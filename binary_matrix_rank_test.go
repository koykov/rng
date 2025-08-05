package rng

import (
	"math"
	"testing"
)

func TestBinaryMatrixRank(t *testing.T) {
	testfn := func(rng Interface, matrixSize, numMatrices int) float64 {
		matrices := make([][][]uint64, numMatrices)
		for i := 0; i < numMatrices; i++ {
			matrix := make([][]uint64, matrixSize)
			for j := 0; j < matrixSize; j++ {
				matrix[j] = make([]uint64, matrixSize)
				for k := 0; k < matrixSize; k++ {
					matrix[j][k] = rng.Uint64()
				}
			}
			matrices[i] = matrix
		}

		rankfn := func(matrix [][]uint64) (rank int) {
			n := len(matrix)
			row := make([]uint64, n)
			copy(row, matrix[0])

			for i := 0; i < n && rank < n; i++ {
				pivotRow := -1
				for j := i; j < n; j++ {
					if matrix[j][i] != 0 {
						pivotRow = j
						break
					}
				}

				if pivotRow == -1 {
					continue
				}

				matrix[i], matrix[pivotRow] = matrix[pivotRow], matrix[i]

				for j := 0; j < n; j++ {
					if j != rank && matrix[j][i] != 0 {
						coeff := matrix[j][i]
						for k := i; k < n; k++ {
							matrix[j][k] -= coeff * matrix[rank][k]
						}
					}
				}

				rank++
			}

			return
		}
		rankCounts := make(map[int]int)
		for i := 0; i < numMatrices; i++ {
			rank := rankfn(matrices[i])
			rankCounts[rank]++
		}

		expectedProbs := map[int]float64{
			matrixSize:     0.2888, // P(rank = 32)
			matrixSize - 1: 0.5776, // P(rank = 31)
			matrixSize - 2: 0.1336, // P(rank <= 30)
		}

		observed := make([]float64, 3)
		expected := make([]float64, 3)
		for rank, count := range rankCounts {
			if rank == matrixSize {
				observed[0] += float64(count)
			} else if rank == matrixSize-1 {
				observed[1] += float64(count)
			} else {
				observed[2] += float64(count)
			}
		}

		for i, prob := range expectedProbs {
			if i == matrixSize {
				expected[0] = prob * float64(numMatrices)
			} else if i == matrixSize-1 {
				expected[1] = prob * float64(numMatrices)
			} else {
				expected[2] += prob * float64(numMatrices)
			}
		}

		var chi2 float64
		for i := 0; i < 3; i++ {
			if expected[i] != 0 {
				chi2 += math.Pow(observed[i]-expected[i], 2) / expected[i]
			}
		}

		pValue := math.Exp(-chi2 / 2)
		return pValue
	}
	testgroup := func(t *testing.T, rng Interface, matrixSize, numMatrices int) {
		t.Run("", func(t *testing.T) {
			p := testfn(rng, matrixSize, numMatrices)
			if p < 0.05 {
				t.Fail()
			}
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 32, 100)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 32, 100)
	})
}
