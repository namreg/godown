package clock

import "time"

//Clock represents clock
type Clock interface {
	//Now returns current time
	Now() time.Time
}
