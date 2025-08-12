package rng

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestOverlappingPermutations(t *testing.T) {
	permutations := func(arr []int) [][]int {
		var helper func([]int, int)
		var res [][]int

		helper = func(arr []int, n int) {
			if n == 1 {
				tmp := make([]int, len(arr))
				copy(tmp, arr)
				res = append(res, tmp)
			} else {
				for i := 0; i < n; i++ {
					helper(arr, n-1)
					if n%2 == 1 {
						arr[0], arr[n-1] = arr[n-1], arr[0]
					} else {
						arr[i], arr[n-1] = arr[n-1], arr[i]
					}
				}
			}
		}
		helper(arr, len(arr))
		return res
	}

	generatePermutations := func(k int) [][]float64 {
		indices := make([]int, k)
		for i := 0; i < k; i++ {
			indices[i] = i
		}

		var perms [][]float64
		for _, p := range permutations(indices) {
			perm := make([]float64, k)
			for i, idx := range p {
				perm[i] = float64(idx)
			}
			perms = append(perms, perm)
		}
		return perms
	}

	equalSlices := func(a, b []float64) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	getPermutationKey := func(group []float64, perms [][]float64) string {
		sorted := make([]float64, len(group))
		copy(sorted, group)
		sort.Float64s(sorted)

		for i, perm := range perms {
			if equalSlices(group, perm) {
				return fmt.Sprintf("perm_%d", i)
			}
		}
		return "unknown"
	}

	ksTest := func(permCounts map[string]int, total int) float64 {
		sorted := make([]float64, 0, len(permCounts))
		for _, cnt := range permCounts {
			sorted = append(sorted, float64(cnt)/float64(total))
		}
		sort.Float64s(sorted)

		maxDiff := 0.0
		for i, p := range sorted {
			expected := float64(i+1) / float64(len(sorted))
			diff := math.Abs(p - expected)
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		return maxDiff
	}

	testfn := func(rng Interface, n, k int) (bool, float64) {
		raw := make([]float64, n)
		for i := 0; i < n; i++ {
			raw[i] = rng.Float64()
		}

		perms := generatePermutations(k)
		permCounts := make(map[string]int)

		for i := 0; i <= len(raw)-k; i++ {
			group := raw[i : i+k]
			key := getPermutationKey(group, perms)
			permCounts[key]++
		}

		// KS-тест вместо χ²
		ksStat := ksTest(permCounts, len(raw)-k+1)
		criticalValue := 1.36 / math.Sqrt(float64(len(perms))) // α=0.05
		isUniform := ksStat < criticalValue

		return isUniform, ksStat
	}

	testgroup := func(t *testing.T, rng Interface, n, k int) {
		t.Run("", func(t *testing.T) {
			u, c := testfn(rng, n, k)
			_ = c
			if !u {
				t.Fail()
			}
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6, 3)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6, 3)
	})
}
