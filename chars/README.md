# chars [![GoDoc](https://godoc.org/github.com/muyo/rush/chars?status.svg)](https://godoc.org/github.com/muyo/rush/chars) 

Currently provides some functions hoisted out of an otherwise private repo due to public dependencies.

### dtou

Typed alternatives to `strconv.ParseUint(dec, 10, bitsN)` (e.g. base10 string encoded numbers) with a non-idiomatic 
API (bool returns instead of err), which happened to present some optimization opportunities.

Since those tend to be in hot paths in a server context (both for arg/request parsing and response encoding),
those seemingly minor performance gains start adding up. SIMD versions of ParseUint64/32 are on the table and
should provide an additional 2-3x speedup.

```
Suffixes:
 - max: max value for a given integer size.
 - mid: a value somewhere in the middle range (about half length of max in dec characters).
 - digit: single digit (gets rough call overhead).

Platform: go 1.12, i7 4770k @ 4.4GHz; ran at 2019/06/01; code included in test file.

BenchmarkStrconv8Max            200000000                7.03 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Max               500000000                3.03 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv8Mid            200000000                6.02 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Mid               1000000000               2.96 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv8Digit          300000000                5.02 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Digit             2000000000               1.87 ns/op            0 B/op          0 allocs/op

BenchmarkStrconv16Max           200000000                9.05 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Max              200000000                6.39 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv16Mid           200000000                7.04 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Mid              200000000                6.34 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv16Digit         300000000                5.18 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Digit            1000000000               2.26 ns/op            0 B/op          0 allocs/op

BenchmarkStrconv64Max           50000000                24.1 ns/op             0 B/op          0 allocs/op
BenchmarkRush64Max              100000000               16.0 ns/op             0 B/op          0 allocs/op
BenchmarkStrconv64Mid           100000000               14.1 ns/op             0 B/op          0 allocs/op
BenchmarkRush64Mid              200000000                9.73 ns/op            0 B/op          0 allocs/op
BenchmarkStrconv64Digit         300000000                5.02 ns/op            0 B/op          0 allocs/op
BenchmarkRush64Digit            1000000000               2.19 ns/op            0 B/op          0 allocs/op
```