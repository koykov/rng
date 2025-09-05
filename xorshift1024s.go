package rng

type xorshift1024s struct {
	x     [16]uint64
	index int
}

func (r *xorshift1024s) Seed(_ int64) {}

func (r *xorshift1024s) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift1024s) Uint64() uint64 {
	index := r.index
	s := r.x[index]
	index++
	t := r.x[index]
	index &= 15
	t ^= t << 31
	t ^= t >> 11
	t ^= s ^ (s >> 30)
	r.x[index] = t
	r.index = index
	return t * 1181783497276652981
}
