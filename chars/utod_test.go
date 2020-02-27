package chars

import (
	"bytes"
	"strconv"
	"testing"
)

const (
	u64max = 18446744073709551615
	u64mid = 1844674407
	u32max = 4294967295
	u32mid = 429496
	u16max = 65535
	u16mid = 655
	u8max  = 255
	u8mid  = 25
)

func TestCopyUint64(t *testing.T) {
	buf := make([]byte, 20)
	for _, c := range []struct {
		name        string
		in          uint64
		expected    []byte
		expectedLen int
	}{
		{"min", 0, []byte("0"), 1},
		{"mid", u64mid, []byte("1844674407"), 10},
		{"max", u64max, []byte("18446744073709551615"), 20},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actualLen := CopyUint64(buf, c.in)

			if actualLen != c.expectedLen {
				t.Errorf("expected [%d], got [%d]", c.expectedLen, actualLen)
			}

			if !bytes.Equal(buf[:actualLen], c.expected) {
				t.Errorf("expected [%s], got [%s]", c.expected, buf[:actualLen])
			}
		})
	}
}

func TestCopyUint32(t *testing.T) {
	buf := make([]byte, 20)
	for _, c := range []struct {
		name        string
		in          uint32
		expected    []byte
		expectedLen int
	}{
		{"min", 0, []byte("0"), 1},
		{"mid", u32mid, []byte("429496"), 6},
		{"max", u32max, []byte("4294967295"), 10},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actualLen := CopyUint32(buf, c.in)

			if actualLen != c.expectedLen {
				t.Errorf("expected [%d], got [%d]", c.expectedLen, actualLen)
			}

			if !bytes.Equal(buf[:actualLen], c.expected) {
				t.Errorf("expected [%s], got [%s]", c.expected, buf[:actualLen])
			}
		})
	}
}

func TestCopyUint16(t *testing.T) {
	buf := make([]byte, 20)
	for _, c := range []struct {
		name        string
		in          uint16
		expected    []byte
		expectedLen int
	}{
		{"min", 0, []byte("0"), 1},
		{"mid", u16mid, []byte("655"), 3},
		{"max", u16max, []byte("65535"), 5},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actualLen := CopyUint16(buf, c.in)

			if actualLen != c.expectedLen {
				t.Errorf("expected [%d], got [%d]", c.expectedLen, actualLen)
			}

			if !bytes.Equal(buf[:actualLen], c.expected) {
				t.Errorf("expected [%s], got [%s]", c.expected, buf[:actualLen])
			}
		})
	}
}

func TestCopyUint8(t *testing.T) {
	buf := make([]byte, 20)
	for _, c := range []struct {
		name        string
		in          uint8
		expected    []byte
		expectedLen int
	}{
		{"min", 0, []byte("0"), 1},
		{"mid", u8mid, []byte("25"), 2},
		{"max", u8max, []byte("255"), 3},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actualLen := CopyUint8(buf, c.in)

			if actualLen != c.expectedLen {
				t.Errorf("expected [%d], got [%d]", c.expectedLen, actualLen)
			}

			if !bytes.Equal(buf[:actualLen], c.expected) {
				t.Errorf("expected [%s], got [%s]", c.expected, buf[:actualLen])
			}
		})
	}
}

func BenchmarkStrconvAppendUint64(b *testing.B) {
	by := make([]byte, 0, 20)

	for n := 0; n < b.N; n++ {
		_ = strconv.AppendUint(by, uint64Max, 10)
	}
}

func BenchmarkStrconvFormatUint64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.FormatUint(uint64Max, 10)
	}
}

func BenchmarkRushCopyUint64(b *testing.B) {
	by := make([]byte, 20)
	for n := 0; n < b.N; n++ {
		_ = CopyUint64(by, uint64Max)
	}
}

func BenchmarkStrconvAppendUint64Precomputed(b *testing.B) {
	by := make([]byte, 0, 20)
	for n := 0; n < b.N; n++ {
		_ = strconv.AppendUint(by, 99, 10)
	}
}

func BenchmarkStrconvFormatUint64Precomputed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.FormatUint(smallsN, 10)
	}
}

func BenchmarkRushCopyUint64Precomputed(b *testing.B) {
	by := make([]byte, 20)
	for n := 0; n < b.N; n++ {
		_ = CopyUint64(by, smallsN)
	}
}

func BenchmarkStrconvAppendUint64Tiny(b *testing.B) {
	by := make([]byte, 0, 20)
	for n := 0; n < b.N; n++ {
		_ = strconv.AppendUint(by, 9, 10)
	}
}

func BenchmarkStrconvFormatUint64Tiny(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.FormatUint(9, 10)
	}
}

func BenchmarkRushCopyUint64Tiny(b *testing.B) {
	by := make([]byte, 20)
	for n := 0; n < b.N; n++ {
		_ = CopyUint64(by, 9)
	}
}

func BenchmarkStrconvAppendUint32(b *testing.B) {
	by := make([]byte, 0, 10)
	for n := 0; n < b.N; n++ {
		_ = strconv.AppendUint(by, uint32Max, 10)
	}
}

func BenchmarkStrconvFormatUint32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.FormatUint(uint32Max, 10)
	}
}

func BenchmarkRushCopyUint32(b *testing.B) {
	by := make([]byte, 10)
	for n := 0; n < b.N; n++ {
		_ = CopyUint32(by, uint32Max)
	}
}

func BenchmarkStrconvAppendUint16(b *testing.B) {
	by := make([]byte, 0, 10)
	for n := 0; n < b.N; n++ {
		_ = strconv.AppendUint(by, uint16Max, 10)
	}
}

func BenchmarkStrconvFormatUint16(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.FormatUint(uint16Max, 10)
	}
}

func BenchmarkRushCopyUint16(b *testing.B) {
	by := make([]byte, 10)
	for n := 0; n < b.N; n++ {
		_ = CopyUint16(by, uint16Max)
	}
}

func BenchmarkStrconvAppendUint8(b *testing.B) {
	by := make([]byte, 0, 10)
	for n := 0; n < b.N; n++ {
		_ = strconv.AppendUint(by, uint8Max, 10)
	}
}

func BenchmarkStrconvFormatUint8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.FormatUint(uint8Max, 10)
	}
}

func BenchmarkRushCopyUint8(b *testing.B) {
	by := make([]byte, 10)
	for n := 0; n < b.N; n++ {
		_ = CopyUint8(by, uint8Max)
	}
}

func BenchmarkRushCopyUint8Precomputed(b *testing.B) {
	by := make([]byte, 10)
	for n := 0; n < b.N; n++ {
		_ = CopyUint8(by, smallsN)
	}
}

func BenchmarkRushCopyUint8Tiny(b *testing.B) {
	by := make([]byte, 10)
	for n := 0; n < b.N; n++ {
		_ = CopyUint8(by, 9)
	}
}
