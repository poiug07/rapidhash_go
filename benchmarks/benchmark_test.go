package benchmarks

import (
	"fmt"
	rand "math/rand/v2"
	"testing"

	"github.com/cespare/xxhash"
	rapidhash "github.com/poiug07/rapidhash_go"
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
	{size: 4096},
	{size: 20480},
}

var randomseed [32]byte = [32]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func generateInputs(size int) []byte {
	input := make([]byte, size)
	cc8 := rand.NewChaCha8(randomseed)
	cc8.Read(input)
	return input
}

var sink uint64

func BenchmarkRapidhash(b *testing.B) {
	for _, v := range table {
		var input = generateInputs(v.size)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			b.SetBytes(int64(v.size))
			for i := 0; i < b.N; i++ {
				sink = rapidhash.Rapidhash(input)
			}
		})
	}
}

func BenchmarkRapidhashMicro(b *testing.B) {
	for _, v := range table {
		var input = generateInputs(v.size)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			b.SetBytes(int64(v.size))
			for i := 0; i < b.N; i++ {
				sink = rapidhash.RapidhashMicro(input)
			}
		})
	}
}

func BenchmarkRapidhashNano(b *testing.B) {
	for _, v := range table {
		var input = generateInputs(v.size)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			b.SetBytes(int64(v.size))
			for i := 0; i < b.N; i++ {
				sink = rapidhash.RapidhashNano(input)
			}
		})
	}
}

func BenchmarkXXHash(b *testing.B) {
	for _, v := range table {
		var input = generateInputs(v.size)
		b.Run(fmt.Sprintf("input_size_%d", v.size), func(b *testing.B) {
			b.SetBytes(int64(v.size))
			for i := 0; i < b.N; i++ {
				sink = xxhash.Sum64(input)
			}
		})
	}
}
