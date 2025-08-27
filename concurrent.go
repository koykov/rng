package rng

import (
	"math/rand"
	"sync"
)

type concurrent struct {
	sync.Pool
}

func (r *concurrent) Seed(_ int64) {}

func (r *concurrent) Int() (x int) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int()
	return
}

func (r *concurrent) Intn(n int) (x int) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Intn(n)
	return
}

func (r *concurrent) Int31() (x int32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int31()
	return
}

func (r *concurrent) Int31n(n int32) (x int32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int31n(n)
	return
}

func (r *concurrent) Int63() (x int64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int63()
	return
}

func (r *concurrent) Int63n(n int64) (x int64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int63n(n)
	return
}

func (r *concurrent) Perm(n int) (x []int) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Perm(n)
	return
}

func (r *concurrent) Read(p []byte) (n int, err error) {
	rng := r.get()
	defer r.put(rng)
	n, err = rng.Read(p)
	return
}

func (r *concurrent) Shuffle(n int, swap func(i, j int)) {
	rng := r.get()
	defer r.put(rng)
	rng.Shuffle(n, swap)
}

func (r *concurrent) Uint32() (x uint32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Uint32()
	return
}

func (r *concurrent) Uint64() (x uint64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Uint64()
	return
}

func (r *concurrent) Float32() (x float32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Float32()
	return
}

func (r *concurrent) Float64() (x float64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Float64()
	return
}

func (r *concurrent) ExpFloat64() (x float64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.ExpFloat64()
	return
}

func (r *concurrent) NormFloat64() (x float64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.NormFloat64()
	return
}

func (r *concurrent) get() *rand.Rand {
	raw := r.Get()
	rng := raw.(*rand.Rand)
	return rng
}

func (r *concurrent) put(rng *rand.Rand) {
	if rng == nil {
		return
	}
	r.Put(rng)
}
