package rng

import (
	"math/rand"
	"sync"
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
		Rand:       rand.New(newXorshift32(rand.Int63())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newXorshift32(rand.Int63())) }}},
	},
	Xorshift64: wrapper{
		Rand:       rand.New(newXorshift64(rand.Int63())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newXorshift64(rand.Int63())) }}},
	},
	Xorshift128: wrapper{
		Rand:       rand.New(newXorshift128(rand.Uint64())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newXorshift128(rand.Uint64())) }}},
	},
	Xorshift128Plus: wrapper{
		Rand:       rand.New(newXorshift128p(rand.Uint64())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newXorshift128p(rand.Uint64())) }}},
	},
	Xorshift1024s: wrapper{
		Rand:       rand.New(newXorshift1024s(rand.Int63())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newXorshift1024s(rand.Int63())) }}},
	},
	Xorshiftr128Plus: wrapper{
		Rand:       rand.New(newXorshiftr128p(rand.Uint64())),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(newXorshiftr128p(rand.Uint64())) }}},
	},
}
