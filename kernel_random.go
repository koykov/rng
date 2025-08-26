package rng

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"hash"
	"io"
	"math/rand"
	"os"
	"sync"
)

const fpDevRandom = "/dev/random"

type kernelRandom struct {
	fp        string
	f         *os.File
	seed      int64
	bseed     []byte
	buf, hbuf []byte
	mac       hash.Hash
	err       error
	once      sync.Once
}

type kernelRandomWrapper struct {
	*rand.Rand
	Concurrent *kernelRandomConcurrent
}

var KernelRandom = &kernelRandomWrapper{
	Rand:       rand.New(&kernelRandom{fp: fpDevRandom}),
	Concurrent: &kernelRandomConcurrent{fp: fpDevRandom},
}

func NewKernelRandom() Interface {
	return &kernelRandomWrapper{
		Rand:       rand.New(&kernelRandom{fp: fpDevRandom}),
		Concurrent: &kernelRandomConcurrent{fp: fpDevRandom},
	}
}

var _ = NewKernelRandom

func (r *kernelRandom) Seed(seed int64) {
	r.seed = seed
	r.bseed = make([]byte, 8)
	binary.LittleEndian.PutUint64(r.bseed, uint64(r.seed))
	r.mac = hmac.New(sha256.New, r.bseed)
}

func (r *kernelRandom) Int63() int64 { return int64(r.Uint64()) }

func (r *kernelRandom) Uint64() uint64 {
	if _, err := r.Read(r.buf); err != nil {
		return 0
	}
	r.mac.Reset()
	r.mac.Write(r.bseed)
	r.mac.Write(r.buf)
	r.hbuf = r.mac.Sum(r.hbuf[:0])
	return binary.LittleEndian.Uint64(r.hbuf[:8])
}

func (r *kernelRandom) Read(p []byte) (n int, err error) {
	if r.once.Do(r.init); r.err != nil {
		return 0, r.err
	}
	return r.f.Read(p)
}

func (r *kernelRandom) init() {
	if r.f, r.err = os.Open(r.fp); r.err != nil {
		return
	}
	r.buf = make([]byte, 8)
	r.Seed(r.seed)
}

// concurrent stuff

type kernelRandomConcurrent struct {
	fp string
	p  sync.Pool
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

func (r *kernelRandomConcurrent) get() *rand.Rand {
	raw := r.p.Get()
	if raw == nil {
		return rand.New(&kernelRandom{fp: r.fp})
	}
	f := raw.(*os.File)
	// check closed file
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		_ = f.Close()
		// return empty object to open file again
		return rand.New(&kernelRandom{fp: r.fp})
	}
	rng := rand.New(&kernelRandom{fp: r.fp, f: f})
	return rng
}

func (r *kernelRandomConcurrent) put(rng *rand.Rand) {
	if rng == nil {
		return
	}
	r.p.Put(rng)
}
