package alltime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurations(t *testing.T) {
	testCases := []struct {
		duration string
		expected []string
	}{
		{"1m", []string{
			"seconds: 60.00s",
			"minutes: 1.00m",
			"hours:   0.02h",
			"days:    0.00d",
			"months:  0.00mo",
			"years:   0.00y"}},
		{"86400s", []string{
			"seconds: 86400.00s",
			"minutes: 1440.00m",
			"hours:   24.00h",
			"days:    1.00d",
			"months:  0.03mo",
			"years:   0.00y"}},
	}

	for _, tc := range testCases {
		t.Run(tc.duration, func(t *testing.T) {
			d, err := time.ParseDuration(tc.duration)
			if err != nil {
				t.Errorf("Error parsing duration: %v", err)
			}

			actual := Durations(d)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
