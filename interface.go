package rng

type Interface interface {
	Seed(seed int64)

	Int() int
	Intn(n int) int

	Int31() int32
	Int31n(n int32) int32

	Int63() int64
	Int63n(n int64) int64

	Perm(n int) []int
	Read(p []byte) (n int, err error)
	Shuffle(n int, swap func(i, j int))

	Uint32() uint32
	Uint64() uint64

	Float32() float32
	Float64() float64
	ExpFloat64() float64
	NormFloat64() float64
}
