package rapidhash

import (
	"encoding/binary"
	"encoding/csv"
	"fmt"

	"math/rand/v2"
	"os"
	"strings"
	"testing"
)

func TestRapidhash(t *testing.T) {
	testCases := []struct {
		input    string
		expected uint64
	}{
		{"test", Rapidhash([]byte("test"))},
		{"hello", Rapidhash([]byte("hello"))},
		{"world", Rapidhash([]byte("world"))},
		{"", Rapidhash([]byte(""))},
		{"123456789", Rapidhash([]byte("123456789"))},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := Rapidhash([]byte(testCase.input))
			if result != testCase.expected {
				t.Errorf("Rapidhash(%q) = %x; want %x", testCase.input, result, testCase.expected)
			} else {
				t.Logf("Rapidhash(%q) = %x", testCase.input, result) // Log the result
			}
		})
	}
}

func unescapeString(s string) string {
	s = strings.ReplaceAll(s, "\\t", "\t")
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\r", "\r")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	return s
}

func TestRapidhashFileInput(t *testing.T) {
	file, err := os.Open("additional/file.csv")
	if err != nil {
		t.Fatalf("Cannot open file %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Cannot read csv file %v", err)
	}

	for _, record := range records {
		if len(record) != 2 {
			t.Errorf("Invalid record format: %v", record)
			continue
		}

		input := unescapeString(record[0])
		if err != nil {
			t.Errorf("Unable to unquote %s: %v", record[0], err)
		}
		expectedHash := record[1]

		result := fmt.Sprintf("%x", Rapidhash([]byte(input)))

		if result != expectedHash {
			t.Errorf("Rapidhash(%q) = %s; want %s", input, result, expectedHash)
		} else {
			// t.Logf("Rapidhash(%q) = %s", input, result)
		}
	}
}

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

func TestRapidhashVariantsMatching(t *testing.T) {
	source := rand.NewPCG(1991, 2025)
	r := rand.New(source)
	const N int = 10
	datas := make([][]byte, N)
	// redundant for len 0, but makes code simpler.
	for idx := range N {
		datas[idx] = getRandomBytes(r, 81)
	}
	for i := 0; i <= 48; i++ {
		for j := range N {
			bytes := datas[j][:i]

			var expected = Rapidhash(bytes)
			micro := RapidhashMicro(bytes)
			nano := RapidhashNano(bytes)

			if micro != expected {
				t.Errorf("RapidhashMicro(%q) = %d; want %d", bytes, micro, expected)
			}
			if nano != expected {
				t.Errorf("RapidhashNano(%q) = %d; want %d", bytes, nano, expected)
			}
		}
	}

	for i := 49; i <= 80; i++ {
		for j := range N {
			bytes := datas[j][:i]

			var expected = Rapidhash(bytes)
			micro := RapidhashMicro(bytes)

			if micro != expected {
				t.Errorf("RapidhashMicro(%q) = %d; want %d", bytes, micro, expected)
			}
		}
	}
}
