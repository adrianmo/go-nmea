
package pid

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "path"
  "strconv"
)

const (
  indexFilename = "index.html"
)

type indexHandler struct {

}

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

var (
  float_params = []string{"loss", "setpoint", "maxpower", "inertia", "volume", "fluctuation", "kp", "kd", "ki"}
)

type graphHandler struct {
}

// parseFloats parses all of the floating point parameters.
func (g *graphHandler) parseFloats(r *http.Request) (map[string]float64, error) {
  m := make(map[string]float64)
  for _, param := range float_params {
    value, ok := r.Form[param]
    if !ok {
      return m, fmt.Errorf("missing param '%s'", param)
    }
    float, err := strconv.ParseFloat(value[0], 64)
    if err != nil {
      return m, fmt.Errorf("error parsing '%s': %s", param, err)
    }
    m[param] = float
  }
  return m, nil
}

func (g *graphHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm() ; err != nil {
    fmt.Printf("/graph error: %v\n", err)
    return
  }
  floats, err := g.parseFloats(r)
  if err != nil {
    fmt.Printf("%s\n", err)
    return
  }
  MaxPower = floats["maxpower"]
  system := &System{}
  system.Init()
  system.Driver.PowerFluctuation = floats["fluctuation"]/1000
  system.Driver.ThermalInertia = floats["inertia"]
  system.Load.Volume = floats["volume"]
  system.Load.ThermalLoss = floats["loss"]
  system.Pid.SetTunings(floats["kp"], floats["ki"]/100, floats["kd"])
  system.RunToTemperature(floats["setpoint"])
  system.PngWriter().WriteTo(w)
}

func StartHttp() {
  http.Handle("/", &indexHandler{})
  http.Handle("/graph", &graphHandler{})
  fmt.Printf("Ready to serve.\n")
  http.ListenAndServe(":8080", nil)
}
