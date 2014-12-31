package pid

import (
	"log"
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
	Parameters() parameters
	// SetParameters sets the parameters of the component.
	SetParameters(parameters)
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
	log.Printf("Registered system: %s\n", g.Name())
}
