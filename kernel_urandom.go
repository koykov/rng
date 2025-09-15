package rng

import (
	"math/rand"
)

const fpDevUrandom = "/dev/urandom"

func NewKernelUrandomSource() rand.Source64 {
	return &kernelRandom{fp: fpDevUrandom}
}

var KernelUrandom = &wrapper{
	Rand:       rand.New(NewKernelUrandomSource()),
	Concurrent: &Pool{New: func() rand.Source64 { return NewKernelUrandomSource() }},
}

func NewKernelUrandom() Interface {
	return &wrapper{
		Rand:       rand.New(NewKernelUrandomSource()),
		Concurrent: &Pool{New: func() rand.Source64 { return NewKernelUrandomSource() }},
	}
}

var _ = NewKernelUrandom
