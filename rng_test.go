package rng

import (
	"sync"
	"testing"
)

func TestRNG(t *testing.T) {
	testfn := func(t *testing.T, rng Interface, n int, async bool) {
		if async {
			var wg sync.WaitGroup
			for i := 0; i < n; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if rng.Int63() == 0 {
						t.Error()
					}
				}()
			}
			wg.Wait()
			return
		}
		for i := 0; i < n; i++ {
			if rng.Int63() == 0 {
				t.Error()
			}
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			t.Run("1", func(t *testing.T) { testfn(t, &KernelRandom, 1, false) })
			t.Run("10", func(t *testing.T) { testfn(t, &KernelRandom, 1, false) })
			t.Run("100", func(t *testing.T) { testfn(t, &KernelRandom, 1, false) })
			t.Run("1000", func(t *testing.T) { testfn(t, &KernelRandom, 1, false) })
		})
		t.Run("async", func(t *testing.T) {
			t.Run("1", func(t *testing.T) { testfn(t, &KernelRandom, 1, true) })
			t.Run("10", func(t *testing.T) { testfn(t, &KernelRandom, 1, true) })
			t.Run("100", func(t *testing.T) { testfn(t, &KernelRandom, 1, true) })
			t.Run("1000", func(t *testing.T) { testfn(t, &KernelRandom, 1, true) })
		})
	})
}

func BenchmarkRNG(b *testing.B) {
	b.Run("kernel/random", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = KernelRandom.Int63()
			}
		})
		b.Run("async", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if KernelRandom.Int63() == 0 {
						b.Error()
					}
				}
			})
		})
	})
}
