package chars

// Typed decimal-string-to-unsigned-int conversions.
//
// While faster than std's strconv and for those cases where this matters or a typed API is preferred,
// the perf diff is to be considered marginal in many cases. ParseUint8 provides the biggest relative
// speedup while ParseUint16 provides the smallest.

// ParseUint64 takes an unsigned integer encoded as base10 (decimal) and converts it
// to an unsigned 64-bit integer.
//
// The max length of the string is 20 characters and the max value of the number is
// maxUint64 (18446744073709551615).
// If either overflows, maxUint64 and false get returned.
// If the string contains non-numeric ASCII characters, 0 and false get returned.
func ParseUint64(s string) (uint64, bool) {
	// All values are offset for the purpose of hinting to the compiler that n never gets
	// mutated after assignment, which allows us to stick to a single bounds check.
	n := len(s) - 1
	if n < 1 {
		if n == -1 {
			return 0, false
		}

		n, ok := ParseDigit(s[0])
		return uint64(n), ok
	}

	var (
		ovf int
		r   uint64
		b   byte
	)

	// Overflow is only ever possible if we're at the max string length
	// of a base10 encoded number, which is why the last pass is unrolled
	// out of the multiplication loop and performed when an overflow is
	// possible (ovf = 1).
	if n >= uint64Digits-1 {
		if n > uint64Digits-1 {
			return uint64Max, false
		}

		ovf = 1
	}

	// BCE hint hoisted out of the loop. Unfortunately the compiler @go 1.12 isn't smart enough yet
	// to figure out that n - ovf is always < len(s), even if uints for the evaluations were used.
	// This tiny line shaves off 1.6ns down from 17.6ns for maxUint64 from the loop below, and another
	// 0.4ns from the access in the ovf condition further down.
	_ = s[n-ovf]

	for i := 0; i <= n-ovf; i++ {
		b = s[i] - '0'
		if b > 9 {
			// Syntax error.
			return 0, false
		}

		r = r*10 + uint64(b)
	}

	if ovf == 0 {
		return r, true
	}

	// If ovf == 0, the loop above would've handled all characters already since it was
	// impossible for it to overflow. This last pass is on the last character, which can
	// cause an overflow.
	// Technically n is always 19, but the 1.12 compiler would insert a bounds check.
	b = s[n] - '0'

	// Max    is 18446744073709551615, last digit can't be > 5 or ADD would overflow.
	// Cutoff is 1844674407370955161 accordingly.
	if r > uint64Cutoff || b > 5 {
		if b > 9 {
			return 0, false
		}

		return uint64Max, false
	}

	return r*10 + uint64(b), true
}

// ParseUint32 takes an unsigned integer encoded as base10 (decimal) and converts it
// to an unsigned 32-bit integer.
//
// The max length of the string is 10 characters and the max value of the number is
// uint32Max (4294967295).
// If either overflows, uint32Max and false get returned.
// If the string contains non-numeric ASCII characters, 0 and false get returned.
func ParseUint32(s string) (uint32, bool) {
	n := len(s) - 1
	if n < 1 {
		if n == -1 {
			return 0, false
		}

		n, ok := ParseDigit(s[0])
		return uint32(n), ok
	}

	var (
		ovf int
		r   uint32
		b   byte
	)

	if n >= uint32Digits-1 {
		if n > uint32Digits-1 {
			return uint32Max, false
		}

		ovf = 1
	}

	_ = s[n-ovf]

	for i := 0; i <= n-ovf; i++ {
		b = s[i] - '0'
		if b > 9 {
			return 0, false
		}

		r = r*10 + uint32(b)
	}

	if ovf == 0 {
		return r, true
	}

	b = s[n] - '0'

	// Max    is 4294967295, last digit can't be > 5 or ADD would overflow.
	// Cutoff is 429496729 accordingly.
	if r > uint32Cutoff || b > 5 {
		if b > 9 {
			return 0, false
		}

		return uint32Max, false
	}

	return r*10 + uint32(b), true
}

// ParseUint16 takes an unsigned integer encoded as base10 (decimal) and converts it
// to an unsigned 16-bit integer.
//
// The max length of the string is 5 characters and the max value of the number is
// uint16Max (65535).
// If either overflows, uint16Max and false get returned.
// If the string contains non-numeric ASCII characters, 0 and false get returned.
func ParseUint16(s string) (uint16, bool) {
	n := len(s) - 1
	if n < 1 {
		if n == -1 {
			return 0, false
		}

		n, ok := ParseDigit(s[0])
		return uint16(n), ok
	}

	var (
		ovf int
		r   uint16
		b   byte
	)

	if n >= uint16Digits-1 {
		if n > uint16Digits-1 {
			return uint16Max, false
		}

		ovf = 1
	}

	_ = s[n-ovf]

	for i := 0; i <= n-ovf; i++ {
		b = s[i] - '0'
		if b > 9 {
			return 0, false
		}

		r = r*10 + uint16(b)
	}

	if ovf == 0 {
		return r, true
	}

	b = s[n] - '0'

	// Max    is 65535, last digit can't be > 5 or ADD would overflow.
	// Cutoff is 6553 accordingly.
	if r > uint16Cutoff || b > 5 {
		if b > 9 {
			return 0, false
		}

		return uint16Max, false
	}

	return r*10 + uint16(b), true
}

// ParseUint8 takes an unsigned integer encoded as base10 (decimal) and converts it
// to an unsigned 8-bit integer.
//
// The max length of the string is 3 characters and the max value of the number is
// uint8Max (255).
// If either overflows, uint8Max and false get returned.
// If the string contains non-numeric ASCII characters, 0 and false get returned.
func ParseUint8(s string) (uint8, bool) {
	switch len(s) {
	case 1:
		x := s[0] - '0'
		if x > 9 {
			return 0, false
		}
		return x, true

	case 2:
		a := s[0] - '0'
		b := s[1] - '0'
		if a < 10 && b < 10 {
			return a*10 + b, true
		}

	case 3:
		a := s[0] - '0'
		b := s[1] - '0'
		c := s[2] - '0'

		if a > 9 || b > 9 || c > 9 {
			// Syntax errors.
			return 0, false
		}

		switch a {
		case 1:
			return 100 + b*10 + c, true

		case 2:
			// Catches overflow since b*10 + c = x will never be > 200, but will roll over to 0
			// when x > 55.
			if r := 200 + b*10 + c; r > 200 {
				return r, true
			}
			return uint8Max, false

		case 0:
			// Leading zero. If b is also zero, this will correctly leave c.
			return b*10 + c, true

		default:
			// a > 2
			return uint8Max, false
		}
	}

	if len(s) > uint8Digits {
		return uint8Max, false
	}

	return 0, false
}

// ParseDigit takes a single character representing a base10 encoded unsigned integer
// and converts it to an unsigned 8-bit integer.
//
// If the character is a non-numeric ASCII character, 0 and false get returned.
//
// This function will get inlined.
func ParseDigit(b uint8) (uint8, bool) {
	if b -= '0'; b < 10 {
		return b, true
	}

	// Syntax error.
	return 0, false
}
