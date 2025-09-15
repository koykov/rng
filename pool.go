package rng

import (
	"math/rand"
	"sync"
	"time"
)

type Pool struct {
	New  func() rand.Source64
	p    sync.Pool
	once sync.Once
}

func (r *Pool) Seed(_ int64) {}

func (r *Pool) Int() (x int) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Int()
	return
}

func (r *Pool) Intn(n int) (x int) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Intn(n)
	return
}

func (r *Pool) Int31() (x int32) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Int31()
	return
}

func (r *Pool) Int31n(n int32) (x int32) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Int31n(n)
	return
}

func (r *Pool) Int63() (x int64) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Int63()
	return
}

func (r *Pool) Int63n(n int64) (x int64) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Int63n(n)
	return
}

func (r *Pool) Perm(n int) (x []int) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Perm(n)
	return
}

func (r *Pool) Read(p []byte) (n int, err error) {
	rng := r.Acquire()
	defer r.Release(rng)
	n, err = rng.Read(p)
	return
}

func (r *Pool) Shuffle(n int, swap func(i, j int)) {
	rng := r.Acquire()
	defer r.Release(rng)
	rng.Shuffle(n, swap)
}

func (r *Pool) Uint32() (x uint32) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Uint32()
	return
}

func (r *Pool) Uint64() (x uint64) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Uint64()
	return
}

func (r *Pool) Float32() (x float32) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Float32()
	return
}

func (r *Pool) Float64() (x float64) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.Float64()
	return
}

func (r *Pool) ExpFloat64() (x float64) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.ExpFloat64()
	return
}

func (r *Pool) NormFloat64() (x float64) {
	rng := r.Acquire()
	defer r.Release(rng)
	x = rng.NormFloat64()
	return
}

func (r *Pool) Acquire() *rand.Rand {
	r.once.Do(r.init)
	raw := r.p.Get()
	rng := raw.(*rand.Rand)
	return rng
}

func (r *Pool) Release(rng *rand.Rand) {
	r.once.Do(r.init)
	if rng == nil {
		return
	}
	r.p.Put(rng)
}

func (r *Pool) init() {
	if r.New == nil {
		r.New = func() rand.Source64 {
			src := rand.NewSource(time.Now().UnixNano())
			return any(src).(rand.Source64)
		}
	}
	r.p.New = func() any { return rand.New(r.New()) }
}
