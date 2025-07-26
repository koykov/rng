package rng

import (
	"io"
	"os"
	"sync"
)

type kernelRandomConcurrent struct {
	p sync.Pool
}

func (r *kernelRandomConcurrent) Seed(_ int64) {}

func (r *kernelRandomConcurrent) Int() (x int) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int()
	return
}

func (r *kernelRandomConcurrent) Intn(n int) (x int) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Intn(n)
	return
}

func (r *kernelRandomConcurrent) Int31() (x int32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int31()
	return
}

func (r *kernelRandomConcurrent) Int31n(n int32) (x int32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int31n(n)
	return
}

func (r *kernelRandomConcurrent) Int63() (x int64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int63()
	return
}

func (r *kernelRandomConcurrent) Int63n(n int64) (x int64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Int63n(n)
	return
}

func (r *kernelRandomConcurrent) Perm(n int) (x []int) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Perm(n)
	return
}

func (r *kernelRandomConcurrent) Read(p []byte) (n int, err error) {
	rng := r.get()
	defer r.put(rng)
	n, err = rng.Read(p)
	return
}

func (r *kernelRandomConcurrent) Shuffle(n int, swap func(i, j int)) {
	rng := r.get()
	defer r.put(rng)
	rng.Shuffle(n, swap)
}

func (r *kernelRandomConcurrent) Uint32() (x uint32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Uint32()
	return
}

func (r *kernelRandomConcurrent) Uint64() (x uint64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Uint64()
	return
}

func (r *kernelRandomConcurrent) Float32() (x float32) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Float32()
	return
}

func (r *kernelRandomConcurrent) Float64() (x float64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.Float64()
	return
}

func (r *kernelRandomConcurrent) ExpFloat64() (x float64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.ExpFloat64()
	return
}

func (r *kernelRandomConcurrent) NormFloat64() (x float64) {
	rng := r.get()
	defer r.put(rng)
	x = rng.NormFloat64()
	return
}

func (r *kernelRandomConcurrent) get() *kernelRandom {
	raw := r.p.Get()
	if raw == nil {
		return &kernelRandom{fp: fpDevRandom}
	}
	f := raw.(*os.File)
	// check closed file
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		_ = f.Close()
		// return empty object to open file again
		return &kernelRandom{fp: fpDevRandom}
	}
	rng := &kernelRandom{fp: fpDevRandom, f: f}
	rng.once.Do(func() {})
	return rng
}

func (r *kernelRandomConcurrent) put(rng *kernelRandom) {
	if rng.err != nil {
		return
	}
	r.p.Put(rng.f)
}
