package rng

import (
	"math"
	"strconv"
	"testing"
)

func TestRanksOfMatrices(t *testing.T) {
	computeMatrixRank := func(matrix [][]float64) int {
		q := len(matrix)
		rank := 0
		rowSelected := make([]bool, q)

		for col := 0; col < q; col++ {
			var pivotRow int
			for pivotRow = 0; pivotRow < q; pivotRow++ {
				if !rowSelected[pivotRow] && matrix[pivotRow][col] != 0 {
					break
				}
			}

			if pivotRow < q {
				rank++
				rowSelected[pivotRow] = true

				pivotVal := matrix[pivotRow][col]
				for j := col; j < q; j++ {
					matrix[pivotRow][j] /= pivotVal
				}

				for i := 0; i < q; i++ {
					if i != pivotRow && matrix[i][col] != 0 {
						scale := matrix[i][col]
						for j := col; j < q; j++ {
							matrix[i][j] -= scale * matrix[pivotRow][j]
						}
					}
				}
			}
		}
		return rank
	}

	testfn := func(rng Interface, sz, q int) float64 {
		n := q * q * 10
		binData := make([]float64, n)
		for i := range binData {
			if rng.Float64() > 0.5 {
				binData[i] = 1.0
			}
		}

		blockSize := q * q
		numM := n / blockSize

		if numM > 0 {
			maxRanks := [3]int{}

			for im := 0; im < numM; im++ {
				matrix := make([][]float64, q)
				for i := 0; i < q; i++ {
					matrix[i] = make([]float64, q)
					for j := 0; j < q; j++ {
						idx := im*blockSize + i*q + j
						if idx < len(binData) {
							matrix[i][j] = binData[idx]
						}
					}
				}

				rank := computeMatrixRank(matrix)

				switch {
				case rank == q:
					maxRanks[0]++
				case rank == q-1:
					maxRanks[1]++
				default:
					maxRanks[2]++
				}
			}

			piks := [3]float64{1.0, 0.0, 0.0}
			for x := 1; x < 50; x++ {
				piks[0] *= 1.0 - 1.0/math.Pow(2.0, float64(x))
			}
			piks[1] = 2 * piks[0]
			piks[2] = 1 - piks[0] - piks[1]

			chi := 0.0
			for i := 0; i < 3; i++ {
				expected := piks[i] * float64(numM)
				chi += math.Pow(float64(maxRanks[i])-expected, 2.0) / expected
			}

			return math.Exp(-chi / 2.0)
		}

		return -1.0
	}
	testgroup := func(t *testing.T, rng Interface, sz, n int) {
		t.Run(strconv.Itoa(sz), func(t *testing.T) {
			r := testfn(rng, sz, n)
			t.Log(r)
			if 1.0-r > 0.1 {
				t.Fail()
			}
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 31, 1000)
		testgroup(t, KernelRandom, 32, 1000)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 31, 1000)
		testgroup(t, KernelUrandom, 32, 1000)
	})
}
