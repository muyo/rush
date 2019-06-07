package chars

// Typed decimal-string-to-unsigned-int conversions.
//
// While faster than std's strconv and for those cases where this matters or a typed API is preferred,
// the perf diff is to be considered marginal in many cases. ParseUint8 provides the biggest relative
// speedup while ParseUint16 provides the smallest.
const (
	maxUint64 = 1<<64 - 1
	maxUint32 = 1<<32 - 1
	maxUint16 = 1<<16 - 1
	maxUint8  = 1<<8 - 1

	cutoffUint8  = maxUint8/10 + 1
	cutoffUint16 = maxUint16/10 + 1
	cutoffUint64 = maxUint64/10 + 1
)

// ParseUint64 takes an unsigned integer encoded as base10 (decimal) and converts it
// to an unsigned 64-bit integer.
//
// The max length of the string is 20 characters and the max value of the number is
// maxUint64 (18446744073709551615).
// If either overflows, maxUint64 and false get returned.
// If the string contains non-numeric ASCII characters, 0 and false get returned.
//
//	BenchmarkStrconv64Max           50000000                24.1 ns/op             0 B/op          0 allocs/op
//	BenchmarkRush64Max              100000000               16.0 ns/op             0 B/op          0 allocs/op
//	BenchmarkStrconv64Mid           100000000               14.1 ns/op             0 B/op          0 allocs/op
//	BenchmarkRush64Mid              200000000                9.73 ns/op            0 B/op          0 allocs/op
//	BenchmarkStrconv64Digit         300000000                5.02 ns/op            0 B/op          0 allocs/op
//	BenchmarkRush64Digit            1000000000               2.19 ns/op            0 B/op          0 allocs/op
func ParseUint64(s string) (uint64, bool) {
	var (
		// All values are offset for the purpose of hinting to the compiler that j never gets
		// mutated after assignment, which allows us to stick to a single BCE.
		j   = len(s) - 1
		ovf int
	)

	if j < 1 {
		if j == -1 {
			return 0, false
		}

		n, ok := ParseDigit(s[0])
		return uint64(n), ok
	}

	// Overflow is only ever possible if we're at the max string length
	// of a base10 encoded number, which is why the last pass is unrolled
	// out of the multiplication loop and performed when an overflow is
	// possible (ovf = 1).
	if j > 18 {
		if j > 19 {
			return maxUint64, false
		}

		ovf = 1
	}

	var (
		r uint64
		d uint8
	)

	// BCE hint hoisted out of the loop. Unfortunately the compiler @go 1.12 isn't smart enough yet
	// to figure out that j - ovf is always < len(s), even if uints for the evaluations were used.
	// This tiny line shaves off 1.6ns down from 17.6ns for maxUint64 from the loop below, and another
	// 0.4ns from the access in the ovf condition further down.
	_ = s[j-ovf]

	for i := 0; i <= j-ovf; i++ {
		b := s[i]

		if '0' <= b && b <= '9' {
			d = b - '0'
		} else {
			// Syntax error.
			return 0, false
		}

		r = r*10 + uint64(d)
	}

	// If ovf == 0, the loop above would've handled all characters already since it was
	// impossible for it to overflow. This last pass is on the last character, which can
	// cause an overflow.
	if ovf == 1 {
		if r >= cutoffUint64 {
			// Mul would overflow.
			return maxUint64, false
		}

		b := s[j] // Technically it's always 19, but the 1.12 compiler would do a bounds-check.

		if '0' <= b && b <= '9' {
			d = b - '0'
		} else {
			// Syntax error.
			return 0, false
		}

		r *= 10
		v := r + uint64(d)

		if v < r {
			// Overflow.
			return maxUint64, false
		}

		r = v
	}

	return r, true
}

// ParseUint16 takes an unsigned integer encoded as base10 (decimal) and converts it
// to an unsigned 16-bit integer.
//
// The max length of the string is 5 characters and the max value of the number is
// maxUint16 (65535).
// If either overflows, maxUint16 and false get returned.
// If the string contains non-numeric ASCII characters, 0 and false get returned.
//
//	BenchmarkStrconv16Max           200000000                9.05 ns/op            0 B/op          0 allocs/op
//	BenchmarkRush16Max              200000000                6.39 ns/op            0 B/op          0 allocs/op
//	BenchmarkStrconv16Mid           200000000                7.04 ns/op            0 B/op          0 allocs/op
//	BenchmarkRush16Mid              200000000                6.34 ns/op            0 B/op          0 allocs/op
//	BenchmarkStrconv16Digit         300000000                5.18 ns/op            0 B/op          0 allocs/op
//	BenchmarkRush16Digit            1000000000               2.26 ns/op            0 B/op          0 allocs/op
func ParseUint16(s string) (uint16, bool) {
	var (
		j   = len(s) - 1
		ovf int
	)

	if j < 1 {
		if j == -1 {
			return 0, false
		}

		n, ok := ParseDigit(s[0])
		return uint16(n), ok
	}

	if j > 3 {
		if j > 4 {
			return maxUint16, false
		}

		ovf = 1
	}

	var (
		r uint16
		d uint8
	)

	_ = s[j-ovf]

	for i := 0; i <= j-ovf; i++ {
		b := s[i]

		if '0' <= b && b <= '9' {
			d = b - '0'
		} else {
			// Syntax error.
			return 0, false
		}

		r = r*10 + uint16(d)
	}

	if ovf == 1 {
		if r >= cutoffUint16 {
			return maxUint16, false
		}

		b := s[j]

		if '0' <= b && b <= '9' {
			d = b - '0'
		} else {
			return 0, false
		}

		r *= 10
		v := r + uint16(d)

		if v < r {
			return maxUint16, false
		}

		r = v
	}

	return r, true
}

// ParseUint8 takes an unsigned integer encoded as base10 (decimal) and converts it
// to an unsigned 8-bit integer.
//
// The max length of the string is 3 characters and the max value of the number is
// maxUint8 (255).
// If either overflows, maxUint8 and false get returned.
// If the string contains non-numeric ASCII characters, 0 and false get returned.
//
//	BenchmarkStrconv8Max            200000000                7.03 ns/op            0 B/op          0 allocs/op
//	BenchmarkRush8Max               500000000                3.03 ns/op            0 B/op          0 allocs/op
//	BenchmarkStrconv8Mid            200000000                6.02 ns/op            0 B/op          0 allocs/op
//	BenchmarkRush8Mid               1000000000               2.96 ns/op            0 B/op          0 allocs/op
//	BenchmarkStrconv8Digit          300000000                5.02 ns/op            0 B/op          0 allocs/op
//	BenchmarkRush8Digit             2000000000               1.87 ns/op            0 B/op          0 allocs/op
func ParseUint8(s string) (uint8, bool) {
	// This is entirely unrolled, but obviously too costly to ever get inlined.
	// When not unrolled, the loop would prevent inlining as well.
	var j = len(s)
	if j < 2 {
		if j == 0 {
			return 0, false
		}

		return ParseDigit(s[0])
	}

	var r, d uint8

	b := s[0]
	if '0' <= b && b <= '9' {
		r = b - '0'
	} else {
		// Syntax error.
		return 0, false
	}

	// Check for j < 2 was earlier so this is safe (and no bounds-checks either).
	b = s[1]
	if '0' <= b && b <= '9' {
		d = b - '0'
	} else {
		// Syntax error.
		return 0, false
	}

	r = r*10 + d

	if j >= 3 {
		if j > 3 || r >= cutoffUint8 {
			// Mul would overflow or more chars than could possibly fit.
			return maxUint8, false
		}

		b = s[2]
		if '0' <= b && b <= '9' {
			d = b - '0'
		} else {
			// Syntax error.
			return 0, false
		}

		r *= 10
		v := r + d

		if v < r {
			// Overflow.
			return maxUint8, false
		}

		r = v
	}

	return r, true
}

// ParseDigit takes a single character representing a base10 encoded unsigned integer
// and converts it to an unsigned 8-bit integer.
//
// If the character is a non-numeric ASCII character, 0 and false get returned.
//
// This function will usually get inlined by the compiler (@go >= 1.12).
// If that's not the case on your platform, avoid the call overhead.
func ParseDigit(b uint8) (uint8, bool) {
	if '0' <= b && b <= '9' {
		return b - '0', true
	} else {
		// Syntax error.
		return 0, false
	}
}
