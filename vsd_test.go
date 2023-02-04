package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVSD(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  VSD
	}{
		{
			name: "good sentence",
			raw:  "$RAVSD,0,4.5,6,@@@@@@@@@@@@@@@@@@@@,220516,01,02,8,*6E",
			msg: VSD{
				TypeOfShipAndCargo:    Int64{Value: 0, Valid: true},
				StaticDraughtMeters:   Float64{Value: 4.5, Valid: true},
				PersonsOnBoard:        Int64{Value: 6, Valid: true},
				Destination:           "@@@@@@@@@@@@@@@@@@@@",
				EstimatedArrivalTime:  Int64{Value: 220516, Valid: true},
				EstimatedArrivalDay:   Int64{Value: 1, Valid: true},
				EstimatedArrivalMonth: Int64{Value: 2, Valid: true},
				NavigationalStatus:    Int64{Value: 8, Valid: true},
				RegionalApplication:   Int64{Value: 0, Valid: false},
			},
		},
		{
			name: "invalid nmea: TypeOfShipAndCargo",
			raw:  "$RAVSD,x,4.5,6,@@@@@@@@@@@@@@@@@@@@,220516,01,02,8,*26",
			err:  "nmea: RAVSD invalid type of ship and cargo: x",
		},
		{
			name: "invalid nmea: StaticDraughtMeters",
			raw:  "$RAVSD,0,4.x,6,@@@@@@@@@@@@@@@@@@@@,220516,01,02,8,*23",
			err:  "nmea: RAVSD invalid maximum present static draught: 4.x",
		},
		{
			name: "invalid nmea: PersonsOnBoard",
			raw:  "$RAVSD,0,4.5,x,@@@@@@@@@@@@@@@@@@@@,220516,01,02,8,*20",
			err:  "nmea: RAVSD invalid persons on-board: x",
		},
		{
			name: "invalid nmea: EstimatedArrivalTime",
			raw:  "$RAVSD,0,4.5,6,@@@@@@@@@@@@@@@@@@@@,22051x,01,02,8,*20",
			err:  "nmea: RAVSD invalid estimated arrival time: 22051x",
		},
		{
			name: "invalid nmea: EstimatedArrivalDay",
			raw:  "$RAVSD,0,4.5,6,@@@@@@@@@@@@@@@@@@@@,220516,x1,02,8,*26",
			err:  "nmea: RAVSD invalid estimated arrival day: x1",
		},
		{
			name: "invalid nmea: EstimatedArrivalMonth",
			raw:  "$RAVSD,0,4.5,6,@@@@@@@@@@@@@@@@@@@@,220516,01,x2,8,*26",
			err:  "nmea: RAVSD invalid estimated arrival month: x2",
		},
		{
			name: "invalid nmea: NavigationalStatus",
			raw:  "$RAVSD,0,4.5,6,@@@@@@@@@@@@@@@@@@@@,220516,01,02,x,*2E",
			err:  "nmea: RAVSD invalid navigational status: x",
		},
		{
			name: "invalid nmea: RegionalApplication",
			raw:  "$RAVSD,0,4.5,6,@@@@@@@@@@@@@@@@@@@@,220516,01,02,8,x*16",
			err:  "nmea: RAVSD invalid Regional application: x",
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
				vsd := m.(VSD)
				vsd.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, vsd)
			}
		})
	}
}
