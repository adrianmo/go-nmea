package pid

import (
	"encoding/json"
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

type systemDescriptor struct {
  Name string
  Description string
  Components []*componentDescriptor
}

type componentDescriptor struct {
  Name string
  Parameters interface{}
}

// A systemsHandler handles the /systems URL.
type systemsHandler struct{}

// ServeHTTP returns the JSOn encoded description of the systems.
func (s *systemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//systems := make(map[string]map[string]interface{})
  systems := make([]*systemDescriptor, 0)
	for name, description := range GetSystems() {
		system := GenerateSystem(name)
    sd := new(systemDescriptor)
    sd.Name = name
    sd.Description = description
    sd.Components = make([]*componentDescriptor, 0)

    cd := new(componentDescriptor)
    cd.Name = system.Name()
    cd.Parameters = system.Parameters()
    sd.Components = append(sd.Components, cd)

    cd = new(componentDescriptor)
    cd.Name = system.Pid.Name()
    cd.Parameters = system.Pid.Parameters()
    sd.Components = append(sd.Components, cd)

    cd = new(componentDescriptor)
    cd.Name = system.Load.Name()
    cd.Parameters = system.Load.Parameters()
    sd.Components = append(sd.Components, cd)

    cd = new(componentDescriptor)
    cd.Name = system.Driver.Name()
    cd.Parameters = system.Driver.Parameters()
    sd.Components = append(sd.Components, cd)

    cd = new(componentDescriptor)
    cd.Name = system.Sensor.Name()
    cd.Parameters = system.Sensor.Parameters()
    sd.Components = append(sd.Components, cd)

    systems = append(systems, sd)
	}
	e := json.NewEncoder(w)
	e.Encode(systems)
}

// StartHttp starts the HTTP server.
func StartHttp() {
	http.Handle("/", &indexHandler{})
	http.Handle("/graph", &graphHandler{})
	http.Handle("/systems", &systemsHandler{})
	fmt.Printf("Ready to serve.\n")
	http.ListenAndServe(":8080", nil)
}
