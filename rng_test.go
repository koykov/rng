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

func BenchmarkRNG(b *testing.B) {
	b.Run("/dev/random", func(b *testing.B) {
		var r KernelRandom
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = r.Int63()
		}
	})
}
