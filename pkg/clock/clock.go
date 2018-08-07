package clock

import "time"

//go:generate minimock -i github.com/namreg/godown-v2/pkg/clock.Clock -o ./ -s "_mock.go" -b test

//Clock represents clock
type Clock interface {
	//Now returns current time
	Now() time.Time
}
