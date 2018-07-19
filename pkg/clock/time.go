package clock

import "time"

//TimeClock implements Clock interface
type TimeClock struct{}

//Now returns current time
func (tm TimeClock) Now() time.Time {
	return time.Now()
}
