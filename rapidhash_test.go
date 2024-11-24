package rapidhash

import (
	"encoding/csv"
	"fmt"
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
	file, err := os.Open("file.csv")
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
