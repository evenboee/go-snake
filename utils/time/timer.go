package time

import "time"

// Timer takes a function and returns the time it took to run
func Timer(f func()) time.Duration {
	start := time.Now()

	f()

	end := time.Now()
	return end.Sub(start)
}
