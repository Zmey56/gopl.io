package main_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	// Prepare test input
	input := "Hello, 世界!" // Sample input
	expectedCounts := map[rune]int{
		'H': 1, 'e': 1, 'l': 2, 'o': 1, ',': 1, ' ': 1, '世': 1, '界': 1, '!': 1,
	}
	expectedUtfLen := [utf8.UTFMax + 1]int{0, 8, 0, 2, 0}
	expectedInvalid := 0

	// Redirect stdin for testing
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	fmt.Fprint(w, input)
	w.Close()

	// Execute the code
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("charcount: %v", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	// Restore stdin
	os.Stdin = oldStdin

	// Compare the results
	for r, count := range expectedCounts {
		if counts[r] != count {
			t.Errorf("Incorrect count for rune %q. Expected: %d, Actual: %d", r, count, counts[r])
		}
	}
	for i, length := range expectedUtfLen {
		if utflen[i] != length {
			t.Errorf("Incorrect count for length %d. Expected: %d, Actual: %d", i, length, utflen[i])
		}
	}
	if invalid != expectedInvalid {
		t.Errorf("Incorrect count of invalid UTF-8 characters. Expected: %d, Actual: %d", expectedInvalid, invalid)
	}
}
