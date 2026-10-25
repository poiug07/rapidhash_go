package rapidhash

import (
	"encoding/binary"
	rand "math/rand/v2"
)

func getRandomBytes(r rand.Source, len int) []byte {
	data := make([]byte, len)
	i := len
	cur := 0
	for i >= 8 {
		binary.LittleEndian.PutUint64(data[cur:], r.Uint64())
		i -= 8
		cur += 8
	}
	if i > 4 {
		binary.LittleEndian.PutUint32(data[cur:], uint32(r.Uint64()))
		cur += 4
	}
	return data
}
