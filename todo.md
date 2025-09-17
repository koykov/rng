## 1. True Random Number Generators (TRNG)
├── **Hardware-based**
│   ├── Thermal noise (resistors, CPUs)
│   ├── Photonic noise (lasers, semiconductors)
│   ├── Radioactive decay (Geiger counters)
│   ├── Avalanche noise (Zener diodes)
│   ├── Atmospheric noise (radio receivers)
│   └── Clock jitter (signal instability)
├── **Quantum-based**
│   ├── Photon polarization (ID Quantique QRNG)
│   └── Quantum optics (NIST standards)
└── **Other physical sources**
├── User input (mouse/keyboard timing)
├── HDD/fan speed variations
└── Microphone/webcam sensor noise

## 2. Pseudorandom Number Generators (PRNG)
├── **Linear methods**
│   ├── Linear Congruential Generator (LCG) +
│   │   ├── ANSI C `rand()` +
│   │   ├── Java `java.util.Random` (legacy) +
│   │   └── Park-Miller MINSTD +
│   └── Linear Feedback Shift Register (LFSR) +
│       ├── Fibonacci LFSR +
│       └── Galois LFSR +
├── **Modern non-cryptographic**
│   ├── Mersenne Twister +
│   │   ├── MT19937 (Python/R default) +
│   │   └── MT19937-64 (64-bit) +
│   ├── Xorshift +
│   │   ├── Xorshift32 +
│   │   ├── Xorshift64 +
│   │   └── Xorshift128+ +
│   └── PCG (Permuted Congruential Generator) +
│   |   ├── PCG32 +
│   |   └── PCG64 +
|   └── xoshiro* +
└── **Cryptographically secure**
├── Block cipher-based
│   ├── AES-CTR (NIST SP800-90A)
│   └── ChaCha20 (Linux `/dev/urandom`)
├── Hash-based
│   ├── HMAC-DRBG (NIST)
│   └── SHAKE (SHA-3)
└── Hybrid/Standards
├── Fortuna (Yarrow successor)
├── Yarrow (deprecated)
└── NIST SP800-90A
├── Hash_DRBG
├── HMAC_DRBG
└── CTR_DRBG

## 3. Hybrid Systems (TRNG + PRNG)
├── **Linux Kernel RNG**
│   ├── `/dev/random` (blocking)
│   └── `/dev/urandom` (non-blocking)
├── **Windows RNG**
│   ├── `CryptGenRandom` (deprecated)
│   └── `BCryptGenRandom` (CNG)
└── **Other OS/Libraries**
├── OpenSSL `RAND_bytes()`
└── Apple `SecRandomCopyBytes()`

## 4. Specialized Subcategories
├── **GPU-optimized**
│   ├── NVIDIA CuRAND
│   └── AMD ROCm-RAND
└── **Quantum-simulated PRNG**
└── Quantum algorithm emulators (software-only)

## Tests
├── Monobit Test +
├── Berlekamp-Massey Algorithm +
├── Entropy Test +
├── Avalanche Test +
├── Batteries
│   ├── Diehard Tests (Marsaglia)
|   │   ├── Birthday Spacing Test +
|   │   ├── Overlapping Permutations Test +
|   │   ├── Ranks of 31×31 and 32×32 Matrices -
|   │   ├── Ranks of 6×8 Matrices -
|   │   ├── Monkey Tests
|   │   |   ├── Count-the-1’s Test +
|   │   |   └── Parking Lot Test -
|   │   ├── Minimum Distance Test +
|   │   ├── Random Spheres Test +
|   │   ├── The Squeeze Test +
|   │   ├── Overlapping Sums Test -
|   │   ├── Runs Test +
|   │   ├── The Craps Test +
|   │   └── Marsaglia and Tsang’s GCD Test +
│   ├── TestU01 (L’Ecuyer)
|   │   ├── SmallCrush
|   │   ├── Crush
|   │   └── BigCrush
│   └── NIST STS
├── Spectral Test
├── Ziggurat Test
└── Collision Test
