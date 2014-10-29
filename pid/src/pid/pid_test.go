package pid

import (
  "testing"
)

func TestPIDMode(t *testing.T) {
  p := new(PID)
  p.SetMode(Auto)
  if p.GetMode() != Auto {
    t.Errorf("GetMode() got %v, expected Auto (%v)", p.GetMode(), Auto)
  }
  p.SetMode(Manual)
  if p.GetMode() != Manual {
    t.Errorf("GetMode() got %v, expected Manual (%v)", p.GetMode(), Manual)
  }
}

func TestPIDDirection(t *testing.T) {
  p := new(PID)
  p.sampleTime = 10
  p.Ki = 100
  p.Kd = 100
  p.SetSampleTime(100)
  if int(p.Ki) != 1000 {
    t.Errorf("Ki got %v, expected 1000", p.Ki)
  }
  if int(p.Kd) != 10 {
    t.Errorf("Kd got %v, expected 10", p.Kd)
  }
  if p.sampleTime != 100 {
    t.Errorf("sampleTime got %v, expected %v", p.sampleTime, 100)
  }
}

func TestSetParameters(t *testing.T) {
  pid := new(PID)
  pid.SetSampleTime(5000)
  params := make([]parameter, 0)
  p := parameter{Name: "kp", Value: 1000}
  params = append(params, p)
  p = parameter{Name: "ki", Value: 2}
  params = append(params, p)
  p = parameter{Name: "kd", Value: 3}
  params = append(params, p)
  p = parameter{Name: "limit_high", Value: 2000}
  params = append(params, p)
  p = parameter{Name: "limit_low", Value: 100}
  params = append(params, p)
  p = parameter{Name: "setpoint", Value: 500}
  params = append(params, p)
  pid.SetParameters(params)
  if pid.outMax != 2000 {
    t.Errorf("outMax got %v, wanted %v", pid.outMax, 2000)
  }
  if pid.outMin != 100 {
    t.Errorf("outMiingot %v, wanted %v", pid.outMax, 100)
  }
  if pid.Kp != 1000 {
    t.Errorf("Kp got %v, wanted %v", pid.Kp, 1000)
  }
  if pid.Ki != 10 {
    t.Errorf("Kp got %v, wanted %v", pid.Ki, 10)
  }
  if pid.Kd != 0.6 {
    t.Errorf("Kd got %v, wanted %v", pid.Kd, 0.6)
  }

  paramValues := []float64{500, 1000, 2, 3, 2000, 100}
  params = pid.Parameters()
  for i := range paramValues {
    if params[i].Value != paramValues[i] {
      t.Errorf("Param got %v, wanted %v", params[i].Value, paramValues[i])
    }
  }
}
