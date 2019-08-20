# chars [![GoDoc](https://godoc.org/github.com/muyo/rush/chars?status.svg)](https://godoc.org/github.com/muyo/rush/chars) 

Currently provides some functions hoisted out of an otherwise private repo due to public dependencies.

### dtou (atou/atoi)

Typed alternatives to `strconv.ParseUint(dec, 10, bitsN)` (e.g. base10 string encoded numbers) with a non-idiomatic 
API (bool returns instead of err), which happened to present some optimization opportunities.

Since those tend to be in hot paths in a server context (both for arg/request parsing and response encoding),
those seemingly minor performance gains start adding up. SIMD versions of ParseUint64/32 are on the table and
should provide an additional speedup.

```
Suffixes:
 - max: max value for a given integer size.
 - mid: a value somewhere in the middle range (about half length of max in dec characters).
 - digit: single digit (gets rough call overhead).

Platform: go 1.12, i7 4770k @ 4.4GHz; ran at 2019/08/20; code included in test file.
```

**chars.ParseUint64()**
```
BenchmarkStrconv64Max           50000000                25.5 ns/op             0 B/op          0 allocs/op
BenchmarkRush64Max              100000000               14.4 ns/op             0 B/op          0 allocs/op
BenchmarkStrconv64Mid           100000000               14.9 ns/op             0 B/op          0 allocs/op
BenchmarkRush64Mid              200000000                9.13 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv64Digit         300000000                5.65 ns/op            0 B/op          0 allocs/op
BenchmarkRush64Digit            1000000000               2.28 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint32()**
```
BenchmarkStrconv32Max           100000000               14.9 ns/op             0 B/op          0 allocs/op
BenchmarkRush32Max              200000000                8.56 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv32Mid           100000000               10.9 ns/op             0 B/op          0 allocs/op
BenchmarkRush32Mid              200000000                8.09 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv32Digit         300000000                5.67 ns/op            0 B/op          0 allocs/op
BenchmarkRush32Digit            1000000000               2.21 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint16()**
```
BenchmarkStrconv16Max           200000000                9.92 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Max              300000000                5.84 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv16Mid           200000000                7.84 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Mid              200000000                6.03 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv16Digit         300000000                5.67 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Digit            1000000000               2.21 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint8()**
```
BenchmarkStrconv8Max            200000000                7.86 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Max               1000000000               2.92 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv8Mid            200000000                6.84 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Mid               1000000000               2.52 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv8Digit          300000000                5.65 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Digit             2000000000               1.77 ns/op            0 B/op          0 allocs/op
```