package rapidhash

import (
	"encoding/binary"
	"math/bits"
)

// Default seed
// const rapid_seed uint64 = 0xBDD89AA982704029

// Default secret parameters
var rapid_secret [8]uint64 = [8]uint64{
	0x2d358dccaa6c78a5,
	0x8bb84b93962eacc9,
	0x4b33a62ed433d4a3,
	0x4d5a2da51de1aa47,
	0xa0761d6478bd642f,
	0xe7037ed1a0b428db,
	0x90ed1765281c388c,
	0xaaaaaaaaaaaaaaaa,
}

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

/*
 *  Multiply and xor mix function.
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

func read64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)
}

func read32(data []byte) uint64 {
	return uint64(binary.LittleEndian.Uint32(data))
}

func rapidhash_internal(data []byte, seed uint64, secret [8]uint64) uint64 {
	bufferlen := uint64(len(data))
	seed ^= rapid_mix(seed^secret[2], secret[1])
	var a, b uint64
	p := data
	i := bufferlen

	if bufferlen <= 16 {
		if bufferlen >= 4 {
			seed ^= bufferlen
			if bufferlen >= 8 {
				var datalast = data[bufferlen-8:]
				a = read64(data)
				b = read64(datalast)
			} else {
				var datalast = data[bufferlen-4:]
				a = read32(data)
				b = read32(datalast)
			}
		} else {
			if bufferlen > 0 {
				a = (((uint64)(data[0])) << 45) | ((uint64)(data[bufferlen-1]))
				b = (uint64)(data[bufferlen>>1])
			} /* else { // Initialized to 0 by default
				a = 0
				b = 0
			} */
		}
	} else {
		var see1, see2 = seed, seed
		var see3, see4 = seed, seed
		var see5, see6 = seed, seed

		// unrolled loop
		if i > 224 {
			for i >= 224 {
				seed = rapid_mix(read64(p)^secret[0], read64(p[8:])^seed)
				see1 = rapid_mix(read64(p[16:])^secret[1], read64(data[24:])^see1)
				see2 = rapid_mix(read64(p[32:])^secret[2], read64(data[40:])^see2)
				see3 = rapid_mix(read64(p[48:])^secret[3], read64(data[56:])^see3)
				see4 = rapid_mix(read64(p[64:])^secret[4], read64(data[72:])^see4)
				see5 = rapid_mix(read64(p[80:])^secret[5], read64(data[88:])^see5)
				see6 = rapid_mix(read64(p[96:])^secret[6], read64(data[104:])^see6)
				seed = rapid_mix(read64(p[112:])^secret[0], read64(p[120:])^seed)
				see1 = rapid_mix(read64(p[128:])^secret[1], read64(data[136:])^see1)
				see2 = rapid_mix(read64(p[144:])^secret[2], read64(data[152:])^see2)
				see3 = rapid_mix(read64(p[160:])^secret[3], read64(data[168:])^see3)
				see4 = rapid_mix(read64(p[176:])^secret[4], read64(data[184:])^see4)
				see5 = rapid_mix(read64(p[192:])^secret[5], read64(data[200:])^see5)
				see6 = rapid_mix(read64(p[208:])^secret[6], read64(data[216:])^see6)
				p = p[224:]
				i -= 224
			}
		}

		if i > 112 {
			seed = rapid_mix(read64(p)^secret[0], read64(p[8:])^seed)
			see1 = rapid_mix(read64(p[16:])^secret[1], read64(data[24:])^see1)
			see2 = rapid_mix(read64(p[32:])^secret[2], read64(data[40:])^see2)
			see3 = rapid_mix(read64(p[48:])^secret[3], read64(data[56:])^see3)
			see4 = rapid_mix(read64(p[64:])^secret[4], read64(data[72:])^see4)
			see5 = rapid_mix(read64(p[80:])^secret[5], read64(data[88:])^see5)
			see6 = rapid_mix(read64(p[96:])^secret[6], read64(data[104:])^see6)
			p = p[112:]
			i -= 112
		}

		seed ^= see1
		see2 ^= see3
		see4 ^= see5
		seed ^= see6
		see2 ^= see4
		seed ^= see2

		if i > 16 {
			seed = rapid_mix(read64(p)^secret[2], read64(p[8:])^seed)
			if i > 32 {
				seed = rapid_mix(read64(p[16:])^secret[2], read64(p[24:])^seed)
				if i > 48 {
					seed = rapid_mix(read64(p[32:])^secret[1], read64(p[40:])^seed)
					if i > 64 {
						seed = rapid_mix(read64(p[48:])^secret[1], read64(p[56:])^seed)
						if i > 80 {
							seed = rapid_mix(read64(p[64:])^secret[2], read64(p[72:])^seed)
							if i > 96 {
								seed = rapid_mix(read64(p[80:])^secret[1], read64(p[88:])^seed)
							}
						}
					}
				}
			}
		}
		a = read64(data[bufferlen-((uint64)(len(p)))+i-16:]) ^ i
		b = read64(data[bufferlen-((uint64)(len(p)))+i-8:])
	}
	a ^= secret[1]
	b ^= seed
	rapid_mum(&a, &b)
	return rapid_mix(a^secret[7], b^secret[1]^i)
}

func RapidhashWithSeed(data []byte, seed uint64) uint64 {
	return rapidhash_internal(data, seed, rapid_secret)
}

func Rapidhash(data []byte) uint64 {
	return RapidhashWithSeed(data, 0)
}

func rapidhash_micro_internal(data []byte, seed uint64, secret [8]uint64) uint64 {
	bufferlen := uint64(len(data))
	p := data
	i := bufferlen
	var a, b uint64

	seed ^= rapid_mix(seed^secret[2], secret[1])
	if bufferlen <= 16 {
		if bufferlen >= 4 {
			seed ^= bufferlen
			if bufferlen >= 8 {
				var datalast = data[bufferlen-8:]
				a = read64(data)
				b = read64(datalast)
			} else {
				var datalast = data[bufferlen-4:]
				a = read32(data)
				b = read32(datalast)
			}
		} else {
			if bufferlen > 0 {
				a = (((uint64)(data[0])) << 45) | ((uint64)(data[bufferlen-1]))
				b = (uint64)(data[bufferlen>>1])
			} /* else { // Initialized to 0 by default
				a = 0
				b = 0
			} */
		}
	} else {
		if i > 80 {
			var see1, see2, see3, see4 = seed, seed, seed, seed
			for i >= 80 {
				seed = rapid_mix(read64(p)^secret[0], read64(p[8:])^seed)
				see1 = rapid_mix(read64(p[16:])^secret[1], read64(p[24:])^see1)
				see2 = rapid_mix(read64(p[32:])^secret[2], read64(p[40:])^see2)
				see3 = rapid_mix(read64(p[48:])^secret[3], read64(p[56:])^see3)
				see4 = rapid_mix(read64(p[64:])^secret[4], read64(p[72:])^see4)
				p = p[80:]
				i -= 80
			}
			seed ^= see1
			see2 ^= see3
			seed ^= see4
			seed ^= see2
		}
		if i > 16 {
			seed = rapid_mix(read64(p)^secret[2], read64(p[8:])^seed)
			if i > 32 {
				seed = rapid_mix(read64(p[16:])^secret[2], read64(p[24:])^seed)
				if i > 48 {
					seed = rapid_mix(read64(p[32:])^secret[1], read64(p[40:])^seed)
					if i > 64 {
						seed = rapid_mix(read64(p[48:])^secret[1], read64(p[56:])^seed)
					}
				}
			}
		}
		a = read64(data[bufferlen-((uint64)(len(p)))+i-16:]) ^ i
		b = read64(data[bufferlen-((uint64)(len(p)))+i-8:])
	}
	a ^= secret[1]
	b ^= seed
	rapid_mum(&a, &b)
	return rapid_mix(a^secret[7], b^secret[1]^i)
}

func rapidhash_nano_internal(data []byte, seed uint64, secret [8]uint64) uint64 {
	bufferlen := uint64(len(data))
	p := data
	i := bufferlen
	var a, b uint64

	seed ^= rapid_mix(seed^secret[2], secret[1])
	if bufferlen <= 16 {
		if bufferlen >= 4 {
			seed ^= bufferlen
			if bufferlen >= 8 {
				var datalast = data[bufferlen-8:]
				a = read64(data)
				b = read64(datalast)
			} else {
				var datalast = data[bufferlen-4:]
				a = read32(data)
				b = read32(datalast)
			}
		} else {
			if bufferlen > 0 {
				a = (((uint64)(data[0])) << 45) | ((uint64)(data[bufferlen-1]))
				b = (uint64)(data[bufferlen>>1])
			} /* else { // Initialized to 0 by default
				a = 0
				b = 0
			} */
		}
	} else {
		if i > 48 {
			var see1, see2 = seed, seed
			for i >= 48 {
				seed = rapid_mix(read64(p)^secret[0], read64(p[8:])^seed)
				see1 = rapid_mix(read64(p[16:])^secret[1], read64(p[24:])^see1)
				see2 = rapid_mix(read64(p[32:])^secret[2], read64(p[40:])^see2)
				p = p[48:]
				i -= 48
			}
			seed ^= see1
			seed ^= see2
		}
		if i > 16 {
			seed = rapid_mix(read64(p)^secret[2], read64(p[8:])^seed)
			if i > 32 {
				seed = rapid_mix(read64(p[16:])^secret[2], read64(p[24:])^seed)
			}
		}
		a = read64(data[bufferlen-((uint64)(len(p)))+i-16:]) ^ i
		b = read64(data[bufferlen-((uint64)(len(p)))+i-8:])
	}
	a ^= secret[1]
	b ^= seed
	rapid_mum(&a, &b)
	return rapid_mix(a^secret[7], b^secret[1]^i)
}

func RapidhashMicroWithSeed(data []byte, seed uint64) uint64 {
	return rapidhash_micro_internal(data, seed, rapid_secret)
}

func RapidhashMicro(data []byte) uint64 {
	return RapidhashMicroWithSeed(data, 0)
}

func RapidhashNanoWithSeed(data []byte, seed uint64) uint64 {
	return rapidhash_nano_internal(data, seed, rapid_secret)
}

func RapidhashNano(data []byte) uint64 {
	return RapidhashNanoWithSeed(data, 0)
}
