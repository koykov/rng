package rng

const fpDevUrandom = "/dev/urandom"

type kernelUrandom = kernelRandom

type kernelUrandomWrapper struct {
	kernelUrandom
}

var KernelUrandom = &kernelUrandomWrapper{kernelUrandom: kernelUrandom{fp: fpDevUrandom}}

func NewKernelUrandom() Interface {
	return &kernelUrandom{fp: fpDevUrandom}
}

var _ = NewKernelUrandom
