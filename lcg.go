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

var (
	lcgNew = func(a, c, m int64) wrapper {
		return wrapper{
			Rand:       rand.New(&lcg{seed: rand.Int63(), a: a, c: c, m: m}),
			Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lcg{seed: rand.Int63(), a: a, c: c, m: m}) }}},
		}
	}
	LCG = &lcgContainer{
		ZXSpectrum:    lcgNew(75, 0, 65537),
		Ranqd1:        lcgNew(1664525, 1013904223, 4294967296),
		BorlandCpp:    lcgNew(22695477, 1, 2147483648),
		BorlandDelphi: lcgNew(134775813, 1, 4294967296),
		TurboPascal:   lcgNew(134775813, 1, 4294967296),
		Glibc:         lcgNew(1103515245, 12345, 2147483648),
		ANSI_C:        lcgNew(1103515245, 12345, 2147483648),
		MSVCpp:        lcgNew(214013, 2531011, 2147483648),
		MSVBasic:      lcgNew(1140671485, 12820163, 16777216),
		RtlUniform:    lcgNew(-18, -60, 2147483647),
		MinstdRand:    lcgNew(48271, 0, 2147483647),
		MinstdRand0:   lcgNew(16807, 0, 2147483647),
		MMIX:          lcgNew(6364136223846793005, 1442695040888963407, 18446744073709600000),
		Musl:          lcgNew(6364136223846793005, 1, 18446744073709600000),
		Java:          lcgNew(25214903917, 11, 281474976710656),
		POSIX:         lcgNew(25214903917, 11, 281474976710656),
		Random0:       lcgNew(8121, 28411, 134456),
		Cc65:          lcgNew(16843009, 826366247, 4294967296),
		RANDU:         lcgNew(65539, 0, 2147483648),
	}
)
