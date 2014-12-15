package pid

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	system.Run()
	w.Header().Set("Content-Type", "image/png")
	system.PngWriter().WriteTo(w)
}

// configHandler returns the current paramters and settings.
func configHandler(w http.ResponseWriter, rq *http.Request) {
	system := GenerateSystem("kettle")
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "text/plain")
	enc.Encode(system.AllParameters())
}

// StartHttp starts the HTTP server.
func StartSimulation() {
	readOnlyValues = []string{}
	http.Handle("/", &indexHandler{})
	http.HandleFunc("/config", configHandler)
	http.Handle("/graph", &graphHandler{})
	fmt.Printf("Simulation ready to serve.\n")
	http.ListenAndServe(":8080", nil)
}
