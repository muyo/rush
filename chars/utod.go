package chars

// LUT for numbers < 100 (from std strconv).
const (
	smallsN = 99
	smalls  = "00010203040506070809" +
		"10111213141516171819" +
		"20212223242526272829" +
		"30313233343536373839" +
		"40414243444546474849" +
		"50515253545556575859" +
		"60616263646566676869" +
		"70717273747576777879" +
		"80818283848586878889" +
		"90919293949596979899"
)

// CopyUint64 copies the base10 representation of a uint64 into dst up to len(dst),
// discarding all overflowing bytes.
//
// Returns the number of bytes copied to dst.
//
// Works similar to strconv.AppendUint() but puts values starting at the beginning of dst
// instead of appending and does not grow dst.
func CopyUint64(dst []byte, u uint64) int {
	if len(dst) == 0 {
		return 0
	}

	if u <= smallsN {
		return copySmalls(dst, uint8(u))
	}

	var (
		i = uint64Digits
		b [uint64Digits]byte
		q uint64
	)

	for u >= 10000 {
		q := u % 10000
		u /= 10000

		d1 := q / 100 * 2
		d2 := q % 100 * 2
		i -= 4

		b[i], b[i+1] = smalls[d1], smalls[d1+1]
		b[i+2], b[i+3] = smalls[d2], smalls[d2+1]
	}

	for u > smallsN {
		q = u % 100 * 2
		u /= 100
		i -= 2

		b[i], b[i+1] = smalls[q], smalls[q+1]
	}

	if u < 10 {
		i--
		b[i] = '0' + byte(u)

		return copy(dst, b[i:])
	}

	u *= 2
	i -= 2
	b[i], b[i+1] = smalls[u], smalls[u+1]

	return copy(dst, b[i:])
}

// CopyUint32 copies the base10 representation of a uint32 into dst up to len(dst),
// discarding all overflowing bytes.
//
// Returns the number of bytes copied to dst.
//
// Works similar to strconv.AppendUint() but puts values starting at the beginning of dst
// instead of appending and does not grow dst.
func CopyUint32(dst []byte, u uint32) int {
	if len(dst) == 0 {
		return 0
	}

	if u <= smallsN {
		return copySmalls(dst, uint8(u))
	}

	var (
		i = uint32Digits
		b [uint32Digits]byte
		q uint32
	)

	for u >= 10000 {
		q := u % 10000
		u /= 10000

		s1 := q / 100 * 2
		s2 := q % 100 * 2
		i -= 4

		b[i], b[i+1] = smalls[s1], smalls[s1+1]
		b[i+2], b[i+3] = smalls[s2], smalls[s2+1]
	}

	for u > smallsN {
		q = u % 100 * 2
		u /= 100
		i -= 2

		b[i], b[i+1] = smalls[q], smalls[q+1]
	}

	if u < 10 {
		i--
		b[i] = '0' + byte(u)

		return copy(dst, b[i:])
	}

	u *= 2
	i -= 2
	b[i], b[i+1] = smalls[u], smalls[u+1]

	return copy(dst, b[i:])
}

// CopyUint16 copies the base10 representation of a uint16 into dst up to len(dst),
// discarding all overflowing bytes.
//
// Returns the number of bytes copied to dst.
//
// Works similar to strconv.AppendUint() but puts values starting at the beginning of dst
// instead of appending and does not grow dst.
func CopyUint16(dst []byte, u uint16) int {
	if len(dst) == 0 {
		return 0
	}

	if u <= smallsN {
		return copySmalls(dst, uint8(u))
	}

	var (
		i = uint16Digits
		b [uint16Digits]byte
		q uint16
	)

	for u > smallsN {
		q = u % 100 * 2
		u /= 100
		i -= 2

		b[i], b[i+1] = smalls[q], smalls[q+1]
	}

	if u < 10 {
		i--
		b[i] = '0' + byte(u)

		return copy(dst, b[i:])
	}

	u *= 2
	i -= 2
	b[i], b[i+1] = smalls[u], smalls[u+1]

	return copy(dst, b[i:])
}

// CopyUint8 copies the base10 representation of a uint8 into dst up to len(dst),
// discarding all overflowing bytes.
//
// Returns the number of bytes copied to dst.
//
// Works similar to strconv.AppendUint() but puts values starting at the beginning of dst
// instead of appending and does not grow dst.
func CopyUint8(dst []byte, u uint8) int {
	if len(dst) == 0 {
		return 0
	}

	if u <= smallsN {
		return copySmalls(dst, u)
	}

	// We can avoid the DIV and MUL here since we're constrained
	// by an uint8 anyways.
	if u > 199 {
		u -= 200
		dst[0] = '2'
	} else {
		u -= 100
		dst[0] = '1'
	}

	if len(dst) > 1 {
		u *= 2
		dst[1] = smalls[u]

		if len(dst) > 2 {
			dst[2] = smalls[u+1]
			return 3
		}

		return 2
	}

	return 1
}

// Gets inlined.
func copySmalls(dst []byte, u uint8) int {
	if u < 10 {
		dst[0] = '0' + u
		return 1
	}

	u *= 2

	if len(dst) > 1 {
		dst[0], dst[1] = smalls[u], smalls[u+1]
		return 2
	}

	dst[0] = smalls[u]
	return 1
}
