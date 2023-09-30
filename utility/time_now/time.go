package time_now

import "time"

func Wib() time.Time {
	return time.Now().UTC().Add(7 * time.Hour)
}
