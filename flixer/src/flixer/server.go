// Package flixer is the flixer http server.
package flixer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	// Directory containing files to serve.
	htmlRoot  = "files"
	indexHTML = "index.html"
)

var (
	port int
)

// Server config and user settings to pass to client browser.
type Status struct {
	// IP address of the client.
	Client string
	// List of all supported regions.
	Regions []string
	// Client's currently selected region.
	Region string
}

func logRequest(r *http.Request) {
	log.Printf("[%s] %s", remoteHost(r.RemoteAddr), r.URL.String())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	// Remove leading slash.
	path := strings.TrimLeft(r.URL.Path, "/")
	// Ensure client is talking to this host directly.
	if !strings.Contains(r.Host, ":") {
		http.Redirect(w, r, fmt.Sprintf("http://%s:%d/%s", r.Host, port, path),
			http.StatusMovedPermanently)
		return
	}
	// Default to the index html if empty path.
	if len(path) == 0 || strings.Contains(path, "/") {
		path = indexHTML
	}
	// Read the file.
	data, err := ioutil.ReadFile(filepath.Join(htmlRoot, path))
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading %s", path)
		return
	}
	// Write the file to the client.
	if _, err := w.Write(data); err != nil {
		log.Printf("Error writing socket: %s", err)
		return
	}
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	// Extract the selected region.
	c, ok := r.URL.Query()["country"]
	if !ok || len(c) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing parameter 'country'")
		return
	}
	// Fetch the supported regions.
	regions, err := RegionCfg()
	if err != nil {
		log.Printf("Error reading config: %v", err)
		return
	}
	// Check that the region is supported.
	cc, ok := regions[c[0]]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported country: %s", c[0])
		return
	}
	fmt.Fprintf(w, "Setting country to: %s [%s]", c[0], cc)
	// Set user region.
	_, err = UserRegion(remoteHost(r.RemoteAddr), c[0])
	if err != nil {
		log.Printf("Error setting region: %v", err)
		return
	}

	// Update IPTables for new country.
	i := IPTables{}
	if err := i.Commit(); err != nil {
		log.Printf("Error updating iptables: %v", err)
	}

	// Restart sniproxy to clear connections.
	s := SNIProxy{}
	if err := s.Commit(); err != nil {
		log.Printf("Error updating sniproxy: %v", err)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	host := remoteHost(r.RemoteAddr)
	// Get users's current region.
	region, err := UserRegion(host, "")
	if err != nil {
		log.Printf("Error marshalling status: %v", err)
		return
	}
	// Fetch all regions available.
	regions, err := Regions()
	if err != nil {
		log.Printf("Error reading config: %v", err)
		return
	}
	// Build status struct to send to client.
	st := Status{
		Regions: regions,
		Region:  region,
		Client:  host,
	}
	js, err := json.Marshal(st)
	if err != nil {
		log.Printf("Error marshalling data: %v\n", err)
		return
	}
	// Send!
	if _, err := w.Write(js); err != nil {
		log.Printf("Error writing socket: %s", err)
		return
	}
}

func remoteHost(addr string) string {
	return strings.Split(addr, ":")[0]
}

func Start(a string, p int) error {
	// Setup iptables with saved rules.
	i := IPTables{}
	if err := i.Commit(); err != nil {
		log.Printf("Error updating iptables: %v", err)
	} else {
		log.Printf("Set saved iptables rules.")
	}
	port = p
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/status", statusHandler)
	log.Printf("Listening on port %d", port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", a, port), nil)
}
