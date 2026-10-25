package rapidhash

import (
	"encoding/binary"
	"math/bits"
)

/*
64*64 -> 128bit multiply function.

Calculates 128-bit C = A * B, returning the high and low 64-bit halves.
*/
func rapid_mum(a, b uint64) (hi, lo uint64) {
	return bits.Mul64(a, b)
}

func rapid_mum_protected(a, b uint64) (hi, lo uint64) {
	hi, lo = bits.Mul64(a, b)
	return hi ^ a, lo ^ b
}

func read64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)
}

func read32(data []byte) uint64 {
	return uint64(binary.LittleEndian.Uint32(data))
}
