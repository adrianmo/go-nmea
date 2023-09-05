package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mcptests = []struct {
	name string
	raw  string
	err  string
	msg  MCP
}{
	{
		name: "good sentence",
		raw:  "$IIMCP,50.0,25.0,10.0,0,0,12345,9876,5432*72",
		msg: MCP{
			JoystickSurgeAxisCommandSetValue: 50.0,
			JoystickSwayAxisCommandSetValue:  25.0,
			JoystickYawAxisCommandSetValue:   10.0,
			Reserved1:                        0,
			Reserved2:                        0,
			ValueErrorStatusWord:             12345,
			ControlStateWord1:                9876,
			ControlStateWord2:                5432,
		},
	},
	{
		name: "bad validity",
		raw:  "$IIMCP,50.0,25.0,10.0,0,0,12345,9876,5432*64",
		err:  "nmea: sentence checksum mismatch [72 != 64]",
	},
}

func TestMCP(t *testing.T) {
	for _, tt := range mcptests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mcp := m.(MCP)
				mcp.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mcp)
			}
		})
	}
}
