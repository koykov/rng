package rng

import "testing"

func TestRNG(t *testing.T) {
	t.Run("kernel/random", func(t *testing.T) {
		if KernelRandom.Int63() == 0 {
			t.Error()
		}
	})
}

func BenchmarkRNG(b *testing.B) {
	b.Run("kernel/random", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = KernelRandom.Int63()
		}
	})
}
