package pid

import (
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
	PowerFluctuation = 3
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

  // Interval since last operation.
  interval float64
}

// NewKettle returns an initialised Kettle object.
func NewKettle() *Kettle {
	k := &Kettle{}
	SetComponentDefaults(k)
	return k
}

// Parameters returns the parameters for the kettle.
func (k *Kettle) Parameters() []parameter {
	p := make([]parameter, 0)
	v := parameter{Name: "volume", Title: "Liquid Volume",
		Minimum: 0.0, Maximum: 20.0,
		Step: 1.0, Default: 10.0, Unit: "L", Value: k.Volume}
	p = append(p, v)
	t := parameter{Name: "ambient", Title: "Ambient Temperature",
		Minimum: 0.0, Maximum: 30.0,
		Step: 2.0, Default: 25.0, Unit: "degC", Value: k.AmbientTemperature}
	p = append(p, t)
	x := parameter{Name: "temperature", Title: "Initial Temperature",
		Minimum: 0.0, Maximum: 100.0,
		Step: 1.0, Default: 25.0, Unit: "degC", Value: k.Temperature}
	p = append(p, x)
	l := parameter{Name: "thermal_loss", Title: "Thermal Loss",
		Minimum: 0.0, Maximum: 30.0,
		Step: 1.0, Default: ThermalLoss, Unit: "W/deg", Value: k.ThermalLoss}
	p = append(p, l)
	return p
}

// SetParameters sets the parameter values for the kettle.
func (k *Kettle) SetParameters(params []parameter) {
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

// Name returns the name of the kettle.
func (k *Kettle) Name() string {
	return "Kettle"
}

// Operate operates the kettle for the given time.
func (k *Kettle) operate() {
	/* To work out the current temperature of the kettle:
	 *
	 * - Work out the total heat supplied (Watts*Time/Volume/ThermalMass).
	 * - Work out the total heat lost (ThermalLoss * time * dT).
	 * - Get the difference.
	 */
  // Work out heat added to the kettle.
  k.Temperature += k.getIncrease()
  // Work out and subtract heat lost.
  k.Temperature -= k.getLoss()
}

// SetInput sets the target input watts to the kettle.
func (k *Kettle) SetInput(watts float64) {
	k.Watts = watts
}

// GetInput returns the configured input watts.
func (k *Kettle) Input() float64 {
	return k.Watts
}

// GetOutput returns the kettle temperature.
func (k *Kettle) Output(interval float64) float64 {
  k.interval = interval
	k.operate()
	return k.Temperature
}

// getIncrease calculates the thermal increase from the heat.
func (k *Kettle) getIncrease() float64 {
	tempRise := (k.Watts * k.interval) / (k.Volume * 1e3) / ThermalMass
	return tempRise
}

// getLoss calculates the thermal loss of the vessel in one second.
func (k *Kettle) getLoss() float64 {
	lossWatts := (k.Temperature - k.AmbientTemperature) * k.ThermalLoss * k.interval
	lossCelcius := lossWatts / (k.Volume * 1e3) / ThermalMass
	return lossCelcius
}

// A Burner is a thermal source for the Kettle.
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

// NewBurner returns an initialised Burner object.
func NewBurner() *Burner {
	b := &Burner{}
	SetComponentDefaults(b)
	return b
}

// Parameters returns the parameters for the Burner.
func (b *Burner) Parameters() []parameter {
	p := make([]parameter, 0)
	f := parameter{Name: "fluctuation", Title: "Power Fluctuation",
		Minimum: 0, Maximum: 50,
		Step: 1, Default: PowerFluctuation, Unit: "%", Value: b.PowerFluctuation * 100}
	p = append(p, f)
	i := parameter{Name: "inertia", Title: "Thermal Inertia",
		Minimum: 0, Maximum: 5000.0,
		Step: 10.0, Default: ThermalInertia, Unit: "W/s", Value: b.ThermalInertia}
	p = append(p, i)

	return p
}

// SetParameters sets the parameters for the Burner.
func (b *Burner) SetParameters(params []parameter) {
	for _, p := range params {
		switch p.Name {
		case "fluctuation":
			b.PowerFluctuation = p.Value / 100
		case "inertia":
			b.ThermalInertia = p.Value
		}
	}
}

// Name returns the name of the Burner.
func (b *Burner) Name() string {
	return "Burner"
}

// SetInput sets the current input value of the Burner.
func (b *Burner) SetInput(value float64) {
	b.inputPowerLevel = value
}

// Input returns the current input value of the Burner.
func (b *Burner) Input() float64 {
	return b.inputPowerLevel
}

// Output fetches the current output value of the Burner.
func (b *Burner) Output(interval float64) float64 {
	if b.outputPowerLevel > b.inputPowerLevel {
		b.outputPowerLevel = math.Max(b.inputPowerLevel, b.outputPowerLevel-(ThermalInertia * interval))
	} else if b.outputPowerLevel < b.inputPowerLevel {
		b.outputPowerLevel = math.Min(b.inputPowerLevel, b.outputPowerLevel+(ThermalInertia * interval))
	}
	if b.PowerFluctuation > 0 {
		fluctuation := (rand.Float64()*2.0 - 1.0) * b.PowerFluctuation
		b.outputPowerLevel = b.outputPowerLevel + (b.outputPowerLevel * fluctuation)
	}
	return b.outputPowerLevel
}

// A Thermometer is a temperature sensor for the Kettle.
type Thermometer struct {
	// Current temperature in celcius.
	temperature float64
	// Granularity of the thermometer, in celcius.
	granularity float64
}

// NewThermometer returns an initialised Thermometer object.
func NewThermometer() *Thermometer {
	return new(Thermometer)
}

// Name returns the name of the object.
func (t *Thermometer) Name() string {
	return "Thermometer"
}

// Parameters returns the parameters of the Thermometer.
func (t *Thermometer) Parameters() []parameter {
	p := make([]parameter, 0)
	f := parameter{Name: "granularity", Title: "Reading Granularity",
		Minimum: 0, Maximum: 1,
		Step: 0.01, Default: 0.0, Unit: "deg", Value: t.granularity}
	p = append(p, f)
	return p
}

// SetParameters sets the parameters of the Thermometer.
func (t *Thermometer) SetParameters(params []parameter) {
	for _, p := range params {
		switch p.Name {
		case "granularity":
			t.granularity = p.Value
		}
	}
}

// SetInput sets the input value of the Thermometer.
func (t *Thermometer) SetInput(value float64) {
	t.temperature = value
}

// Input gets the input value of the Thermometer.
func (t *Thermometer) Input() float64 {
	return t.temperature
}

// Output gets the output value of the Thermometer.
func (t *Thermometer) Output(interval float64) float64 {
	if t.granularity > 0 {
		return math.Floor(t.temperature/t.granularity) * t.granularity
	}
	return t.temperature
}

// A KettleSystemGenerator is a generator for a Kettle system.
type KettleSystemGenerator struct{}

// Name returns the name of the generator system.
func (g KettleSystemGenerator) Name() string {
	return "kettle"
}

// Description returns a description of the system.
func (g KettleSystemGenerator) Description() string {
	return "A kettle of liquid with a heating element"
}

// GenerateSystem returns an initialised Kettle system.
func (g KettleSystemGenerator) GenerateSystem() System {
	s := System{}
	SetComponentDefaults(&s)
	s.Load = NewKettle()
	SetComponentDefaults(s.Load)
	s.Driver = NewBurner()
	SetComponentDefaults(s.Driver)
	s.Sensor = NewThermometer()
	SetComponentDefaults(s.Sensor)
	s.Pid = NewPID(Auto, Direct)
	// Pid takes samples every 5 seconds.
	s.Pid.SetSampleTime(5000)
	s.Pid.SetOutputLimits(MinPower, MaxPower)
	s.Pid.Initialize()
	s.Init()
	return s
}

// init registers the Kettle system.
func init() {
	g := KettleSystemGenerator{}
	RegisterSystemGenerator(g)
}
