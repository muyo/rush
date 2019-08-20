# chars [![GoDoc](https://godoc.org/github.com/muyo/rush/chars?status.svg)](https://godoc.org/github.com/muyo/rush/chars) 

Currently provides some functions hoisted out of an otherwise private repo due to public dependencies.

### dtou (atou/atoi)

Typed alternatives to `strconv.ParseUint(dec, 10, bitsN)` (e.g. base10 string encoded numbers) with a non-idiomatic 
API (bool returns instead of err), which happened to present some optimization opportunities.

The behaviour on overflows is also different compared to strconv - overflows are reported before syntax errors, e.g. 
input may consist of entirely invalid (non-numeric) characters but exceed expected max lengths and will be reported as 
overflow first, not bad syntax. In either case the result would be invalid regardless - this simply avoids some CPU 
cycles on iterations over invalid input.

Since those tend to be in hot paths in a server context (both for arg/request parsing and response encoding),
those nanosecond gains start adding up.

```
Suffixes:
 - max: max value for a given integer size.
 - mid: a value somewhere in the middle range (about half length of max in dec characters).
 - digit: single digit (gets rough call overhead).

Note: strconv.Atoi() included only as reference for its fast path - they do different things: 
strconv.Atoi("300") != strconv.ParseUint("300", 10, 8) after all (the latter would correctly report 
an overflow, the former would yield an int of 300, also correctly).

Platform: go 1.12, i7 4770k @ 4.4GHz; ran at 2019/08/20; code included in test file.
```

**chars.ParseUint64()**
```
BenchmarkStrconvParse64Max-8        50000000                25.5 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi64Max-8         20000000                70.6 ns/op            48 B/op          1 allocs/op
BenchmarkRush64Max-8                100000000               14.2 ns/op             0 B/op          0 allocs/op

BenchmarkStrconvParse64Mid-8        100000000               14.8 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi64Mid-8         200000000                9.83 ns/op            0 B/op          0 allocs/op
BenchmarkRush64Mid-8                200000000                8.57 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse64Digit-8      300000000                5.85 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi64Digit-8       300000000                4.29 ns/op            0 B/op          0 allocs/op
BenchmarkRush64Digit-8              1000000000               2.28 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint32()**
```
BenchmarkStrconvParse32Max-8        100000000               15.0 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi32Max-8         200000000                9.35 ns/op            0 B/op          0 allocs/op
BenchmarkRush32Max-8                200000000                8.53 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse32Mid-8        100000000               11.0 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi32Mid-8         200000000                7.32 ns/op            0 B/op          0 allocs/op
BenchmarkRush32Mid-8                200000000                7.03 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse32Digit-8      300000000                5.68 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi32Digit-8       300000000                4.04 ns/op            0 B/op          0 allocs/op
BenchmarkRush32Digit-8              1000000000               2.28 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint16()**
```
BenchmarkStrconvParse16Max-8        200000000                9.83 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi16Max-8         200000000                6.83 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Max-8                300000000                5.80 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse16Mid-8        200000000                8.07 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi6Mid-8          300000000                5.72 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Mid-8                300000000                5.63 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse16Digit-8      300000000                5.66 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi16Digit-8       300000000                4.06 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Digit-8              1000000000               2.27 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint8()**
```
BenchmarkStrconvParse8Max-8         200000000                7.83 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi8Max-8          300000000                5.77 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Max                   1000000000               2.92 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse8Mid-8         200000000                7.04 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi8Mid-8          300000000                4.79 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Mid                   1000000000               2.52 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse8Digit-8       300000000                5.66 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi8Digit-8        300000000                4.05 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Digit                 2000000000               1.77 ns/op            0 B/op          0 allocs/op
```