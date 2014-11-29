package pid

import (
	"testing"
)

type testComponent struct {
	min float64
	max float64
	now float64
}

func (c *testComponent) Name() string             { return "testComponent" }
func (c *testComponent) SetInput(i float64)       {}
func (c *testComponent) Input() float64           { return 0.0 }
func (c *testComponent) Output(i float64) float64 { return 0.0 }
func (c *testComponent) Parameters() parameters {
	p := parameters{}
	p = append(p, &parameter{Name: "min", Default: 5.0, Value: 0.0})
	p = append(p, &parameter{Name: "max", Default: 10.0, Value: 0.0})
	p = append(p, &parameter{Name: "now", Default: 7.0, Value: 0.0})
	return p
}
func (c *testComponent) SetParameters(params parameters) {
	for _, p := range params {
		switch p.Name {
		case "min":
			c.min = p.Value
		case "max":
			c.max = p.Value
		case "now":
			c.now = p.Value
		}
	}
}

type testSystemGenerator struct{}

func (s *testSystemGenerator) Name() string        { return "systemGenerator" }
func (s *testSystemGenerator) Description() string { return "Generator Description" }
func (s *testSystemGenerator) GenerateSystem() System {
	return System{}
}

func TestGetSystems(t *testing.T) {
	// Unregister other generators.
	SystemGenerators = make(map[string]SystemGenerator)
	// Register a test one.
	RegisterSystemGenerator(new(testSystemGenerator))
	ss := GetSystems()
	if len(ss) != 1 {
		t.Errorf("Got %d systems, expected 1", len(ss))
	}
	if ss["systemGenerator"] != "Generator Description" {
		t.Errorf("Got description %s, expected 'Generator Description'", ss["SystemGenerator"])
	}
}
