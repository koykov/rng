# RNG

Collection of various sources for RNG (random number generator) and tests.

List of implemented RNGs:
* [KernelRandom](kernel_random.go) (based on `/dev/random` thus work only in Linux based OS)
* [KernelUrandom](kernel_urandom.go) (based on `/dev/urandom` thus work only in Linux based OS)
* [LCG (Linear Congruential Generator)](lcg.go)
  * ZXSpectrum
  * Ranqd1
  * BorlandCpp
  * BorlandDelphi
  * TurboPascal
  * Glibc
  * ANSI_C
  * MSVCpp
  * MSVBasic
  * RtlUniform
  * MinstdRand
  * MinstdRand0
  * MMIX
  * Musl
  * Java
  * POSIX
  * Random0
  * Cc65
  * RANDU
* [LSFR (Linear Feedback Shift Register)](lsfr.go)
  * [Fibonacci](lsfr_fibonacci.go)
  * [Galois](lsfr_galois.go)
* [Mersenne Twister](mersenne_twister.go)
  * [MT19937](mt19937.go)
  * [MT19937-64](mt19937_64.go)
* [Xorshift](xorshift.go)
  * [Xorshift32](xorshift32.go)
  * [Xorshift64](xorshift64.go)
  * [Xorshift128](xorshift128.go)
  * [Xorshift128Plus](xorshift128p.go)
  * [Xorshift1024s](xorshift1024s.go)
  * [Xorshiftr128Plus](xorshiftr128p.go)
* [Xoshiro](xoshiro.go)
  * [Xoshiro256Plus](xoshiro256p.go)
  * [Xoshiro256PlusPlus](xoshiro256pp.go)
  * [Xoshiro256SS](xoshiro256ss.go)

Each implementation has exported name to use globally, eg:
```go
import "github.com/koykov/rng"

println(rng.KernelRandom.Uint64()) // random unsigned number
println(rng.KernelRandom.Int63n(1000)) // random number in range [0..1000)
...
```

To use in multithread environments each implementation has concurrent implementation, eg:
```go
import "github.com/koykov/rng"

go func() { _ = rng.KernelRandom.Concurrent.Uint64() }
go func() { _ = rng.KernelRandom.Concurrent.Float64() }
go func() { _ = rng.KernelRandom.Concurrent.Int31n(1000) }
...
```

### Tests

List of implemented statistical tests:
* [Diehard](https://en.wikipedia.org/wiki/Diehard_tests)
  * [Birthday Spacing Test](https://www.pcg-random.org/posts/birthday-test.html)
  * [Overlapping Permutations Test](overlapping_permutations_test.go)
  * [Count-the-1’s Test](count_the_1st_test.go)
  * [Minimum Distance Test](minimum_distance_test.go)
  * [Random Spheres Test](random_spheres_test.go)
  * [The Squeeze Test](squeeze_test.go)
  * [Runs Test](runs_test.go)
  * [The Craps Test](crasp_test.go)
  * [Marsaglia and Tsang’s GCD Test](marsaglia_tsang_gcd_test.go)
* [Monobit Test](monobit_test.go)
* [Berlekamp-Massey Algorithm](berlekamp_massey_test.go)
* [Entropy Test](entropy_test.go)
* [Avalanche Test](avalanche_test.go)
* [ChiSquare Test](chi_square_test.go)
* [Kolmogorov-Smirnov Test](kolmogorov_smirnov_test.go)
* [Autocorrelation Test](autocorrelation_test.go)
* [Uniformity Test](uniformity_test.go)
