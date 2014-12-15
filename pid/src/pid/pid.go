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
	// Minimum and maximum output values.
	outMin float64
	outMax float64
	// Auto or Manual control mode.
	modeAuto bool
	// Whether a positive output moves the input higher or lower.
	direction int16

	system *System
}

// timeMillis returns the current time as epoch milliseconds.
func timeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

// New returns a new PID object.
func NewPID(s *System, mode, direction int16) *PID {
	p := new(PID)
	p.system = s
	p.SetMode(mode)
	p.SetSampleTime(defaultSampleTime)
	p.SetControllerDirection(direction)
	return p
}

// Status returns the input and output values.
func (p *PID) Status() parameters {
	s := p.system.values[p.Name()]
	s.SetValue("input", p.input)
	s.SetValue("output", p.output)
	return s
}

// Parameters returns the parameters of the PID object.
func (p *PID) Parameters() parameters {
	pa := p.system.parameters[p.Name()]
	pa.SetValue("setpoint", p.Setpoint)
	pa.SetValue("kp", p.dispKp)
	pa.SetValue("ki", p.dispKi)
	pa.SetValue("kd", p.dispKd)
	pa.SetValue("limit_high", p.outMax)
	pa.SetValue("limit_low", p.outMin)
	return pa
}

// SetParameters sets the parameters of the PID object.
func (p *PID) SetParameters(params parameters) {
	var kp, ki, kd, ll, lh float64
	params.GetValueIfPresent("kp", &kp)
	params.GetValueIfPresent("ki", &ki)
	params.GetValueIfPresent("kd", &kd)
	params.GetValueIfPresent("limit_high", &lh)
	params.GetValueIfPresent("limit_low", &ll)
	params.GetValueIfPresent("setpoint", &p.Setpoint)
	p.SetTunings(kp, ki, kd)
	p.SetOutputLimits(ll, lh)
}

// Name returns the name of the PID object.
func (p *PID) Name() string {
	return "PID"
}

// Output gets the output value of the PID.
func (p *PID) Output(interval float64) float64 {
	p.SetSampleTime(int64(interval * 1e3))

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
	return p.output
}

// PID performs a PID computation.
func (p *PID) SetInput(input float64) {
	p.input = input
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

func (p *PID) GetSampleTime() int64 {
	return p.sampleTime
}
