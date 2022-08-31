package clients

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetWindDirection(t *testing.T) {
	client := NewOpenWeatherMap("")

	testCases := []struct {
		Degrees  int
		Expected string
	}{
		{
			Degrees:  0,
			Expected: "N",
		},
		{
			Degrees:  20,
			Expected: "NNE",
		},
		{
			Degrees:  34,
			Expected: "NE",
		},
		{
			Degrees:  70,
			Expected: "ENE",
		},
		{
			Degrees:  90,
			Expected: "E",
		},
		{
			Degrees:  120,
			Expected: "ESE",
		},
		{
			Degrees:  140,
			Expected: "SE",
		},
		{
			Degrees:  160,
			Expected: "SSE",
		},
		{
			Degrees:  180,
			Expected: "S",
		},
		{
			Degrees:  200,
			Expected: "SSW",
		},
		{
			Degrees:  220,
			Expected: "SW",
		},
		{
			Degrees:  250,
			Expected: "WSW",
		},
		{
			Degrees:  270,
			Expected: "W",
		},
		{
			Degrees:  300,
			Expected: "WNW",
		},
		{
			Degrees:  320,
			Expected: "NW",
		},
		{
			Degrees:  340,
			Expected: "NNW",
		},
	}

	for _, testCase := range testCases {
		require.Equal(t, testCase.Expected, client.getWindString(testCase.Degrees))
	}
}
