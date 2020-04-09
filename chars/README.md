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

<details>
<summary>Benchmarks</summary>
<p>

```
Suffixes:
 - max: max value for a given integer size.
 - mid: a value somewhere in the middle range (about half length of max in dec characters).
 - digit: single digit (gets rough call overhead).

Note: strconv.Atoi() included only as reference for its fast path - they do different things: 
strconv.Atoi("300") != strconv.ParseUint("300", 10, 8) after all (the latter would correctly report 
an overflow, the former would yield an int of 300, also correctly).

Platform: go 1.14, i7 4770k @ 4.4GHz; ran at 2020/02/27; code included in test file.
```

**chars.ParseUint64()**
```
BenchmarkStrconvParse64Max-8        50000000                29.9 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi64Max-8         20000000                81.5 ns/op            48 B/op          1 allocs/op
BenchmarkRush64Max-8                100000000               14.2 ns/op             0 B/op          0 allocs/op

BenchmarkStrconvParse64Mid-8        100000000               17.2 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi64Mid-8         200000000                9.84 ns/op            0 B/op          0 allocs/op
BenchmarkRush64Mid-8                200000000                8.36 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse64Digit-8      300000000                6.17 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi64Digit-8       300000000                4.60 ns/op            0 B/op          0 allocs/op
BenchmarkRush64Digit-8              1000000000               2.29 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint32()**
```
BenchmarkStrconvParse32Max-8        100000000               17.3 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi32Max-8         200000000                9.86 ns/op            0 B/op          0 allocs/op
BenchmarkRush32Max-8                200000000                8.41 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse32Mid-8        100000000               12.2 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi32Mid-8         200000000                7.44 ns/op            0 B/op          0 allocs/op
BenchmarkRush32Mid-8                200000000                6.58 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse32Digit-8      300000000                6.06 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi32Digit-8       300000000                4.89 ns/op            0 B/op          0 allocs/op
BenchmarkRush32Digit-8              1000000000               2.26 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint16()**
```
BenchmarkStrconvParse16Max-8        200000000               11.1 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvAtoi16Max-8         200000000                7.33 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Max-8                300000000                5.92 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse16Mid-8        200000000                8.56 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi6Mid-8          300000000                6.31 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Mid-8                300000000                5.30 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse16Digit-8      300000000                6.04 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi16Digit-8       300000000                4.55 ns/op            0 B/op          0 allocs/op
BenchmarkRush16Digit-8              1000000000               2.27 ns/op            0 B/op          0 allocs/op
```
**chars.ParseUint8()**
```
BenchmarkStrconvParse8Max-8         200000000                9.07 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi8Max-8          300000000                6.30 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Max                   1000000000               2.91 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse8Mid-8         200000000                7.41 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi8Mid-8          300000000                5.71 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Mid                   1000000000               2.29 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvParse8Digit-8       300000000                6.65 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvAtoi8Digit-8        300000000                4.54 ns/op            0 B/op          0 allocs/op
BenchmarkRush8Digit                 2000000000               1.77 ns/op            0 B/op          0 allocs/op
```
</p>
</details>

### utod (utoa/itoa)

Zero-allocation supplements to `strconv.AppendUint()` and `strconv.FormatUint`, which instead of appending
and growing a buffer work similar to the `copy()` builtin, performing the integer-to-decimals (to ASCII bytes)
conversion in place.

<details>
<summary>Benchmarks</summary>
<p>

```
Suffixes:
 - max: max value for a given integer size.
 - mid: a value somewhere in the middle range (about half length of max in dec characters).
 - tiny: single digit (gets rough call overhead).

Platform: go 1.14, i7 4770k @ 4.4GHz; ran at 2020/02/27; code included in test file.
```

**chars.CopyUint64()**
```
BenchmarkStrconvAppendUint64-8                  39872275          30.0 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvFormatUint64-8                  24490545          50.1 ns/op            32 B/op          1 allocs/op
BenchmarkRushCopyUint64-8                       47996928          22.7 ns/op             0 B/op          0 allocs/op

BenchmarkStrconvAppendUint64Precomputed-8       210293774          5.71 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvFormatUint64Precomputed-8       364729335          3.30 ns/op            0 B/op          0 allocs/op
BenchmarkRushCopyUint64Precomputed-8            339739305          3.53 ns/op            0 B/op          0 allocs/op

BenchmarkStrconvAppendUint64Tiny-8              223835829          5.38 ns/op            0 B/op          0 allocs/op
BenchmarkStrconvFormatUint64Tiny-8              472445036          2.53 ns/op            0 B/op          0 allocs/op
BenchmarkRushCopyUint64Tiny-8                   472303716          2.54 ns/op            0 B/op          0 allocs/op
```
**chars.CopyUint32()**
```
BenchmarkStrconvAppendUint32-8      59977208                19.7 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvFormatUint32-8      32433133                37.5 ns/op            16 B/op          1 allocs/op
BenchmarkRushCopyUint32-8           85328478                13.8 ns/op             0 B/op          0 allocs/op
```
**chars.CopyUint16()**
```
BenchmarkStrconvAppendUint16-8      79951495                15.1 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvFormatUint16-8      41373888                28.7 ns/op             5 B/op          1 allocs/op
BenchmarkRushCopyUint16-8           122195684                9.65 ns/op            0 B/op          0 allocs/op
```
**chars.CopyUint8()**
```
BenchmarkStrconvAppendUint8-8       100000000               12.7 ns/op             0 B/op          0 allocs/op
BenchmarkStrconvFormatUint8-8       50002082                25.0 ns/op             3 B/op          1 allocs/op
BenchmarkRushCopyUint8-8            373325401                3.20 ns/op            0 B/op          0 allocs/op
BenchmarkRushCopyUint8Precomputed-8 394176307                3.08 ns/op            0 B/op          0 allocs/op
BenchmarkRushCopyUint8Tiny-8        481904486                2.51 ns/op            0 B/op          0 allocs/op
```
</p>
</details>
