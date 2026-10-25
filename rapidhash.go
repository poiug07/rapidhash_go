package rapidhash

import "encoding/binary"

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
				a = read64(data)
				b = read64(data[bufferlen-8:])
			} else {
				a = read32(data)
				b = read32(data[bufferlen-4:])
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
			for i > 224 {
				_ = p[223] // bounds-check elimination hint: p has >= 224 bytes
				seed = rapid_mix(read64(p)^secret[0], read64(p[8:])^seed)
				see1 = rapid_mix(read64(p[16:])^secret[1], read64(p[24:])^see1)
				see2 = rapid_mix(read64(p[32:])^secret[2], read64(p[40:])^see2)
				see3 = rapid_mix(read64(p[48:])^secret[3], read64(p[56:])^see3)
				see4 = rapid_mix(read64(p[64:])^secret[4], read64(p[72:])^see4)
				see5 = rapid_mix(read64(p[80:])^secret[5], read64(p[88:])^see5)
				see6 = rapid_mix(read64(p[96:])^secret[6], read64(p[104:])^see6)
				seed = rapid_mix(read64(p[112:])^secret[0], read64(p[120:])^seed)
				see1 = rapid_mix(read64(p[128:])^secret[1], read64(p[136:])^see1)
				see2 = rapid_mix(read64(p[144:])^secret[2], read64(p[152:])^see2)
				see3 = rapid_mix(read64(p[160:])^secret[3], read64(p[168:])^see3)
				see4 = rapid_mix(read64(p[176:])^secret[4], read64(p[184:])^see4)
				see5 = rapid_mix(read64(p[192:])^secret[5], read64(p[200:])^see5)
				see6 = rapid_mix(read64(p[208:])^secret[6], read64(p[216:])^see6)
				p = p[224:]
				i -= 224
			}
		}

		if i > 112 {
			_ = p[111] // bounds-check elimination hint: p has >= 112 bytes
			seed = rapid_mix(read64(p)^secret[0], read64(p[8:])^seed)
			see1 = rapid_mix(read64(p[16:])^secret[1], read64(p[24:])^see1)
			see2 = rapid_mix(read64(p[32:])^secret[2], read64(p[40:])^see2)
			see3 = rapid_mix(read64(p[48:])^secret[3], read64(p[56:])^see3)
			see4 = rapid_mix(read64(p[64:])^secret[4], read64(p[72:])^see4)
			see5 = rapid_mix(read64(p[80:])^secret[5], read64(p[88:])^see5)
			see6 = rapid_mix(read64(p[96:])^secret[6], read64(p[104:])^see6)
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
		a = read64(data[bufferlen-16:]) ^ i
		b = read64(data[bufferlen-8:])
	}
	a ^= secret[1]
	b ^= seed
	hi, lo := rapid_mum(a, b)
	a, b = lo, hi
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
				a = read64(data)
				b = read64(data[bufferlen-8:])
			} else {
				a = read32(data)
				b = read32(data[bufferlen-4:])
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
			for i > 80 {
				_ = p[79] // bounds-check elimination hint: p has >= 80 bytes
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
		a = read64(data[bufferlen-16:]) ^ i
		b = read64(data[bufferlen-8:])
	}
	a ^= secret[1]
	b ^= seed
	hi, lo := rapid_mum(a, b)
	a, b = lo, hi
	return rapid_mix(a^secret[7], b^secret[1]^i)
}

func RapidhashMicroWithSeed(data []byte, seed uint64) uint64 {
	return rapidhash_micro_internal(data, seed, rapid_secret)
}

func RapidhashMicro(data []byte) uint64 {
	return RapidhashMicroWithSeed(data, 0)
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
				a = read64(data)
				b = read64(data[bufferlen-8:])
			} else {
				a = read32(data)
				b = read32(data[bufferlen-4:])
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
			for i > 48 {
				_ = p[47] // bounds-check elimination hint: p has >= 48 bytes
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
		a = read64(data[bufferlen-16:]) ^ i
		b = read64(data[bufferlen-8:])
	}
	a ^= secret[1]
	b ^= seed
	hi, lo := rapid_mum(a, b)
	a, b = lo, hi
	return rapid_mix(a^secret[7], b^secret[1]^i)
}

func RapidhashNanoWithSeed(data []byte, seed uint64) uint64 {
	return rapidhash_nano_internal(data, seed, rapid_secret)
}

func RapidhashNano(data []byte) uint64 {
	return RapidhashNanoWithSeed(data, 0)
}

type Digest struct {
	buffer   [chunkPrev + chunkSize]byte
	buffered int    // Length of the chunk part of buffer currently in use
	chunks   uint64 // Have we processed a full 112-byte chunk?
	seed     uint64 // Primary seed state, updated on each processed chunk
	see1     uint64 // Secondary states, carried across chunks
	see2     uint64
	see3     uint64
	see4     uint64
	see5     uint64
	see6     uint64
}

const chunkSize = 112 // The size of a single rapidhash processing chunk
const chunkPrev = 16  // The minimum tail we must keep in the buffer for processing

func New() *Digest {
	d := Digest{}
	d.ResetWithSeed(0)
	return &d
}

func NewSeed(seed uint64) *Digest {
	d := Digest{}
	d.ResetWithSeed(seed)
	return &d
}

func (d *Digest) Reset() {
	d.ResetWithSeed(0)
}

/*
 *  Resets the internal state of the hasher to a new seed, equivalent to
 *  reinitializing a new hasher with the given seed
 */
func (d *Digest) ResetWithSeed(seed uint64) {
	d.buffered = 0
	d.chunks = 0
	d.seed = seed ^ rapid_mix(seed^rapid_secret[2], rapid_secret[1])
	d.see1, d.see2 = d.seed, d.seed
	d.see3, d.see4 = d.seed, d.seed
	d.see5, d.see6 = d.seed, d.seed
}

/*
 *  Writes the given bytes into the hasher.
 *
 *  Accepts any amount of data, and will buffer any trailing bytes that do not
 *  fill a complete 112-byte chunk until further Write() or Sum64() calls.
 */
func (d *Digest) Write(data []byte) (n int, err error) {
	n = len(data)

	if len(data)+d.buffered > chunkSize {
		// Slow path: the data doesn't fit in the remaining buffer space
		d.writeInner(data)
	} else {
		start := chunkPrev + d.buffered
		copy(d.buffer[start:start+len(data)], data)
		d.buffered += len(data)
	}
	return n, nil
}

/*
 *  Slow-path write that processes full chunks directly from data. We keep back
 *  the last chunk when chunk-aligned: rapidhash uses `pos + 112 < len` (not
 *  <=), so the final 112 bytes must go through the tail path in finish().
 */
func (d *Digest) writeInner(data []byte) {
	chunkBuffer := d.buffer[chunkPrev:]

	// Fill the current chunk up to 112 bytes and process it
	copyBytes := chunkSize - d.buffered
	copy(chunkBuffer[d.buffered:chunkSize], data[:copyBytes])
	d.processChunk(chunkBuffer)
	d.buffered = 0

	// Process chunks without copying, keeping the last chunk back
	remaining := data[copyBytes:]
	stop := (len(remaining) - 1) / chunkSize * chunkSize
	var lastChunk []byte
	pos := 0
	for pos < stop {
		d.processChunk(remaining[pos : pos+chunkSize])
		lastChunk = remaining[pos : pos+chunkSize]
		pos += chunkSize
	}

	// Copy the final 16 data bytes from the previous chunk into chunkPrev
	if lastChunk != nil {
		copy(d.buffer[:chunkPrev], lastChunk[chunkSize-chunkPrev:])
	} else {
		copy(d.buffer[:chunkPrev], chunkBuffer[chunkSize-chunkPrev:])
	}

	// Copy the remainder into the chunk part of the buffer
	unprocessed := remaining[pos:]
	copy(chunkBuffer[:len(unprocessed)], unprocessed)
	d.buffered = len(unprocessed)
}

func (d *Digest) Sum64() uint64 {
	return d.finish()
}

func (d *Digest) Sum(b []byte) []byte {
	var out [8]byte
	binary.LittleEndian.PutUint64(out[:], d.Sum64())
	return append(b, out[:]...)
}

func (d *Digest) Size() int {
	return 8
}

// Returns the internal buffer block size in bytes, implementing the hash.Hash interface. We process data in 112-byte chunks
func (d *Digest) BlockSize() int {
	return 112
}

func (d *Digest) processChunk(p []byte) {
	seed := d.seed
	_ = p[111] // bounds-check elimination hint: p has == 112 bytes
	d.seed = rapid_mix(read64(p)^rapid_secret[0], read64(p[8:])^seed)
	d.see1 = rapid_mix(read64(p[16:])^rapid_secret[1], read64(p[24:])^d.see1)
	d.see2 = rapid_mix(read64(p[32:])^rapid_secret[2], read64(p[40:])^d.see2)
	d.see3 = rapid_mix(read64(p[48:])^rapid_secret[3], read64(p[56:])^d.see3)
	d.see4 = rapid_mix(read64(p[64:])^rapid_secret[4], read64(p[72:])^d.see4)
	d.see5 = rapid_mix(read64(p[80:])^rapid_secret[5], read64(p[88:])^d.see5)
	d.see6 = rapid_mix(read64(p[96:])^rapid_secret[6], read64(p[104:])^d.see6)
	d.chunks++
}

/*
 *  Finishes the hash on the buffered data (at most 112 bytes). Does not
 *  modify the hasher state.
 */
func (d *Digest) finish() uint64 {
	seed := d.seed
	bufferlen := uint64(d.buffered)
	var a, b uint64

	if d.chunks == 0 && bufferlen <= 16 {
		// Short <=16 pass only if we haven't processed a full chunk yet
		data := d.buffer[chunkPrev : chunkPrev+int(bufferlen)]
		if bufferlen >= 4 {
			seed ^= bufferlen
			if bufferlen >= 8 {
				a = read64(data)
				b = read64(data[bufferlen-8:])
			} else {
				a = read32(data)
				b = read32(data[bufferlen-4:])
			}
		} else if bufferlen > 0 {
			a = (((uint64)(data[0])) << 45) | ((uint64)(data[bufferlen-1]))
			b = (uint64)(data[bufferlen>>1])
		}
	} else {
		if d.chunks > 0 {
			// Merge independent lanes if we'd previously processed a full chunk
			seed = d.seed ^ d.see1 ^ d.see2 ^ d.see3 ^ d.see4 ^ d.see5 ^ d.see6
		}

		// The >16 tail is the same whether we've processed a full chunk or not
		slice := d.buffer[chunkPrev : chunkPrev+int(bufferlen)]
		if bufferlen > 16 {
			seed = rapid_mix(read64(slice)^rapid_secret[2], read64(slice[8:])^seed)
			if bufferlen > 32 {
				seed = rapid_mix(read64(slice[16:])^rapid_secret[2], read64(slice[24:])^seed)
				if bufferlen > 48 {
					seed = rapid_mix(read64(slice[32:])^rapid_secret[1], read64(slice[40:])^seed)
					if bufferlen > 64 {
						seed = rapid_mix(read64(slice[48:])^rapid_secret[1], read64(slice[56:])^seed)
						if bufferlen > 80 {
							seed = rapid_mix(read64(slice[64:])^rapid_secret[2], read64(slice[72:])^seed)
							if bufferlen > 96 {
								seed = rapid_mix(read64(slice[80:])^rapid_secret[1], read64(slice[88:])^seed)
							}
						}
					}
				}
			}
		}

		// The final 16 bytes may read from the chunkPrev part of the buffer
		data := d.buffer[:chunkPrev+int(bufferlen)]
		a = read64(data[len(data)-16:]) ^ bufferlen
		b = read64(data[len(data)-8:])
	}
	a ^= rapid_secret[1]
	b ^= seed
	hi, lo := rapid_mum(a, b)
	a, b = lo, hi
	return rapid_mix(a^rapid_secret[7], b^rapid_secret[1]^bufferlen)
}
