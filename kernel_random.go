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

var KernelRandom = &wrapper{
	Rand:       rand.New(&kernelRandom{fp: fpDevRandom}),
	Concurrent: &Pool{New: func() rand.Source64 { return &kernelRandom{fp: fpDevRandom} }},
}

func NewKernelRandom() Interface {
	return &wrapper{
		Rand:       rand.New(&kernelRandom{fp: fpDevRandom}),
		Concurrent: &Pool{New: func() rand.Source64 { return &kernelRandom{fp: fpDevRandom} }},
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
	// check closed file
	if _, err = r.f.Seek(0, io.SeekStart); err != nil {
		_ = r.f.Close()
		// return empty object to open file again
		if r.f, r.err = os.Open(r.fp); r.err != nil {
			err = r.err
			return
		}
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
