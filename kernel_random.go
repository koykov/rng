package rng

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"hash"
	"io"
	"math"
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
	kernelRandom
	Concurrent *kernelRandomConcurrent
}

var KernelRandom = &kernelRandomWrapper{
	kernelRandom: kernelRandom{fp: fpDevRandom},
	Concurrent:   &kernelRandomConcurrent{fp: fpDevRandom},
}

func NewKernelRandom() Interface {
	return &kernelRandomWrapper{
		kernelRandom: kernelRandom{fp: fpDevRandom},
		Concurrent:   &kernelRandomConcurrent{fp: fpDevRandom},
	}
}

var _ = NewKernelRandom

func (r *kernelRandom) Seed(seed int64) {
	r.seed = seed
	r.bseed = make([]byte, 8)
	binary.LittleEndian.PutUint64(r.bseed, uint64(r.seed))
	r.mac = hmac.New(sha256.New, r.bseed)
}

func (r *kernelRandom) Int() int { return int(uint(r.Int63()) << 1 >> 1) }

func (r *kernelRandom) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

func (r *kernelRandom) Int31() int32 { return int32(r.Int63() >> 32) }

func (r *kernelRandom) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	mx := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > mx {
		v = r.Int31()
	}
	return v % n
}

func (r *kernelRandom) Int63() int64 { return int64(r.Uint64()) }

func (r *kernelRandom) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int63() & (n - 1)
	}
	mx := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > mx {
		v = r.Int63()
	}
	return v % n
}

func (r *kernelRandom) Uint() uint { return uint(r.Uint64()) }

func (r *kernelRandom) Uint32() uint32 { return uint32(r.Int63() >> 31) }

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

func (r *kernelRandom) Float64() float64 {
again:
	f := float64(r.Int63()) / (1 << 63)
	if f == 1 {
		goto again
	}
	return f
}

func (r *kernelRandom) Float32() float32 {
again:
	f := float32(r.Float64())
	if f == 1 {
		goto again
	}
	return f
}

func (r *kernelRandom) ExpFloat64() float64 {
	for {
		j := r.Uint32()
		i := j & 0xFF
		x := float64(j) * float64(we[i])
		if j < ke[i] {
			return x
		}
		if i == 0 {
			return re - math.Log(r.Float64())
		}
		if fe[i]+float32(r.Float64())*(fe[i-1]-fe[i]) < float32(math.Exp(-x)) {
			return x
		}
	}
}

func (r *kernelRandom) NormFloat64() float64 {
	for {
		j := int32(r.Uint32())
		i := j & 0x7F
		x := float64(j) * float64(wn[i])
		abs := uint32(j)
		if j < 0 {
			abs = uint32(-j)
		}
		if abs < kn[i] {
			// This case should be hit better than 99% of the time.
			return x
		}

		if i == 0 {
			// This extra work is only required for the base strip.
			for {
				x = -math.Log(r.Float64()) * (1.0 / rn)
				y := -math.Log(r.Float64())
				if y+y >= x*x {
					break
				}
			}
			if j > 0 {
				return rn + x
			}
			return -rn - x
		}
		if fn[i]+float32(r.Float64())*(fn[i-1]-fn[i]) < float32(math.Exp(-.5*x*x)) {
			return x
		}
	}
}

func (r *kernelRandom) Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

func (r *kernelRandom) Read(p []byte) (n int, err error) {
	if r.once.Do(r.init); r.err != nil {
		return 0, r.err
	}
	return r.f.Read(p)
}

func (r *kernelRandom) Shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}

	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.int31n(int32(i + 1)))
		swap(i, j)
	}
}

func (r *kernelRandom) int31n(n int32) int32 {
	v := r.Uint32()
	prod := uint64(v) * uint64(n)
	low := uint32(prod)
	if low < uint32(n) {
		thresh := uint32(-n) % uint32(n)
		for low < thresh {
			v = r.Uint32()
			prod = uint64(v) * uint64(n)
			low = uint32(prod)
		}
	}
	return int32(prod >> 32)
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

func (r *kernelRandomConcurrent) get() *kernelRandom {
	raw := r.p.Get()
	if raw == nil {
		return &kernelRandom{fp: r.fp}
	}
	f := raw.(*os.File)
	// check closed file
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		_ = f.Close()
		// return empty object to open file again
		return &kernelRandom{fp: r.fp}
	}
	rng := &kernelRandom{fp: r.fp, f: f}
	rng.once.Do(rng.init)
	return rng
}

func (r *kernelRandomConcurrent) put(rng *kernelRandom) {
	if rng.err != nil {
		return
	}
	r.p.Put(rng.f)
}
