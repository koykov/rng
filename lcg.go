package rng

type lcg struct {
	seed, a, c, m uint64
}

func (r *lcg) Uint64() uint64 {
	r.seed = (r.a*r.seed + r.c) % r.m
	return r.seed
}
