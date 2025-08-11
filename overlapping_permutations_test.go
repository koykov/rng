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

	getChi2CriticalValue := func(df int) float64 {
		chi2Table := map[int]float64{
			1:  3.84,
			5:  11.07,
			23: 35.17,
		}
		return chi2Table[df]
	}

	testfn := func(rng Interface, n, k int) (bool, map[string]int) {
		raw := make([]float64, n)
		for i := 0; i < n; i++ {
			raw[i] = rng.Float64()
		}

		perms := generatePermutations(k)
		permCounts := make(map[string]int)

		for i := 0; i <= len(raw)-k; i++ {
			group := raw[i : i+k]
			permKey := getPermutationKey(group, perms)
			permCounts[permKey]++
		}

		expected := float64(len(raw)-k+1) / float64(len(perms))
		chi2 := 0.0
		for _, count := range permCounts {
			chi2 += math.Pow(float64(count)-expected, 2) / expected
		}

		criticalValue := getChi2CriticalValue(len(perms) - 1)
		isUniform := chi2 < criticalValue

		return isUniform, permCounts
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
		testgroup(t, KernelRandom, 1e6, 5)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6, 5)
	})
}
