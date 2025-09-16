package rng

import (
	"math/rand"
)

type pcgContainer struct {
	PCG32 wrapper
	PCG64 wrapper
}

var PCG = pcgContainer{
	PCG32: wrapper{
		Rand: rand.New(NewPCG32Source(rand.Uint64())),
		Concurrent: &Pool{New: func() rand.Source64 {
			return rand.New(NewPCG32Source(rand.Uint64()))
		}},
	},
	PCG64: wrapper{
		Rand:       rand.New(NewPCG64Source(rand.Uint64(), rand.Uint64())),
		Concurrent: &Pool{New: func() rand.Source64 { return NewPCG64Source(rand.Uint64(), rand.Uint64()) }},
	},
}
