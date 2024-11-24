package rapidhash

import (
	"encoding/binary"
	"math/bits"
)

// Default seed
const rapid_seed uint64 = 0xBDD89AA982704029

// Default secret parameters
var rapid_secret [3]uint64 = [3]uint64{0x2D358DCCAA6C78A5, 0x8BB84B93962EACC9, 0x4B33A62ED433D4A3}

/*
64*64 -> 128bit multiply function.

@param A  Address of 64-bit number.
@param B  Address of 64-bit number.

Calculates 128-bit C = *A * *B.
*/
func rapid_mum(A *uint64, B *uint64) {
	hi, lo := bits.Mul64(*A, *B)
	*A = lo
	*B = hi
	// Below and above are equivalent, didn't know go has 128 bit multiply
	/* var ha uint64 = *A >> 32
	var hb uint64 = *B >> 32

	var la uint64 = uint64(*(*uint32)(unsafe.Pointer(A)))
	var lb uint64 = uint64(*(*uint32)(unsafe.Pointer(B)))
	var hi uint64
	var lo uint64
	var rh uint64 = ha * hb
	var rm0 uint64 = ha * lb
	var rm1 uint64 = hb * la
	var rl uint64 = la * lb
	var t uint64 = rl + (rm0 << 32)
	var c uint64 = 0
	if t < rl {
		c = 1
	}
	lo = t + (rm1 << 32)
	if lo < t {
		c += 1
	}
	hi = rh + (rm0 >> 32) + (rm1 >> 32) + c
	*A = lo
	*B = hi */
}

/*
 *  Multiply and xor mix function.
 *
 *  @param A  64-bit number.
 *  @param B  64-bit number.
 *
 *  Calculates 128-bit C = A * B.
 *  Returns 64-bit xor between high and low 64 bits of C.
 */
func rapid_mix(A, B uint64) uint64 {
	rapid_mum(&A, &B)
	return A ^ B
}

func readSmall(data []byte, bufferlen uint64) uint64 {
	return (uint64(data[0]) << 56) | (uint64(data[bufferlen>>1]) << 32) | uint64(data[bufferlen-1])
}

func rapidhash_internal(data []byte, seed uint64, secret [3]uint64) uint64 {
	bufferlen := uint64(len(data))
	seed ^= rapid_mix(seed^secret[0], secret[1]) ^ bufferlen
	var a, b uint64
	if bufferlen <= 16 {
		if bufferlen >= 4 {
			a = uint64(binary.LittleEndian.Uint32(data))<<32 |
				uint64(binary.LittleEndian.Uint32(data[bufferlen-4:bufferlen]))
			var delta uint64 = ((bufferlen & 24) >> (bufferlen >> 3))
			b = uint64(binary.LittleEndian.Uint32(data[delta:]))<<32 |
				uint64(binary.LittleEndian.Uint32(data[bufferlen-4-delta:]))
		} else {
			if bufferlen > 0 {
				a = readSmall(data, bufferlen)
				b = 0
			} else {
				a = 0
				b = 0
			}
		}
	} else {
		datapos := uint64(0)
		i := bufferlen
		if i > 48 {
			var see1 uint64 = seed
			var see2 uint64 = seed
			for i >= 48 {
				seed = rapid_mix(binary.LittleEndian.Uint64(data[datapos:])^secret[0], binary.LittleEndian.Uint64(data[datapos+8:])^seed)
				see1 = rapid_mix(binary.LittleEndian.Uint64(data[datapos+16:])^secret[1], binary.LittleEndian.Uint64(data[datapos+24:])^see1)
				see2 = rapid_mix(binary.LittleEndian.Uint64(data[datapos+32:])^secret[2], binary.LittleEndian.Uint64(data[datapos+40:])^see2)
				datapos += 48
				i -= 48
			}
			seed ^= see1 ^ see2
		}
		if i > 16 {
			seed = rapid_mix(binary.LittleEndian.Uint64(data[datapos:])^secret[2],
				binary.LittleEndian.Uint64(data[datapos+8:])^seed^secret[1])
			if i > 32 {
				seed = rapid_mix(binary.LittleEndian.Uint64(data[datapos+16:])^secret[2],
					binary.LittleEndian.Uint64(data[datapos+24:])^seed)
			}
		}
		a = binary.LittleEndian.Uint64(data[datapos+i-16:])
		b = binary.LittleEndian.Uint64(data[datapos+i-8:])
	}
	a ^= secret[1]
	b ^= seed
	rapid_mum(&a, &b)
	return rapid_mix(a^secret[0]^bufferlen, b^secret[1])
}

func RapidhashWithSeed(data []byte, seed uint64) uint64 {
	return rapidhash_internal(data, seed, rapid_secret)
}

func Rapidhash(data []byte) uint64 {
	return RapidhashWithSeed(data, rapid_seed)
}
