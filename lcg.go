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
