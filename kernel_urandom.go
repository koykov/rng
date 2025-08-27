package rng

import (
	"math/rand"
	"sync"
)

const fpDevUrandom = "/dev/urandom"

var KernelUrandom = &kernelRandomWrapper{
	Rand:       rand.New(&kernelRandom{fp: fpDevUrandom}),
	Concurrent: &concurrent{Pool: sync.Pool{New: func() interface{} { return rand.New(&kernelRandom{fp: fpDevUrandom}) }}},
}

func NewKernelUrandom() Interface {
	return &kernelRandomWrapper{
		Rand:       rand.New(&kernelRandom{fp: fpDevUrandom}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() interface{} { return rand.New(&kernelRandom{fp: fpDevUrandom}) }}},
	}
}

var _ = NewKernelUrandom
