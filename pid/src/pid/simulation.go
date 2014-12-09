package pid

import (
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

// StartHttp starts the HTTP server.
func StartSimulation() {
	http.Handle("/", &indexHandler{})
	http.Handle("/graph", &graphHandler{})
	fmt.Printf("Ready to serve.\n")
	http.ListenAndServe(":8080", nil)
}
