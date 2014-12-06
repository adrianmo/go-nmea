package pid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	indexFilename = "index.html"
)

// An indexHandler handles file requests.
type indexHandler struct {
}

// ServeHTTP returns the contents of the requested file.
func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := path.Base(r.URL.Path)
	if filename == "/" {
		filename = indexFilename
	}
	fmt.Printf("Open: %s\n", filename)
	indexData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(w, "error reading %s: %s", indexFilename, err)
		return
	}
	fmt.Fprintf(w, "%s", indexData)
}

// A graphHandler handles graph requests.
type graphHandler struct {
}

// ServeHTTP returns the graph for the supplied parameters.
func (g *graphHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Printf("/graph error: %v\n", err)
		return
	}
	system := GenerateSystem("kettle")
	system.SetFormParameters(r.Form)
	system.RunToTemperature()
	w.Header().Set("Content-Type", "image/png")
	system.PngWriter().WriteTo(w)
}

// StartHttp starts the HTTP server.
func StartHttp() {
	http.Handle("/", &indexHandler{})
	http.Handle("/graph", &graphHandler{})
	fmt.Printf("Ready to serve.\n")
	http.ListenAndServe(":8080", nil)
}
