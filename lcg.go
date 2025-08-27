package rng

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
	// todo fill me
}
