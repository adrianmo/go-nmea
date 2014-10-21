/* Package PID is an implementation of a Proportional, Integral Differential
 * analogue controller.
 */

package pid

import (
	"time"
)

const (
	// Direction of input change when output is changed.
	Direct  = 0
	Reverse = 1

	// Mode of the controller.
	Auto   = 0
	Manual = 1

	// Sample frequency in millis.
	defaultSampleTime = 100
	// Low and high output limit defaults.
	defaultLimitLow  = 0
	defaultLimitHigh = 255
)

type PID struct {
	// The value to aim for.
	Setpoint float64
	// Input and output variables.
	input     float64
	output    float64
	iTerm     float64
	lastInput float64
	// PID algorithm tunings (post-mangling).
	Kp float64
	Ki float64
	Kd float64
	// PID tunings (pre-mangling)
	dispKp float64
	dispKi float64
	dispKd float64
	// Millis between samples.
	sampleTime int64
	lastTime   int64
	// Minimum and maximum output values.
	outMin float64
	outMax float64
	// Auto or Manual control mode.
	modeAuto bool
	// Whether a positive output moves the input higher or lower.
	direction int16
}

// timeMillis returns the current time as epoch milliseconds.
func timeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

// New returns a new PID object.
func New(kp, ki, kd, setPoint float64, mode, direction int16) *PID {
	p := new(PID)
	p.SetMode(mode)
	p.SetSampleTime(defaultSampleTime)
	p.SetControllerDirection(direction)
	SetComponentDefaults(p)
	p.lastTime = timeMillis() - defaultSampleTime
	return p
}

// Parameters returns the parameters of the PID object.
func (p *PID) Parameters() []parameter {
	params := make([]parameter, 0)
	sp := parameter{Name: "setpoint", Title: "Set Point",
		Minimum: 0, Maximum: 100,
		Step: 1, Default: 50, Unit: "deg",
		Value: p.Setpoint,
	}
	params = append(params, sp)
	kp := parameter{Name: "kp", Title: "kp",
		Minimum: 0, Maximum: 10000,
		Step: 10, Default: 6000, Unit: "",
		Value: p.dispKp,
	}
	params = append(params, kp)
	ki := parameter{Name: "ki", Title: "ki",
		Minimum: 0, Maximum: 5,
		Step: 0.1, Default: 0.1, Unit: "",
		Value: p.dispKi,
	}
	params = append(params, ki)
	kd := parameter{Name: "kd", Title: "kd",
		Minimum: 0, Maximum: 5000,
		Step: 10, Default: 10, Unit: "",
		Value: p.dispKd,
	}
	params = append(params, kd)
	lh := parameter{Name: "limit_high", Title: "High Limit",
		Minimum: 0, Maximum: 3000,
		Step: 1, Default: 2000, Unit: "",
		Value: p.outMax,
	}
	params = append(params, lh)
	ll := parameter{Name: "limit_low", Title: "Lower Limit",
		Minimum: 0, Maximum: 3000,
		Step: 1, Default: 0, Unit: "",
		Value: p.outMin,
	}
	params = append(params, ll)

	return params
}

// SetParameters sets the parameters of the PID object.
func (p *PID) SetParameters(params []parameter) {
	var kp, ki, kd, ll, lh float64
	for _, param := range params {
		switch param.Name {
		case "kp":
			kp = param.Value
		case "ki":
			ki = param.Value
		case "kd":
			kd = param.Value
		case "limit_high":
			ll = param.Value
		case "limit_low":
			lh = param.Value
		case "setpoint":
			p.Setpoint = param.Value
		}
	}
	p.SetTunings(kp, ki, kd)
	p.SetOutputLimits(ll, lh)
}

// Name returns the name of the PID object.
func (p *PID) Name() string {
	return "PID"
}

// Output gets the output value of the PID.
func (p *PID) Output() float64 {
	return p.output
}

// PID performs a PID computation.
func (p *PID) SetInput(input float64) {
	/*
		now := timeMillis()
		timeChange := now - p.lastTime
		if timeChange <= p.sampleTime {
			return false
		}
	*/

	p.input = input

	err := p.Setpoint - p.input
	p.iTerm += (p.input * err)

	if p.iTerm > p.outMax {
		p.iTerm = p.outMax
	} else if p.iTerm < p.outMin {
		p.iTerm = p.outMin
	}

	dInput := p.input - p.lastInput

	output := p.Kp*err + p.Ki*p.iTerm - p.Kd*dInput

	if output > p.outMax {
		output = p.outMax
	} else if output < p.outMin {
		output = p.outMin
	}
	p.output = output

	p.lastInput = p.input
	//	p.lastTime = now
}

// Input returns the input value of the PID.
func (p *PID) Input() float64 {
	return p.input
}

// SetTunings sets the Kp/Ki/Kd tuning parmeters.
func (p *PID) SetTunings(kp, ki, kd float64) {
	if kp < 0 || ki < 0 || kd < 0 {
		return
	}

	p.dispKp = kp
	p.dispKi = ki
	p.dispKd = kd

	sampleTimeSec := float64(p.sampleTime) / 1000

	p.Kp = kp
	p.Ki = ki * sampleTimeSec
	p.Kd = kd / sampleTimeSec

	if p.direction == Reverse {
		p.Kp = (0 - p.Kp)
		p.Ki = (0 - p.Ki)
		p.Kd = (0 - p.Kd)
	}
}

// SetSampleTime sets the PID sample time in milliseconds.
func (p *PID) SetSampleTime(sampleTime int64) {
	if sampleTime > 0 {
		ratio := float64(sampleTime) / float64(p.sampleTime)
		p.Ki *= ratio
		p.Kd /= ratio
		p.sampleTime = int64(sampleTime)
	}
}

// SetOutputLimits sets the lower and upper output limits.
func (p *PID) SetOutputLimits(min, max float64) {
	if min >= max {
		return
	}
	p.outMin = min
	p.outMax = max

	if p.modeAuto {
		if p.output > p.outMax {
			p.output = p.outMax
		} else if p.output < p.outMin {
			p.output = p.outMin
		}

		if p.iTerm > p.outMax {
			p.iTerm = p.outMax
		} else if p.iTerm < p.outMin {
			p.iTerm = p.outMin
		}
	}
}

// SetMode sets the PID mode to auto or manual.
func (p *PID) SetMode(mode int16) {
	var newIsAuto = bool(mode == Auto)

	if newIsAuto != p.modeAuto {
		p.Initialize()
	}
	p.modeAuto = newIsAuto
}

// Initialize sets up the controller.
func (p *PID) Initialize() {
	p.iTerm = p.output
	p.lastInput = p.input
	if p.iTerm > p.outMax {
		p.iTerm = p.outMax
	} else if p.iTerm < p.outMin {
		p.iTerm = p.outMin
	}
}

// SetControllerDirection sets the input to output direction ratio.
func (p *PID) SetControllerDirection(direction int16) {
	if p.modeAuto && p.direction != direction {
		p.Kp = (0 - p.Kp)
		p.Ki = (0 - p.Ki)
		p.Kd = (0 - p.Kd)
	}
	p.direction = direction
}

// GetKp returns the Kp tuning parameter.
func (p *PID) GetKp() float64 {
	return p.dispKp
}

// GetKi returns the Ki tuning parameter.
func (p *PID) GetKi() float64 {
	return p.dispKi
}

// GetKd returns the Kd tuning parameter.
func (p *PID) GetKd() float64 {
	return p.dispKd
}

// GetMode returns PID mode (Auto or Manual).
func (p *PID) GetMode() int32 {
	if p.modeAuto {
		return Auto
	}
	return Manual
}

// GetDirection returns the in/out direction ratio.
func (p *PID) GetDirection() int16 {
	return p.direction
}
