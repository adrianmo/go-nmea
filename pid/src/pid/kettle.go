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

	kettleJson = "systems.json"
)

type Kettle struct {
	// Volume of the kettle, in litres.
	Volume float64

	// Ambient air temperature, in Celcius.
	AmbientTemperature float64

	// Kettle temperature, in Celcius.
	Temperature float64

	// Thermal loss, in watts per degree.
	ThermalLoss float64

	// Watts of power being supplied.
	Watts float64

	// Interval since last operation.
	interval float64

	// System of ths Kettle.
	system *System
}

// NewKettle returns an initialised Kettle object.
func NewKettle(s *System) *Kettle {
	k := &Kettle{system: s}
	return k
}

// Status returns the input and output values.
func (k *Kettle) Status() parameters {
	p := k.system.values[k.Name()]
	p.SetValue("input", k.Watts)
	p.SetValue("output", k.Temperature)
	return p
}

// Parameters returns the parameters and values for the kettle.
func (k *Kettle) Parameters() parameters {
	p := k.system.parameters[k.Name()]
	p.SetValue("volume", k.Volume)
	p.SetValue("ambient", k.AmbientTemperature)
	p.SetValue("temperature", k.Temperature)
	p.SetValue("thermal_loss", k.ThermalLoss)
	return p
}

// SetParameters sets the parameter values for the kettle.
func (k *Kettle) SetParameters(params parameters) {
	params.GetValueIfPresent("thermal_loss", &k.ThermalLoss)
	params.GetValueIfPresent("ambient", &k.AmbientTemperature)
	params.GetValueIfPresent("temperature", &k.Temperature)
	params.GetValueIfPresent("volume", &k.Volume)
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
	// System of ths Kettle.
	system *System
}

// NewBurner returns an initialised Burner object.
func NewBurner(s *System) *Burner {
	b := &Burner{system: s}
	return b
}

// Status returns the input and output values.
func (b *Burner) Status() parameters {
	p := b.system.values[b.Name()]
	p.SetValue("input", b.inputPowerLevel)
	p.SetValue("output", b.outputPowerLevel)
	return p
}

// Parameters returns the parameters for the Burner.
func (b *Burner) Parameters() parameters {
	p := b.system.parameters[b.Name()]
	p.SetValue("fluctuation", b.PowerFluctuation*100)
	p.SetValue("inertia", b.ThermalInertia)
	return p
}

// SetParameters sets the parameters for the Burner.
func (b *Burner) SetParameters(params parameters) {
	if p, ok := params.Get("fluctuation"); ok {
		b.PowerFluctuation = p.Value / 100
	}
	params.GetValueIfPresent("inertia", &b.ThermalInertia)
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
		b.outputPowerLevel = math.Max(b.inputPowerLevel, b.outputPowerLevel-(ThermalInertia*interval))
	} else if b.outputPowerLevel < b.inputPowerLevel {
		b.outputPowerLevel = math.Min(b.inputPowerLevel, b.outputPowerLevel+(ThermalInertia*interval))
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
	// System of ths Thermometer.
	system *System
}

// NewThermometer returns an initialised Thermometer object.
func NewThermometer(s *System) *Thermometer {
	return &Thermometer{system: s}
}

// Name returns the name of the object.
func (t *Thermometer) Name() string {
	return "Thermometer"
}

// Status returns the input and output values.
func (t *Thermometer) Status() parameters {
	p := t.system.values[t.Name()]
	p.SetValue("input", t.temperature)
	p.SetValue("output", t.temperature)
	return p
}

// Parameters returns the parameters of the Thermometer.
func (t *Thermometer) Parameters() parameters {
	p := t.system.parameters[t.Name()]
	p.SetValue("granularity", t.granularity)
	return p
}

// SetParameters sets the parameters of the Thermometer.
func (t *Thermometer) SetParameters(params parameters) {
	params.GetValueIfPresent("granularity", &t.granularity)
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
	s.Load = NewKettle(&s)
	s.Driver = NewBurner(&s)
	s.Sensor = NewThermometer(&s)
	s.Pid = NewPID(&s, Auto, Direct)
	// Pid takes samples every 5 seconds.
	s.Pid.SetSampleTime(5000)
	s.Pid.SetOutputLimits(MinPower, MaxPower)
	s.Pid.Initialize()
	s.Init(g.Name())
	return s
}

// init registers the Kettle system.
func init() {
	g := KettleSystemGenerator{}
	RegisterSystemGenerator(g)
}
