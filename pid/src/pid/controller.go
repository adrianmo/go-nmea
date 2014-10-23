package pid

import (
	"io"
	"net/url"
	"strconv"
	"strings"
)

const (
  // How long to run for (seconds).
	runTime = 200
)

var (
  MinPower = 0.0
	MaxPower = 2500.0
)

// A System is a closed loop PID controlled system.
type System struct {
	// The load being driven.
	Load *Kettle
	// The driver.
	Driver *Burner
	// Reads the load's current value/
	Sensor *Thermometer
	// The PID machine.
	Pid *PID
	// Draw pretty graphs.
	graph *Graph
	// Time since start, in seconds.
	time float64
}

// Init initalises the System, setting up graphing.
func (s *System) Init() {
	var err error
	s.graph, err = NewGraph()
	if err != nil {
		panic(err)
	}
	s.graph.AddInput(s.time, s.Sensor.Output())
	s.graph.AddOutput(s.time, s.Pid.Output())
}

// RunToTemperature runs the controller with the given setpoint.
func (s *System) RunToTemperature() {
  sampleInterval := int(s.Pid.GetSampleTime()/1000)
	for i := 0; i < runTime; i++ {
		// sensor -> controller
		s.Pid.SetInput(s.Sensor.Output())
		// controller -> driver
		s.Driver.SetInput(s.Pid.Output())
    // Allow driver to supply load for 5 seconds.
		for j := 0; j < sampleInterval ; j++ {
			// driver -> load
			s.Load.SetInput(s.Driver.Output())
			s.time++
		}
		// load -> sensor
		s.Sensor.SetInput(s.Load.Output())
		// Graph this run.
		s.graph.AddInput(s.time, s.Sensor.Output())
		s.graph.AddOutput(s.time, 100*s.Driver.Output()/MaxPower)
	}
}

// SetParameters sets parameters of the system and components.
// It uses the Form from the web request.
func (s *System) SetParameters(values url.Values) {
	params := make(map[string][]parameter)
	// Extract the parameters and their values for each component.
	for k, v := range values {
		// The form fields are name <component>_<parameter>.
		parts := strings.SplitN(k, "_", 2)
		if len(parts) == 2 && len(v) == 1 {
			component, param := parts[0], parts[1]
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				continue
			}
			p := parameter{Name: param, Value: val}
			params[component] = append(params[component], p)
		}
	}
	// Now set the parameters for each component.
	for c, p := range params {
		if s.Driver.Name() == c {
			s.Driver.SetParameters(p)
		}
		if s.Load.Name() == c {
			s.Load.SetParameters(p)
		}
		if s.Sensor.Name() == c {
			s.Sensor.SetParameters(p)
		}
		if s.Pid.Name() == c {
			s.Pid.SetParameters(p)
		}
	}
}

// PngWriter renders the graph and returns a WriterTo.
func (s *System) PngWriter() io.WriterTo {
	if err := s.graph.Draw(); err != nil {
		panic(err)
	}
	return s.graph.PngWriter()
}
