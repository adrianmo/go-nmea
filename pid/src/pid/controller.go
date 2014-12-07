package pid

import (
	"fmt"
	"io"
	"net/url"
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
	// The JSON file holding the system definitions.
	systemJson = "systems.json"
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
	// Parameters loaded from JSON
	parameters systemParameters
}

// Init initalises the System, setting up graphing.
// n is the system name to load from systemJson.
func (s *System) Init(n string) {
	var err error
	s.graph, err = NewGraph()
	if err != nil {
		panic(err)
	}
	var all allSystems
	err = all.ReadJson(systemJson)
	if err != nil {
		panic(err)
	}
	paras, ok := all[n]
	if !ok {
		panic(fmt.Errorf("Cant find system %s in %s", n, systemJson))
	}
	s.parameters = paras.Components
	s.SetParameters(s.parameters[s.Name()])
	if s.Load != nil {
		s.Load.SetParameters(s.parameters[s.Load.Name()])
	}
	if s.Driver != nil {
		s.Driver.SetParameters(s.parameters[s.Driver.Name()])
	}
	if s.Sensor != nil {
		s.Sensor.SetParameters(s.parameters[s.Sensor.Name()])
	}
	if s.Pid != nil {
		s.Pid.SetParameters(s.parameters[s.Pid.Name()])
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

// Parameters returns all parameters set by the system.
func (s *System) Parameters() parameters {
	p := make(parameters, 0)
	ps, ok := s.parameters[s.Name()]
	if !ok {
		return p
	}
	for _, param := range ps {
		p = append(p, param.Copy())
	}
	return p
}

// SetParameters sets the parameter values for the system.
func (s *System) SetParameters(params parameters) {
	s.runTime = params.GetValue("runtime")
	s.interval = params.GetValue("interval")
}

// SetFormParameters sets parameters of the system and components.
// It uses the Form from the web request.
func (s *System) SetFormParameters(values url.Values) {
	params := systemParameters{}
	params.ReadURLValues(values)
	// Now set the parameters for each component.
	if p, ok := params[s.Name()]; ok {
		s.SetParameters(p)
	}
	if p, ok := params[s.Driver.Name()]; ok {
		s.Driver.SetParameters(p)
	}
	if p, ok := params[s.Load.Name()]; ok {
		s.Load.SetParameters(p)
	}
	if p, ok := params[s.Sensor.Name()]; ok {
		s.Sensor.SetParameters(p)
	}
	if p, ok := params[s.Pid.Name()]; ok {
		s.Pid.SetParameters(p)
	}
}

// Process run the controller for one interval cycle.
func (s *System) ProcessInterval() {
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

// Run runs the controller for the full runtime.
func (s *System) Run() {
	for s.time = 0; s.time < s.runTime; s.time += s.interval {
		s.ProcessInterval()
	}
}

// PngWriter renders the graph and returns a WriterTo.
func (s *System) PngWriter() io.WriterTo {
	if err := s.graph.Draw(); err != nil {
		panic(err)
	}
	return s.graph.PngWriter()
}
