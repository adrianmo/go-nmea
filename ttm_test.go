package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTTM(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  TTM
	}{
		{
			name: "good sentence",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,,,M*2A",
			msg: TTM{
				BaseSentence:      BaseSentence{},
				TargetNumber:      2,
				TargetDistance:    1.43,
				Bearing:           170.5,
				BearingType:       "T",
				TargetSpeed:       0.16,
				TargetCourse:      264.4,
				CourseType:        "T",
				DistanceCPA:       1.42,
				TimeCPA:           36.9,
				SpeedUnits:        "N",
				TargetName:        "",
				TargetStatus:      "T",
				ReferenceTarget:   "",
				TimeUTC:           Time{Valid: false, Hour: 0, Minute: 0, Second: 0, Millisecond: 0},
				TypeOfAcquisition: "M",
			},
		},
		{
			name: "invalid nmea: TargetNumber",
			raw:  "$RATTM,x02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid target number: x02",
		},
		{
			name: "invalid nmea: TargetDistance",
			raw:  "$RATTM,02,x1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid target Distance: x1.43",
		},
		{
			name: "invalid nmea: Bearing",
			raw:  "$RATTM,02,1.43,x170.5,T,0.16,264.4,T,1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid bearing: x170.5",
		},
		{
			name: "invalid nmea: BearingType",
			raw:  "$RATTM,02,1.43,170.5,xT,0.16,264.4,T,1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid bearing type: xT",
		},
		{
			name: "invalid nmea: TargetSpeed",
			raw:  "$RATTM,02,1.43,170.5,T,x0.16,264.4,T,1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid target speed: x0.16",
		},
		{
			name: "invalid nmea: TargetCourse",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,x264.4,T,1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid target course: x264.4",
		},
		{
			name: "invalid nmea: CourseType",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,xT,1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid course type: xT",
		},
		{
			name: "invalid nmea: DistanceCPA",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,x1.42,36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid distance CPA: x1.42",
		},
		{
			name: "invalid nmea: TimeCPA",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,x36.9,N,,T,,,M*52",
			err:  "nmea: RATTM invalid time of CPA: x36.9",
		},
		{
			name: "invalid nmea: SpeedUnits",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,xN,,T,,,M*52",
			err:  "nmea: RATTM invalid speed units: xN",
		},
		{
			name: "invalid nmea: ReferenceTarget",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,x,,M*52",
			err:  "nmea: RATTM invalid reference target: x",
		},
		{
			name: "invalid nmea: ReferenceTarget",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,x,,M*52",
			err:  "nmea: RATTM invalid reference target: x",
		},
		{
			name: "invalid nmea: TimeUTC",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,,x,M*52",
			err:  "nmea: RATTM invalid UTC time: x",
		},
		{
			name: "invalid nmea: TypeOfAcquisition",
			raw:  "$RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,,x,M*52",
			err:  "nmea: RATTM invalid UTC time: x",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mm := m.(TTM)
				mm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mm)
			}
		})
	}
}
