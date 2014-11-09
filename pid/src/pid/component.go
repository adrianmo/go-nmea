package pid

import (
	"fmt"
)

var (
	// The generators for each type of system.
	SystemGenerators map[string]SystemGenerator
)

// An IO component is a component in a closed system.
type IOComponent interface {
	// GetName gets the human readable name of the component.
	Name() string
	// SetInput sets the input value for the component.
	SetInput(float64)
	// Input gets the input value for the component.
	Input() float64
	// Output gets the output value for the component, after the given duration has elapsed.
	Output(duration float64) float64
	// Parameters gets the settable parameters of the component.
	Parameters() []parameter
	// SetParameters sets the parameters of the component.
	SetParameters([]parameter)
}

// A parameter is a configurable parameter for an IOComponent.
type parameter struct {
	// Name is the shortname of the parameter.
	Name string
	// Title is a more human readable name (for UI display)
	Title string
	// Minimum value of the parameter.
	Minimum float64
	// Maximum value of the parameter.
	Maximum float64
	// The desired step size of the parameter.
	Step float64
	// The default value.
	Default float64
	// The units to display.
	Unit string
	// The value (for setting parameters)
	Value float64
}

type Driver interface {
	IOComponent
}

type Load interface {
	IOComponent
}

type Sensor interface {
	IOComponent
}

// SetComponentDefaults sets the parameters of the component to its defaults.
func SetComponentDefaults(c IOComponent) {
	params := c.Parameters()
	for i := 0; i < len(params); i++ {
		params[i].Value = params[i].Default
	}
	c.SetParameters(params)
}

// A SystemGenerator is registered on init, and is used to build a complete
// System struct when a simulation is requested.
type SystemGenerator interface {
	// Name returns the name of the Generator.
	Name() string
	// Description returns a more complete description of the system.
	Description() string
	// GenerateSystem returns a System.
	GenerateSystem() System
}

// Generate returns a System of the given type.
func GenerateSystem(s string) System {
	return SystemGenerators[s].GenerateSystem()
}

// GetSystems gets the available systems.
func GetSystems() map[string]string {
	systems := make(map[string]string)
	for k, v := range SystemGenerators {
		systems[k] = v.Description()
	}
	return systems
}

//RegisterSystemGenerator registers a system generator module.
func RegisterSystemGenerator(g SystemGenerator) {
	if SystemGenerators == nil {
		SystemGenerators = make(map[string]SystemGenerator)
	}
	SystemGenerators[g.Name()] = g
	fmt.Printf("Registered system: %s\n", g.Name())
}
