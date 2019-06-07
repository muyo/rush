package chars

import (
	"strconv"
	"testing"
)

const (
	dec64max = "18446744073709551615"
	dec64mid = "1844674407"
	dec32    = "4294967295"
	dec16max = "65535"
	dec16mid = "655"
	dec8max  = "255"
	dec8mid  = "25"
	digit    = "9"
)

func TestParseUint8(t *testing.T) {
	for _, c := range []struct {
		name       string
		in         string
		expected   uint8
		expectedOk bool
	}{
		{"empty", "", 0, false},
		{"min", "0", 0, true},
		{"mid", "25", 25, true},
		{"max", "255", 255, true},
		{"overflow-len", "2555", maxUint8, false},
		{"overflow-num", "300", maxUint8, false},
		{"syntax", "2a6", 0, false},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actual, actualOk := ParseUint8(c.in)

			if actual != c.expected {
				t.Errorf("expected [%d], got [%d]", c.expected, actual)
			}

			if actualOk != c.expectedOk {
				t.Errorf("expected [%t], got [%t]", c.expectedOk, actualOk)
			}
		})
	}
}

func TestParseUint16(t *testing.T) {
	for _, c := range []struct {
		name       string
		in         string
		expected   uint16
		expectedOk bool
	}{
		{"empty", "", 0, false},
		{"min", "0", 0, true},
		{"mid", "655", 655, true},
		{"max", "65535", 65535, true},
		{"overflow-len", "655355", maxUint16, false},
		{"overflow-num", "99999", maxUint16, false},
		{"syntax", "6aaa5", 0, false},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actual, actualOk := ParseUint16(c.in)

			if actual != c.expected {
				t.Errorf("expected [%d], got [%d]", c.expected, actual)
			}

			if actualOk != c.expectedOk {
				t.Errorf("expected [%t], got [%t]", c.expectedOk, actualOk)
			}
		})
	}
}

func TestParseUint64(t *testing.T) {
	for _, c := range []struct {
		name       string
		in         string
		expected   uint64
		expectedOk bool
	}{
		{"empty", "", 0, false},
		{"min", "0", 0, true},
		{"mid", "1844674407", 1844674407, true},
		{"max", "18446744073709551615", 18446744073709551615, true},
		{"overflow-len", "984467440737095516150", maxUint64, false},
		{"overflow-num", "98446744073709551615", maxUint64, false},
		{"syntax", "984467dddddd", 0, false},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actual, actualOk := ParseUint64(c.in)

			if actual != c.expected {
				t.Errorf("expected [%d], got [%d]", c.expected, actual)
			}

			if actualOk != c.expectedOk {
				t.Errorf("expected [%t], got [%t]", c.expectedOk, actualOk)
			}
		})
	}
}

func BenchmarkStrconv8Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec8max, 10, 8)
	}
}

func BenchmarkRush8Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint8(dec8max)
	}
}

func BenchmarkStrconv8Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec8mid, 10, 8)
	}
}

func BenchmarkRush8Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint8(dec8mid)
	}
}

func BenchmarkStrconv8Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(digit, 10, 8)
	}
}

func BenchmarkRush8Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint8(digit)
	}
}

func BenchmarkStrconv16Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec16max, 10, 16)
	}
}

func BenchmarkRush16Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint16(dec16max)
	}
}

func BenchmarkStrconv16Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec16mid, 10, 16)
	}
}

func BenchmarkRush16Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint16(dec16mid)
	}
}

func BenchmarkStrconv16Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(digit, 10, 16)
	}
}

func BenchmarkRush16Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint16(digit)
	}
}

func BenchmarkStrconv64Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec64max, 10, 64)
	}
}

func BenchmarkRush64Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint64(dec64max)
	}
}

func BenchmarkStrconv64Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec64mid, 10, 64)
	}
}

func BenchmarkRush64Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint64(dec64mid)
	}
}

func BenchmarkStrconv64Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(digit, 10, 64)
	}
}

func BenchmarkRush64Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint64(digit)
	}
}
