package rng

import (
	"math/rand"
)

type xorshiftContainer struct {
	Xorshift32       wrapper
	Xorshift64       wrapper
	Xorshift128      wrapper
	Xorshift128Plus  wrapper
	Xorshift1024s    wrapper
	Xorshiftr128Plus wrapper
}

var Xorshift = xorshiftContainer{
	Xorshift32: wrapper{
		Rand:       rand.New(NewXorshift32Source(rand.Int63())),
		Concurrent: &Pool{New: func() rand.Source64 { return rand.New(NewXorshift32Source(rand.Int63())) }},
	},
	Xorshift64: wrapper{
		Rand:       rand.New(NewXorshift64Source(rand.Int63())),
		Concurrent: &Pool{New: func() rand.Source64 { return rand.New(NewXorshift64Source(rand.Int63())) }},
	},
	Xorshift128: wrapper{
		Rand:       rand.New(NewXorshift128Source(rand.Uint64())),
		Concurrent: &Pool{New: func() rand.Source64 { return rand.New(NewXorshift128Source(rand.Uint64())) }},
	},
	Xorshift128Plus: wrapper{
		Rand:       rand.New(NewXorshift128pSource(rand.Uint64())),
		Concurrent: &Pool{New: func() rand.Source64 { return rand.New(NewXorshift128pSource(rand.Uint64())) }},
	},
	Xorshift1024s: wrapper{
		Rand:       rand.New(NewXorshift1024sSource(rand.Int63())),
		Concurrent: &Pool{New: func() rand.Source64 { return rand.New(NewXorshift1024sSource(rand.Int63())) }},
	},
	Xorshiftr128Plus: wrapper{
		Rand:       rand.New(NewXorshiftr128pSource(rand.Uint64())),
		Concurrent: &Pool{New: func() rand.Source64 { return rand.New(NewXorshiftr128pSource(rand.Uint64())) }},
	},
}
