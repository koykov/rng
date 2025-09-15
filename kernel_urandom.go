package rng

import (
	"math/rand"
)

const fpDevUrandom = "/dev/urandom"

var KernelUrandom = &wrapper{
	Rand:       rand.New(&kernelRandom{fp: fpDevUrandom}),
	Concurrent: &Pool{New: func() rand.Source64 { return &kernelRandom{fp: fpDevUrandom} }},
}

func NewKernelUrandom() Interface {
	return &wrapper{
		Rand:       rand.New(&kernelRandom{fp: fpDevUrandom}),
		Concurrent: &Pool{New: func() rand.Source64 { return &kernelRandom{fp: fpDevUrandom} }},
	}
}

var _ = NewKernelUrandom
