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
  systemJson = "testdata/systems.json"
	pid := new(PID)
	pid.system = &System{}
	pid.system.Init("kettle")
	pid.SetSampleTime(5000)
	params := make(parameters, 0)
	p := &parameter{Name: "kp", Value: 1000, Minimum: 0, Maximum: 2000}
	params = append(params, p)
	p = &parameter{Name: "ki", Value: 2, Minimum: 0, Maximum: 100}
	params = append(params, p)
	p = &parameter{Name: "kd", Value: 3, Minimum: 0, Maximum: 100}
	params = append(params, p)
	p = &parameter{Name: "limit_high", Value: 2000, Minimum: 0, Maximum: 3000}
	params = append(params, p)
	p = &parameter{Name: "limit_low", Value: 100, Minimum: 0, Maximum: 500}
	params = append(params, p)
	p = &parameter{Name: "setpoint", Value: 500, Minimum: 0, Maximum: 500}
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

	paramValues := map[string]float64{
		"setpoint": 500, "kp": 1000, "ki": 2, "kd": 3, "limit_high": 2000, "limit_low": 100}
	params = pid.Parameters()
	for k, v := range paramValues {
		if params.GetValue(k) != v {
			t.Errorf("Param %s got %v, wanted %v", k, params.GetValue(k), v)
		}
	}
}
