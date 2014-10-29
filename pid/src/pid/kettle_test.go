package pid

import (
  "testing"
)

func TestThermometer(t *testing.T) {
  th := NewThermometer()
  if th.Name() != "Thermometer" {
    t.Errorf("Name got %v, expected %v", th.Name(), "Thermometer")
  }

  th.SetInput(69.9)
  if th.Input() != 69.9 {
    t.Errorf("Input got %v, expected %v", th.Input(), 69.9)
  }
  if th.Output() != 69.9 {
    t.Errorf("Output got %v, expected %v", th.Output(), 69.9)
  }

  params := make([]parameter, 0)
  params = append(params, parameter{Name: "granularity", Value: 0.25})
  th.SetParameters(params)
  if th.granularity != 0.25 {
    t.Errorf("Granularity got %v, expected %v", th.granularity, 0.25)
  }
  if th.Output() != 69.75 {
    t.Errorf("Output got %v, expected %v", th.Output(), 69.75)
  }
}

func TestBurner(t *testing.T) {
  b := NewBurner()
  if b.Name() != "Burner" {
    t.Errorf("Name got %v, expected %v", b.Name(), "Burner")
  }

  b.SetInput(1000.0)
  if b.Input() != 1000.0 {
    t.Errorf("Input got %v, expected %v", b.Input(), 1000.0)
  }
}
