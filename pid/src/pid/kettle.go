
package pid

import (
  "fmt"
  "math"
  "math/rand"
)

const (
  // Thermal mass in Joules per gram per degree.
  ThermalMass = 4.2

  // Thermal loss of the kettle. Watts per degree.
  ThermalLoss = 13

  // Thermal inertia of the Burner. This is the rate of change
  // of power, in watts per second.
  ThermalInertia = 500

  // Power fluctuation of the Burner. Each time Operate() is called, the input power
  // will fluctuate by this percent (0.0-1.0).
  PowerFluctuation = 0.1
)

type Kettle struct {
  // Volume of the kettle, in litres.
  Volume float64

  // Ambient temperature, in Celcius.
  AmbientTemperature float64

  // Kettle temperature, in Celcius.
  Temperature float64

  // Thermal loss, in watts per degree.
  ThermalLoss float64

  // Watts of power being supplied.
  Watts float64
}

func NewKettle() *Kettle {
  k := &Kettle{}
  SetComponentDefaults(k)
  return k
}

func (k *Kettle) Parameters() []parameter {
  p := make([]parameter, 0)
  v := parameter{Name: "volume", Title: "Liquid Volume",
                 Minimum: 0.0, Maximum: 20.0,
                 Step: 1.0, Default: 10.0, Unit: "L"}
  p = append(p, v)
  t := parameter{Name: "ambient", Title: "Ambient Temperature",
                 Minimum: 0.0, Maximum: 30.0,
                 Step: 2.0, Default: 25.0, Unit: "degC"}
  p = append(p, t)
  x := parameter{Name: "temperature", Title: "Initial Temperature",
                 Minimum: 0.0, Maximum: 100.0,
                 Step: 1.0, Default: 25.0, Unit: "degC"}
  p = append(p, x)
  l := parameter{Name: "thermal_loss", Title: "Thermal Loss",
                 Minimum: 0.0, Maximum: 30.0,
                 Step: 1.0, Default: 13.0, Unit: "W/deg"}
  p = append(p, l)
  return p
}

func (k *Kettle) SetParameters(params []parameter) {
  fmt.Printf("%v\n", params)
  for _, p := range params {
    switch p.Name {
    case "thermal_loss":
      k.ThermalLoss = p.Value
    case "ambient":
      k.AmbientTemperature = p.Value
    case "temperature":
      k.Temperature = p.Value
    case "volume":
      k.Volume = p.Value
    }
  }
}

// Operate operates the kettle for the given time.
func (k *Kettle) operate(seconds int) {
  /* To work out the current temperature of the kettle:
   *
   * - Work out the total heat supplied (Watts*Time/Volume/ThermalMass).
   * - Work out the total heat lost (ThermalLoss * time * dT).
   * - Get the difference.
   */
   for s := 0 ; s < seconds ; s++ {
     // Work out heat added to the kettle.
     k.Temperature += k.getIncrease()
     // Work out and subtract heat lost.
     k.Temperature -= k.getLoss()
   }
}

// SetInput sets the target input watts to the kettle.
func (k *Kettle) SetInput(watts float64) {
  k.Watts = watts
  k.operate(1)
}

// GetInput returns the configured input watts.
func (k *Kettle) Input() float64 {
  return k.Watts
}

// GetOutput returns the kettle temperature.
func (k *Kettle) Output() float64 {
  return k.Temperature
}

// getIncrease calculates the thermal increase from the heat.
func (k *Kettle) getIncrease() float64 {
   tempRise := k.Watts / (k.Volume * 1e3) / ThermalMass
   return tempRise
}

// getLoss calculates the thermal loss of the vessel in one second.
func (k *Kettle) getLoss() float64 {
  lossWatts := (k.Temperature - k.AmbientTemperature) * k.ThermalLoss
  lossCelcius := lossWatts / (k.Volume * 1e3) / ThermalMass
  return lossCelcius
}


type Burner struct {
  // How much the output power fluctuates (0.0 - 1.0)
  PowerFluctuation float64
  // How many watts per second the output power lags the input.
  ThermalInertia float64

  // Input power value (Watts)
  inputPowerLevel float64
  // Output power value (Watts)
  outputPowerLevel float64
}

func NewBurner() *Burner {
  b := &Burner{}
  SetComponentDefaults(b)
  return b
}

func (b *Burner) Parameters() []parameter {
  p := make([]parameter, 0)
  f := parameter{Name: "fluctuation", Title: "Power Fluctuation",
                 Minimum: 0.0, Maximum: 1.0,
                 Step: 0.01, Default: 0.05, Unit: "%"}
  p = append(p, f)
  i := parameter{Name: "inertia", Title: "Thermal Inertia",
                 Minimum: 0, Maximum: 5000.0,
                 Step: 10.0, Default: 5000.0, Unit: "W/s"}
  p = append(p, i)

  return p
}

func (b *Burner) SetParameters(params []parameter) {
  for _, p := range params {
    switch p.Name {
    case "fluctuation":
      b.PowerFluctuation = p.Value
    case "inertia":
      b.ThermalInertia = p.Value
    }
  }
}

func (b *Burner) SetInput (value float64) {
  b.inputPowerLevel = value
}

func (b *Burner) Input () float64 {
  return b.inputPowerLevel
}

func (b *Burner) Output () float64 {
  if b.outputPowerLevel > b.inputPowerLevel {
    b.outputPowerLevel = math.Max(b.inputPowerLevel, b.outputPowerLevel - ThermalInertia)
  } else if b.outputPowerLevel < b.inputPowerLevel {
    b.outputPowerLevel = math.Min(b.inputPowerLevel, b.outputPowerLevel + ThermalInertia)
  }
  if b.PowerFluctuation > 0 {
    fluctuation := (rand.Float64()*2.0-1.0) * b.PowerFluctuation
    b.outputPowerLevel = b.outputPowerLevel + (b.outputPowerLevel * fluctuation)
  }
  return b.outputPowerLevel
}

type Thermometer struct {
  temperature float64
}

func NewThermometer() *Thermometer {
  return &Thermometer{}
}

func (t *Thermometer) Parameters() []parameter {
  p := make([]parameter, 0)
  return p
}

func (t *Thermometer) SetParameters(params []parameter) {
}

func (t *Thermometer) SetInput (value float64) {
  t.temperature = value
}

func (t *Thermometer) Input () float64 {
  return t.temperature
}

func (t *Thermometer) Output () float64 {
  return t.temperature
}

type KettleSystemGenerator struct{}

func (g KettleSystemGenerator) Name() string {
  return "kettle"
}
func (g KettleSystemGenerator) GenerateSystem() System {
  s := System{}
  s.Load = NewKettle()
  SetComponentDefaults(s.Load)
  s.Driver = NewBurner()
  SetComponentDefaults(s.Driver)
  s.Sensor = NewThermometer()
  SetComponentDefaults(s.Sensor)
  return s;
}

func init() {
  g := KettleSystemGenerator{}
  RegisterSystemGenerator(g)
}
