package pid

import (
  "testing"
)

func TestSystemParameters( t *testing.T) {
  s := &System{runTime: 69.0}
  p := s.Parameters()
  if p[0].Name != "runtime" {
    t.Errorf("Name: Got %s, wanted %s", p[0].Name, "runtime")
  }
  if p[0].Value != 69.0 {
    t.Errorf("runTime: Got %f, wanted %f", p[0].Value, 69.0)
  }
}
