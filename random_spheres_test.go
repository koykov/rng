package rng

import (
	"math"
	"sort"
	"strconv"
	"testing"
)

func TestRandomSpheres(t *testing.T) {
	type point struct {
		X, Y, Z float64
	}

	distanceBetween := func(p, other *point) float64 {
		dx := p.X - other.X
		dy := p.Y - other.Y
		dz := p.Z - other.Z
		return math.Sqrt(dx*dx + dy*dy + dz*dz)
	}

	genPoints := func(n int, rng Interface) []point {
		points := make([]point, 0, n)

		for len(points) < n {

			x := 2*rng.Float64() - 1
			y := 2*rng.Float64() - 1
			z := 2*rng.Float64() - 1

			if x*x+y*y+z*z <= 1.0 {
				points = append(points, point{X: x, Y: y, Z: z})
			}
		}
		return points
	}

	findMinDistance := func(points []point) []float64 {
		minDistances := make([]float64, len(points))

		for i := range points {
			minDist := math.MaxFloat64

			for j := range points {
				if i == j {
					continue
				}

				dist := distanceBetween(&points[i], &points[j])
				if dist < minDist {
					minDist = dist
				}
			}

			minDistances[i] = minDist
		}

		return minDistances
	}

	testfn := func(distances []float64, lambda float64) (float64, float64) {
		sorted := make([]float64, len(distances))
		copy(sorted, distances)
		sort.Float64s(sorted)

		n := float64(len(sorted))
		maxDiff := 0.0

		theoreticalCDF := func(x float64) float64 {
			return 1 - math.Exp(-lambda*x)
		}

		for i, x := range sorted {
			empirical := float64(i+1) / n
			theoretical := theoreticalCDF(x)
			diff := math.Abs(empirical - theoretical)

			if diff > maxDiff {
				maxDiff = diff
			}
		}

		criticalValue := 1.36 / math.Sqrt(n)

		return maxDiff, criticalValue
	}
	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				points := genPoints(step, rng)
				minDistances := findMinDistance(points)
				density := float64(step) / (4.0 * math.Pi / 3.0)
				lambda := math.Pow(4.0*math.Pi*density/3.0, 1.0/3.0)

				ksStat, criticalValue := testfn(minDistances, lambda)
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
