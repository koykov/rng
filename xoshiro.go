package rng

import (
	"math/rand"
	"sync"
)

type xoshiroContainer struct {
	Xoshiro256Plus     wrapper
	Xoshiro256PlusPlus wrapper
	Xoshiro256SS       wrapper
}

var Xoshiro = xoshiroContainer{
	Xoshiro256Plus: wrapper{
		Rand: rand.New(newXoshiro256p([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any {
			return rand.New(newXoshiro256p([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}))
		}}},
	},
	Xoshiro256PlusPlus: wrapper{
		Rand: rand.New(newXoshiro256pp([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any {
			return rand.New(newXoshiro256pp([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}))
		}}},
	},
	Xoshiro256SS: wrapper{
		Rand: rand.New(newXoshiro256ss([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any {
			return rand.New(newXoshiro256ss([4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}))
		}}},
	},
}

type xoshiroBase struct {
	s [4]uint64
}

func (r *xoshiroBase) rol64(x uint64, k int) uint64 {
	return (x << k) | (x >> (64 - k))
}
