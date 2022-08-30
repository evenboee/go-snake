package random

import "math/rand"

// RandomIntBetween returns a random number between min and max
// min <= n < max
func RandomIntBetween(min int, max int) int {
	return RandomInt(max-min) + min
}

// RandomInt returns a random number < max
func RandomInt(max int) int {
	return rand.Intn(max)
}
