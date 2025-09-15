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
		Rand:       rand.New(NewMt19937Source(rand.Int63())),
		Concurrent: &Pool{New: func() rand.Source64 { return NewMt19937Source(rand.Int63()) }},
	},
	mt19937_64: wrapper{
		Rand:       rand.New(NewMt19937_64Source(rand.Int63())),
		Concurrent: &Pool{New: func() rand.Source64 { return NewMt19937_64Source(rand.Int63()) }},
	},
}
