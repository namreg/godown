package clock

import "time"

//Clock is the system clock
type Clock struct{}

//Now returns current time
func (c *Clock) Now() time.Time {
	return time.Now()
}

//New creates a new clock
func New() *Clock {
	return &Clock{}
}
