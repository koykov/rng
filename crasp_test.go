package rng

import (
	"math"
	"strconv"
	"testing"
)

func TestCrasp(t *testing.T) {
	testfn := func(rng Interface, games int) float64 {
		wins := 0
		roll := func() int {
			dice1 := int(math.Abs(rng.Float64())*6) + 1
			dice2 := int(math.Abs(rng.Float64())*6) + 1
			return dice1 + dice2
		}

		for i := 0; i < games; i++ {
			firstRoll := roll()
			switch firstRoll {
			case 7, 11:
				wins++
			case 2, 3, 12:
			default:
				point := firstRoll
				for {
					newRoll := roll()
					if newRoll == point {
						wins++
						break
					} else if newRoll == 7 {
						break
					}
				}
			}
		}

		expectedWinProb := 244.0 / 495.0
		expectedWins := expectedWinProb * float64(games)
		chi2 := math.Pow(float64(wins)-expectedWins, 2) / expectedWins
		pValue := math.Exp(-chi2 / 2)
		return pValue
	}
	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				pValue := testfn(rng, step)
				if pValue <= 0.05 {
					t.Fail()
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6)
	})
}
