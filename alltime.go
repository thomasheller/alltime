package alltime

import (
	"fmt"
	"time"
)

const (
	perDay   = 24
	perMonth = perDay * 30.436875
	perYear  = perMonth * 365.2425
)

// Durations formats a duration as seconds, minutes, hours, days,
// months and years.
func Durations(d time.Duration) []string {
	hours := d.Hours()
	return []string{
		// fmt.Sprintf("%dns", d.Nanoseconds()),
		fmt.Sprintf("seconds: %.2fs", d.Seconds()),
		fmt.Sprintf("minutes: %.2fm", d.Minutes()),
		fmt.Sprintf("hours:   %.2fh", hours),
		fmt.Sprintf("days:    %.2fd", hours/perDay),
		fmt.Sprintf("months:  %.2fmo", hours/perMonth),
		fmt.Sprintf("years:   %.2fy", hours/perYear),
	}
}
