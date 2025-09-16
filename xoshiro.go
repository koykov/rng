package rng

import (
	"math/rand"
)

type xoshiroContainer struct {
	Xoshiro256Plus     wrapper
	Xoshiro256PlusPlus wrapper
	Xoshiro256SS       wrapper
}

var Xoshiro = xoshiroContainer{
	Xoshiro256Plus: wrapper{
		Rand: rand.New(NewXoshiro256pSource([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})),
		Concurrent: &Pool{New: func() rand.Source64 {
			return rand.New(NewXoshiro256pSource([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}))
		}},
	},
	Xoshiro256PlusPlus: wrapper{
		Rand: rand.New(NewXoshiro256ppSource([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})),
		Concurrent: &Pool{New: func() rand.Source64 {
			return rand.New(NewXoshiro256ppSource([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}))
		}},
	},
	Xoshiro256SS: wrapper{
		Rand: rand.New(NewXoshiro256ssSource([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})),
		Concurrent: &Pool{New: func() rand.Source64 {
			return rand.New(NewXoshiro256ssSource([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}))
		}},
	},
}

type xoshiroBase struct {
	s [4]uint64
}

func (r *xoshiroBase) rol64(x uint64, k int) uint64 {
	return (x << k) | (x >> (64 - k))
}
