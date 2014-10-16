
package pid

import (
  "fmt"
  "io"
)

const (
  minPower = 100
  setPointTime = 200

  /*
  kP = 25
  kI = 35
  kD = 20000
  */

  kP = 6000
  kI = 0.25
  kD = 5000
)

var (
  MaxPower = 2500.0
  SystemGenerators map[string]SystemGenerator
)

func RegisterSystemGenerator(g SystemGenerator) {
  SystemGenerators[g.Name()] = g
  fmt.Printf("Registered system: %s\n", g.Name())
}

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


func (s *System) Init() {
  var err error
  s.Load = NewKettle()
  s.Driver = NewBurner()
  s.Sensor = NewThermometer()

  s.Pid = New(kP, kI, kD, 80, Auto, Direct)
  s.Pid.SetSampleTime(5000)
  s.Pid.SetOutputLimits(minPower, MaxPower)
  s.Pid.Initialize()

  s.graph, err = NewGraph()
  if err != nil {
    panic(err)
  }
  s.graph.AddInput(s.time, s.Sensor.Output())
  s.graph.AddOutput(s.time, s.Pid.Output())
}

// RunToTemperature runs the controller with the given setpoint.
func (s *System) RunToTemperature(t float64) {
  s.Pid.Setpoint = t
  for i := 0 ; i < setPointTime ; i++ {
    // sensor -> controller
    s.Pid.SetInput(s.Sensor.Output())
    // controller -> driver
    s.Driver.SetInput(s.Pid.Output())
    for j := 0 ; j < 5 ; j++ {
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

// PngWriter renders the graph and returns a WriterTo.
func (s *System) PngWriter() io.WriterTo {
  if err := s.graph.Draw() ; err != nil {
    panic(err)
  }
  return s.graph.PngWriter()
}

func init() {
  SystemGenerators = make(map[string]SystemGenerator)
}
