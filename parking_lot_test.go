package rng

import (
	"strconv"
	"testing"
)

func TestParkingLot(t *testing.T) {
	overlaps := func(x1, y1, x2, y2, lotSize, carSize float64) bool {
		return !(x1+carSize <= x2 || x2+carSize <= x1 ||
			y1+carSize <= y2 || y2+carSize <= y1)
	}
	testfn := func(rng Interface, maxAttempts int, lotSize, carSize float64) (filledArea float64, parkedCars int) {
		type point struct {
			x, y float64
		}
		lot := make(map[point]bool)
		for i := 0; i < maxAttempts; i++ {
			x := rng.Float64() * (lotSize - carSize)
			y := rng.Float64() * (lotSize - carSize)

			canPark := true
			for car := range lot {
				if overlaps(x, y, car.x, car.y, lotSize, carSize) {
					canPark = false
					break
				}
			}

			if canPark {
				lot[point{x, y}] = true
				parkedCars++
			}
		}
		filledArea = float64(parkedCars) * (carSize * carSize) / (lotSize * lotSize) * 100
		return
	}

	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				area, _ := testfn(rng, step, 100, 1)
				if area < 70.0 || area > 74.0 {
					t.Errorf("area %f outside of (70.0..74.0)", area)
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1000, 1e4, 1e5, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1000, 1e4, 1e5)
	})
}
