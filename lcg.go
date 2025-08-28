package rng

import (
	"math/rand"
	"sync"
)

type lcg struct {
	seed, a, c, m int64
}

func (r *lcg) Seed(v int64) {
	r.seed = v
}

func (r *lcg) Int63() int64 {
	return int64(r.Uint64())
}

func (r *lcg) Uint64() uint64 {
	r.seed = (r.a*r.seed + r.c) % r.m
	return uint64(r.seed)
}

type lcgContainer struct {
	ZXSpectrum    wrapper
	Ranqd1        wrapper
	BorlandCpp    wrapper
	BorlandDelphi wrapper
	TurboPascal   wrapper
	Glibc         wrapper
	ANSI_C        wrapper
	MSVCpp        wrapper
	MSVBasic      wrapper
	RtlUniform    wrapper
	MinstdRand    wrapper
	MinstdRand0   wrapper
	MMIX          wrapper
	Musl          wrapper
	Java          wrapper
	POSIX         wrapper
	Random0       wrapper
	Cc65          wrapper
	RANDU         wrapper
}

var LCG = &lcgContainer{
	ZXSpectrum: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 75, c: 0, m: 65537}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 75, c: 0, m: 65537}) }}},
	},
	Ranqd1: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 1664525, c: 1013904223, m: 4294967296}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 1664525, c: 1013904223, m: 4294967296}) }}},
	},
	BorlandCpp: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 22695477, c: 1, m: 2147483648}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 22695477, c: 1, m: 2147483648}) }}},
	},
	BorlandDelphi: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 134775813, c: 1, m: 4294967296}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 134775813, c: 1, m: 4294967296}) }}},
	},
	TurboPascal: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 134775813, c: 1, m: 4294967296}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 134775813, c: 1, m: 4294967296}) }}},
	},
	Glibc: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 1103515245, c: 12345, m: 2147483648}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 1103515245, c: 12345, m: 2147483648}) }}},
	},
	ANSI_C: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 1103515245, c: 12345, m: 2147483648}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 1103515245, c: 12345, m: 2147483648}) }}},
	},
	MSVCpp: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 214013, c: 2531011, m: 2147483648}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 214013, c: 2531011, m: 2147483648}) }}},
	},
	MSVBasic: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 1140671485, c: 12820163, m: 16777216}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 1140671485, c: 12820163, m: 16777216}) }}},
	},
	RtlUniform: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: -18, c: -60, m: 2147483647}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: -18, c: -60, m: 2147483647}) }}},
	},
	MinstdRand: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 48271, c: 0, m: 2147483647}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 48271, c: 0, m: 2147483647}) }}},
	},
	MinstdRand0: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 16807, c: 0, m: 2147483647}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 16807, c: 0, m: 2147483647}) }}},
	},
	MMIX: wrapper{
		Rand: rand.New(&lcg{seed: rand.Int63(), a: 6364136223846793005, c: 1442695040888963407, m: 18446744073709600000}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any {
			return rand.New(&lcg{seed: rand.Int63(), a: 6364136223846793005, c: 1442695040888963407, m: 18446744073709600000})
		}}},
	},
	Musl: wrapper{
		Rand: rand.New(&lcg{seed: rand.Int63(), a: 6364136223846793005, c: 1, m: 18446744073709600000}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any {
			return rand.New(&lcg{seed: rand.Int63(), a: 6364136223846793005, c: 1, m: 18446744073709600000})
		}}},
	},
	Java: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 25214903917, c: 11, m: 281474976710656}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 25214903917, c: 11, m: 281474976710656}) }}},
	},
	POSIX: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 25214903917, c: 11, m: 281474976710656}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 25214903917, c: 11, m: 281474976710656}) }}},
	},
	Random0: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 8121, c: 28411, m: 134456}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 8121, c: 28411, m: 134456}) }}},
	},
	Cc65: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 16843009, c: 826366247, m: 4294967296}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 16843009, c: 826366247, m: 4294967296}) }}},
	},
	RANDU: wrapper{
		Rand:       rand.New(&lcg{seed: rand.Int63(), a: 65539, c: 0, m: 2147483648}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: 65539, c: 0, m: 2147483648}) }}},
	},
}
