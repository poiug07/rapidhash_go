package rapidhash

import (
	"encoding/binary"
	"math/bits"
)

/*
64*64 -> 128bit multiply function.

Calculates 128-bit C = *A * *B.
*/
func rapid_mum(A *uint64, B *uint64) {
	hi, lo := bits.Mul64(*A, *B)
	*A = lo
	*B = hi
}

func rapid_mum_protected(A *uint64, B *uint64) {
	hi, lo := bits.Mul64(*A, *B)
	*A ^= lo
	*B ^= hi
}

func rapid_mix(A, B uint64) uint64 {
	if protected {
		rapid_mum_protected(&A, &B)
	} else {
		rapid_mum(&A, &B)
	}
	return A ^ B
}

func read64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)
}

func read32(data []byte) uint64 {
	return uint64(binary.LittleEndian.Uint32(data))
}
