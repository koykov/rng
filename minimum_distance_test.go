package rng

import (
	"math"
	"sort"
	"strconv"
	"testing"
)

func TestMinimumDistance(t *testing.T) {
	type point []float64

	genPoints := func(n, d int, rng Interface) []point {
		points := make([]point, n)
		for i := range points {
			points[i] = make(point, d)
			for j := range points[i] {
				points[i][j] = rng.Float64()
			}
		}
		return points
	}

	euclideanDistance := func(p1, p2 point) float64 {
		sum := 0.0
		for i := range p1 {
			diff := p1[i] - p2[i]
			sum += diff * diff
		}
		return math.Sqrt(sum)
	}

	findMinDistances := func(points []point) []float64 {
		n := len(points)
		minDistances := make([]float64, n)

		for i := 0; i < n; i++ {
			minDist := math.MaxFloat64
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				dist := euclideanDistance(points[i], points[j])
				if dist < minDist {
					minDist = dist
				}
			}
			minDistances[i] = minDist
		}

		return minDistances
	}

	testfn := func(minDistances []float64, lambda float64) float64 {
		n := float64(len(minDistances))
		sorted := make([]float64, len(minDistances))
		copy(sorted, minDistances)
		sort.Float64s(sorted)

		maxDiff := 0.0
		for i, x := range sorted {
			empirical := float64(i+1) / n

			theoretical := 1 - math.Exp(-lambda*x)

			diff := math.Abs(empirical - theoretical)
			if diff > maxDiff {
				maxDiff = diff
			}
		}

		return maxDiff
	}
	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				points := genPoints(step, 3, rng)
				minDistances := findMinDistances(points)
				meanDist := 0.0
				for _, dist := range minDistances {
					meanDist += dist
				}
				meanDist /= float64(len(minDistances))
				lambda := 1.0 / meanDist

				ksStat := testfn(minDistances, lambda)
				criticalValue := 1.36 / math.Sqrt(float64(step))
				if ksStat >= criticalValue {
					t.Errorf("ksStat %f overflows critial value %f", ksStat, criticalValue)
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e4)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e4)
	})
}
