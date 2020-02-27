package chars

import (
	"strconv"
	"testing"
)

const (
	dec64max = "18446744073709551615"
	dec64mid = "1844674407"
	dec32max = "4294967295"
	dec32mid = "429496"
	dec16max = "65535"
	dec16mid = "655"
	dec8max  = "255"
	dec8mid  = "25"
	decDigit = "9"
)

func TestParseUint64(t *testing.T) {
	for _, c := range []struct {
		name       string
		in         string
		expected   uint64
		expectedOk bool
	}{
		{"empty", "", 0, false},
		{"min", "0", 0, true},
		{"mid", dec64mid, 1844674407, true},
		{"max", dec64max, uint64Max, true},
		{"overflow-len", "984467440737095516150", uint64Max, false},
		{"overflow-num", "98446744073709551615", uint64Max, false},
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

func TestParseUint32(t *testing.T) {
	for _, c := range []struct {
		name       string
		in         string
		expected   uint32
		expectedOk bool
	}{
		{"empty", "", 0, false},
		{"min", "0", 0, true},
		{"mid", dec32mid, 429496, true},
		{"max", dec32max, uint32Max, true},
		{"overflow-len", "42949672950", uint32Max, false},
		{"overflow-num", "5294967295", uint32Max, false},
		{"syntax", "4w9x9x7x95", 0, false},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actual, actualOk := ParseUint32(c.in)

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
		{"mid", dec16mid, 655, true},
		{"max", dec16max, uint16Max, true},
		{"overflow-len", "655355", uint16Max, false},
		{"overflow-num", "99999", uint16Max, false},
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

func TestParseUint8(t *testing.T) {
	for _, c := range []struct {
		name       string
		in         string
		expected   uint8
		expectedOk bool
	}{
		{"empty", "", 0, false},
		{"min", "0", 0, true},
		{"mid", dec8mid, 25, true},
		{"max", dec8max, uint8Max, true},
		{"overflow-len", "2555", uint8Max, false},
		{"overflow-num", "300", uint8Max, false},
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

func BenchmarkStrconvParse8Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec8max, 10, 8)
	}
}

func BenchmarkStrconvAtoi8Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec8max)
	}
}

func BenchmarkRush8Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint8(dec8max)
	}
}

func BenchmarkStrconvParse8Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec8mid, 10, 8)
	}
}

func BenchmarkStrconvAtoi8Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec8mid)
	}
}

func BenchmarkRush8Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint8(dec8mid)
	}
}

func BenchmarkStrconvParse8Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(decDigit, 10, 8)
	}
}

func BenchmarkStrconvAtoi8Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(decDigit)
	}
}

func BenchmarkRush8Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint8(decDigit)
	}
}

func BenchmarkStrconvParse16Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec16max, 10, 16)
	}
}

func BenchmarkStrconvAtoi16Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec16max)
	}
}

func BenchmarkRush16Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint16(dec16max)
	}
}

func BenchmarkStrconvParse16Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec16mid, 10, 16)
	}
}

func BenchmarkStrconvAtoi6Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec16mid)
	}
}

func BenchmarkRush16Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint16(dec16mid)
	}
}

func BenchmarkStrconvParse16Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(decDigit, 10, 16)
	}
}

func BenchmarkStrconvAtoi16Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(decDigit)
	}
}

func BenchmarkRush16Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint16(decDigit)
	}
}

func BenchmarkStrconvParse32Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec32max, 10, 32)
	}
}

func BenchmarkStrconvAtoi32Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec32max)
	}
}

func BenchmarkRush32Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint32(dec32max)
	}
}

func BenchmarkStrconvParse32Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec32mid, 10, 32)
	}
}

func BenchmarkStrconvAtoi32Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec32mid)
	}
}

func BenchmarkRush32Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint32(dec32mid)
	}
}

func BenchmarkStrconvParse32Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(decDigit, 10, 32)
	}
}

func BenchmarkStrconvAtoi32Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(decDigit)
	}
}

func BenchmarkRush32Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint32(decDigit)
	}
}

func BenchmarkStrconvParse64Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec64max, 10, 64)
	}
}

func BenchmarkStrconvAtoi64Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec64max)
	}
}

func BenchmarkRush64Max(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint64(dec64max)
	}
}

func BenchmarkStrconvParse64Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(dec64mid, 10, 64)
	}
}

func BenchmarkStrconvAtoi64Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(dec64mid)
	}
}

func BenchmarkRush64Mid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint64(dec64mid)
	}
}

func BenchmarkStrconvParse64Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.ParseUint(decDigit, 10, 64)
	}
}

func BenchmarkStrconvAtoi64Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = strconv.Atoi(decDigit)
	}
}

func BenchmarkRush64Digit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ParseUint64(decDigit)
	}
}
