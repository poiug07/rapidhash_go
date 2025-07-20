package rapidhash

import (
	"fmt"
	rand "math/rand/v2"
	"testing"
)

var table = []struct {
	size int
}{
	{size: 0},
	{size: 3},
	{size: 4},
	{size: 5},
	{size: 16},
	{size: 24},
	{size: 32},
	{size: 64},
	{size: 127},
	{size: 128},
	{size: 129},
	{size: 256},
	{size: 2048},
}

var randomseed [32]byte = [32]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func BenchmarkRapidhash(b *testing.B) {
	for _, v := range table {
		input := make([]byte, v.size)
		cc8 := rand.NewChaCha8(randomseed)
		cc8.Read(input)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Rapidhash(input)
			}
		})
	}
}

func BenchmarkRapidhashMicro(b *testing.B) {
	for _, v := range table {
		input := make([]byte, v.size)
		cc8 := rand.NewChaCha8(randomseed)
		cc8.Read(input)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RapidhashMicro(input)
			}
		})
	}
}

func BenchmarkRapidhashNano(b *testing.B) {
	for _, v := range table {
		input := make([]byte, v.size)
		cc8 := rand.NewChaCha8(randomseed)
		cc8.Read(input)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RapidhashNano(input)
			}
		})
	}
}

func BenchmarkXXHash(b *testing.B) {
	for _, v := range table {
		input := make([]byte, v.size)
		cc8 := rand.NewChaCha8(randomseed)
		cc8.Read(input)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// xxhash.Sum64(input)
			}
		})
	}
}
