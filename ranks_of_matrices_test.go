package rng

import (
	"strconv"
	"testing"
)

func TestRanksOfMatrices(t *testing.T) {
	generateBinaryMatrix := func(rng Interface, sz int) [][]uint64 {
		matrix := make([][]uint64, sz)
		for i := range matrix {
			matrix[i] = make([]uint64, (sz+63)/64)
			for j := 0; j < len(matrix[i]); j++ {
				num1 := rng.Uint64()
				num2 := rng.Uint64()
				matrix[i][j] = (num1 << 32) | num2
			}
		}
		return matrix
	}

	matrixRank := func(matrix [][]uint64) int {
		rows := len(matrix)
		if rows == 0 {
			return 0
		}
		cols := len(matrix[0]) * 64
		rank := 0

		for col := 0; col < cols && rank < rows; col++ {
			pivot := rank
			for pivot < rows && (matrix[pivot][col/64]>>(63-col%64))&1 == 0 {
				pivot++
			}
			if pivot == rows {
				continue
			}

			matrix[rank], matrix[pivot] = matrix[pivot], matrix[rank]

			for i := 0; i < rows; i++ {
				if i != rank && (matrix[i][col/64]>>(63-col%64))&1 == 1 {
					for j := 0; j < len(matrix[i]); j++ {
						matrix[i][j] ^= matrix[rank][j]
					}
				}
			}
			rank++
		}
		return rank
	}

	testfn := func(rng Interface, sz, n int) float64 {
		rankStats := make(map[int]int)
		for i := 0; i < n; i++ {
			matrix := generateBinaryMatrix(rng, sz)
			rank := matrixRank(matrix)
			rankStats[rank]++
		}
		return float64(rankStats[sz]) / float64(n)
	}
	testgroup := func(t *testing.T, rng Interface, sz, n int) {
		t.Run(strconv.Itoa(sz), func(t *testing.T) {
			r := testfn(rng, sz, n)
			t.Log(r)
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
