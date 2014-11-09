package pid

import (
	"io"
	"net/url"
	"strconv"
	"strings"
)

const (
	// How long to run for (seconds).
	runTime = 900

  // default iteration interval (seconds).
  interval = 5.0
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
	// How long to run (seconds).
	runTime float64
	// Time since start, in seconds.
	time float64
  // Interval between system iterations (simulated).
  interval float64
}

// Init initalises the System, setting up graphing.
func (s *System) Init() {
	var err error
	s.graph, err = NewGraph()
	if err != nil {
		panic(err)
	}
}

func (s *System) Name() string {
	return "System"
}

func (s *System) SetInput(v float64) {
}
func (s *System) Input() float64 {
	return 0.0
}

func (s *System) Output(i float64) float64 {
	return 0.0
}

func (s *System) Parameters() []parameter {
	p := make([]parameter, 0)
	v := parameter{Name: "runtime", Title: "Run Time",
		Minimum: 1.0, Maximum: 1000.0,
		Step: 1.0, Default: runTime, Unit: "s", Value: s.runTime}
	p = append(p, v)
	v = parameter{Name: "interval", Title: "Run Interval",
		Minimum: 0.1, Maximum: 10.0,
		Step: 0.1, Default: interval, Unit: "s", Value: s.interval}
	p = append(p, v)
	return p
}

// SetParameters sets the parameter values for the system.
func (s *System) SetParameters(params []parameter) {
	for _, p := range params {
		switch p.Name {
		case "runtime":
			s.runTime = p.Value
		case "interval":
			s.interval = p.Value
		}
	}
}

// RunToTemperature runs the controller with the given setpoint.
func (s *System) RunToTemperature() {
  for s.time = 0 ; s.time < s.runTime ; s.time += s.interval {
		// sensor -> controller
    sensOut := s.Sensor.Output(s.interval)
		s.Pid.SetInput(sensOut)
		s.graph.AddInput(s.time, sensOut)
		// controller -> driver
		s.Driver.SetInput(s.Pid.Output(s.interval))
    // driver -> load
    driveOut := s.Driver.Output(s.interval)
		s.Load.SetInput(driveOut)
		s.graph.AddOutput(s.time, 100*driveOut/MaxPower)
		// load -> sensor
		s.Sensor.SetInput(s.Load.Output(s.interval))
	}
}

// SetFormParameters sets parameters of the system and components.
// It uses the Form from the web request.
func (s *System) SetFormParameters(values url.Values) {
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
		if s.Name() == c {
			s.SetParameters(p)
		}
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
