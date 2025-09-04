package rng

import (
	"math/rand"
	"sync"
)

type mtContainer struct {
	mt19937    wrapper
	mt19937_64 wrapper
}

var MersenneTwister = &mtContainer{
	mt19937: wrapper{
		Rand:       rand.New(newMt19937(rand.Int63())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newMt19937(rand.Int63())) }}},
	},
	mt19937_64: wrapper{
		Rand:       rand.New(&mt19937_64{seed: rand.Uint64()}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&mt19937_64{seed: rand.Uint64()}) }}},
	},
}
