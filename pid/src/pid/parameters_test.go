package pid

import (
	"reflect"
	"testing"
)

func TestParameter(t *testing.T) {
	p := &parameter{
		Name:    "duration",
		Title:   "The duration",
		Minimum: 0.1,
		Maximum: 1.2,
		Step:    0.2,
		Default: 0.5,
		Unit:    "s",
		Value:   0.7,
	}

	p1 := p.Copy()

	if !reflect.DeepEqual(p, p1) {
		t.Errorf("Copy() failed.\np: %v\np1: %v\n", p, p1)
	}
}

func TestParameters(t *testing.T) {
	p1 := &parameter{Name: "length", Value: 1.0}
	p2 := &parameter{Name: "width", Value: 2.0}
	p3 := &parameter{Name: "height", Value: 3.0}
	p4 := &parameter{Name: "volume", Value: 4.0}
	p5 := &parameter{Name: "mass", Value: 5.0}

	ps := parameters{p1, p2, p3, p4, p5}

	// Test Get() functionality.
	p, ok := ps.Get("height")
	if !ok {
		t.Errorf("Get() was not ok.")
	} else if p.Name != "height" || p.Value != 3.0 {
		t.Errorf("Returned was %v, expected %v", p, p3)
	}

	p, ok = ps.Get("invalid")
	if ok {
		t.Errorf("Returned was ok, should be not ok.")
	}

	// Test GetValue() functionality.
	if v := ps.GetValue("height"); v != 3.0 {
		t.Errorf("GetValue() got %v, expected %v", v, 3.0)
	}
	if v := ps.GetValue("invalid"); v != 0.0 {
		t.Errorf("GetValue() got %v, expected %v", v, 0.0)
	}

	// Test SetValue() functionality.
	ps.SetValue("height", 3.5)
	if v := ps.GetValue("height"); v != 3.5 {
		t.Errorf("GetValue() got %v, expected %v", v, 3.5)
	}
}
