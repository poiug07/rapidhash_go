package rapidhash

import (
	"bytes"
	"encoding/binary"
	"io"
	rand "math/rand/v2"
	"testing"
)

// Test that streaming Digest output always matches the one-shot Rapidhash.
func TestDigestMatchesOneShot(t *testing.T) {
	source := rand.NewPCG(1991, 2025)
	r := rand.New(source)

	const N = 10
	datas := make([][]byte, N)
	for idx := range N {
		datas[idx] = getRandomBytes(r, 1000)
	}

	splitPoints := []int{0, 1, 2, 3, 4, 5, 7, 8, 15, 16, 17, 48, 80, 111, 112, 113, 223, 224, 225, 500, 999, 1000}

	for i := 0; i <= 1000; i++ {
		for j := range N {
			data := datas[j][:i]
			expected := Rapidhash(data)

			// Single Write.
			h := New()
			h.Write(data)
			if got := h.Sum64(); got != expected {
				t.Errorf("single write len=%d: Sum64() = %x; want %x", i, got, expected)
			}

			// Write in every possible split point combination (two writes).
			for _, s := range splitPoints {
				if s >= i {
					continue
				}
				h := New()
				h.Write(data[:s])
				h.Write(data[s:])
				if got := h.Sum64(); got != expected {
					t.Errorf("two writes split=%d len=%d: Sum64() = %x; want %x", s, i, got, expected)
				}
			}

			// Write byte by byte.
			h = New()
			for _, b := range data {
				h.Write([]byte{b})
			}
			if got := h.Sum64(); got != expected {
				t.Errorf("byte-by-byte len=%d: Sum64() = %x; want %x", i, got, expected)
			}
		}
	}
}

// Test hash.Hash interface compliance via io.
func TestDigestHashInterface(t *testing.T) {
	data := []byte("the quick brown fox jumps over the lazy dog")
	expected := Rapidhash(data)

	h := New()
	ioCopy := bytes.NewReader(data)
	n, err := io.Copy(h, ioCopy)
	if err != nil {
		t.Fatalf("io.Copy failed: %v", err)
	}
	if n != int64(len(data)) {
		t.Errorf("io.Copy wrote %d bytes; want %d", n, len(data))
	}
	if got := h.Sum64(); got != expected {
		t.Errorf("Sum64() = %x; want %x", got, expected)
	}
	if b := h.Sum(nil); binary.LittleEndian.Uint64(b[len(b)-8:]) != expected {
		t.Errorf("Sum(nil) = %x; want %x", b, expected)
	}
	if h.Size() != 8 {
		t.Errorf("Size() = %d; want 8", h.Size())
	}
	if h.BlockSize() != 112 {
		t.Errorf("BlockSize() = %d; want 112", h.BlockSize())
	}
}

// Test that Sum64 is non-destructive and matches one-shot at every point.
func TestDigestNonDestructive(t *testing.T) {
	source := rand.NewPCG(7, 13)
	r := rand.New(source)
	data := getRandomBytes(r, 1000)

	h := New()
	for i := 0; i <= len(data); i++ {
		expected := Rapidhash(data[:i])
		if got := h.Sum64(); got != expected {
			t.Errorf("after %d bytes: Sum64() = %x; want %x", i, got, expected)
		}
		if i < len(data) {
			h.Write(data[i : i+1])
		}
	}
}

// Test that a seeded Digest matches the one-shot seeded function.
func TestDigestWithSeed(t *testing.T) {
	data := []byte("some data for seed test")
	var seed uint64 = 0x1234567890abcdef
	expected := RapidhashWithSeed(data, seed)

	h := NewSeed(seed)
	h.Write(data)
	if got := h.Sum64(); got != expected {
		t.Errorf("seeded Sum64() = %x; want %x", got, expected)
	}
}
