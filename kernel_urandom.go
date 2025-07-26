package rng

const fpDevUrandom = "/dev/urandom"

type kernelUrandomWrapper struct {
	kernelRandom
	Concurrent kernelRandomConcurrent
}

var KernelUrandom = &kernelUrandomWrapper{
	kernelRandom: kernelRandom{fp: fpDevUrandom},
	Concurrent:   kernelRandomConcurrent{fp: fpDevUrandom},
}

func NewKernelUrandom() Interface {
	return &kernelRandom{fp: fpDevUrandom}
}

var _ = NewKernelUrandom
