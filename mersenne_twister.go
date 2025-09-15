package rng

import (
	"math/rand"
)

type mtContainer struct {
	mt19937    wrapper
	mt19937_64 wrapper
}

var MersenneTwister = &mtContainer{
	mt19937: wrapper{
		Rand:       rand.New(newMt19937(rand.Int63())),
		Concurrent: &Pool{New: func() rand.Source64 { return newMt19937(rand.Int63()) }},
	},
	mt19937_64: wrapper{
		Rand:       rand.New(newMt19937_64(rand.Int63())),
		Concurrent: &Pool{New: func() rand.Source64 { return newMt19937_64(rand.Int63()) }},
	},
}
