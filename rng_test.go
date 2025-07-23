package rng

import "testing"

func TestRNG(t *testing.T) {
	t.Run("kernel/random", func(t *testing.T) {
		var r KernelRandom
		if r.Int63() == 0 {
			t.Error()
		}
	})
}
