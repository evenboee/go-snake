package random_test

import (
	"testing"

	"github.com/evenboee/go-snake/utils/random"
	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	iter := 100
	max := 3
	// number can be 0, 1, 2
	// chance of failure with correct implementation: (1/3)^100

	for i := 0; i < iter; i++ {
		require.Less(t, random.RandomInt(max), max)
	}
}

func TestRandomIntBetween(t *testing.T) {
	iter := 100
	min := 1
	max := 4
	// number can be 1, 2, 3
	// chance of failure with correct implementation: (1/3)^100

	for i := 0; i < iter; i++ {
		res := random.RandomIntBetween(min, max)
		require.Less(t, res, max)
		require.GreaterOrEqual(t, res, min)
	}
}
