package rng

import (
	"math"
	"testing"
)

func TestEntropy(t *testing.T) {
	testfn := func(rng Interface, n, m int) (entropy float64) {
		raw := make([]byte, n)
		for i := 0; i < n; i++ {
			raw[i] = byte(rng.Intn(m))
		}

		counts := make([]int, m)
		for _, bit := range raw {
			counts[bit]++
		}

		total := float64(len(raw))
		for _, cnt := range counts {
			if cnt == 0 {
				continue
			}
			p := float64(cnt) / total
			entropy -= p * math.Log2(p)
		}
		return
	}
	testgroup := func(t *testing.T, rng Interface, n int) {
		t.Run("bit", func(t *testing.T) {
			entropy := testfn(rng, n, 2)
			if 1.0-entropy > 0.01 {
				t.Fail()
			}
		})
		t.Run("byte", func(t *testing.T) {
			entropy := testfn(rng, n, 256)
			if 1.0-entropy > 0.25 {
				t.Fail()
			}
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6)
	})
}
