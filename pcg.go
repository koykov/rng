package rng

import (
	"math/rand"
	"sync"
)

type pcgContainer struct {
	PCG32 wrapper
	PCG64 wrapper
}

var PCG = pcgContainer{
	PCG32: wrapper{
		Rand: rand.New(newPCG32(rand.Uint64())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any {
			return rand.New(newPCG32(rand.Uint64()))
		}}},
	},
	PCG64: wrapper{
		Rand:       rand.New(newPCG64(rand.Uint64(), rand.Uint64())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newPCG64(rand.Uint64(), rand.Uint64())) }}},
	},
}
