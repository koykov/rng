package rng

type lcg struct {
	seed, a, c, m int64
}

func (r *lcg) Seed(v int64) {
	r.seed = v
}

func (r *lcg) Int() int {
	return int(r.Uint64())
}

func (r *lcg) Intn(n int) int {
	return int(r.Uint64() % uint64(n))
}

func (r *lcg) Int31() int32 {
	return int32(r.Uint64() >> 32)
}

func (r *lcg) Int31n(n int32) int32 {
	return int32(r.Uint64() % uint64(n))
}

func (r *lcg) Int63() int64 {
	return int64(r.Uint64())
}

func (r *lcg) Int63n(n int64) int64 {
	return int64(r.Uint64() % uint64(n))
}

func (r *lcg) Uint32() uint32 {
	return uint32(r.Uint64())
}

func (r *lcg) Uint64() uint64 {
	r.seed = (r.a*r.seed + r.c) % r.m
	return uint64(r.seed)
}
