package rng

import "math/rand"

type xorshift1024s struct {
	x     [16]uint64
	index int
}

func NewXorshift1024sSource(seed int64) rand.Source64 {
	r := &xorshift1024s{}
	r.Seed(seed)
	return r
}

func (r *xorshift1024s) Seed(seed int64) {
	r.x[0] = uint64(seed)
}

func (r *xorshift1024s) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift1024s) Uint64() uint64 {
	index := r.index
	s := r.x[index]
	index++
	index &= 15
	t := r.x[index]
	t ^= t << 31
	t ^= t >> 11
	t ^= s ^ (s >> 30)
	r.x[index] = t
	r.index = index
	return t * 1181783497276652981
}
