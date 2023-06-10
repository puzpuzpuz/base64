# base64

Base64 experiment written in Go. Not meant to be used as a production ready library.

Based in ideas from the [fast-base64](https://github.com/npodonnell/fast-base64) library.
The goal is experiment with scalar (non-SIMD) encoder/decoder and compare it with some
other Go libraries, including the standard one.

Decoder-only for now. Encoder and an algo explanation will follow.

Limitations:
* Scalar instructions and little-endian only.
* Fairly unsafe implementation along with lack of throughout testing, hence not ready for production use.
* Decoder has no input validation.

## Benchmarks

```bash
$ go test -bench=. -cpu=1
goos: linux
goarch: amd64
pkg: github.com/puzpuzpuz/base64
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkDecode_RandomBytes/std/Decode         	  844734	      1395 ns/op	1468.60 MB/s
BenchmarkDecode_RandomBytes/cristalhq/Decode   	 1490624	       805.7 ns/op	2541.74 MB/s
BenchmarkDecode_RandomBytes/puzpuzpuz/Decode   	 1578878	       759.5 ns/op	2696.40 MB/s
BenchmarkDecode_RandomAscii/std/Decode         	 1000000	      1049 ns/op	1951.51 MB/s
BenchmarkDecode_RandomAscii/cristalhq/Decode   	 2001566	       606.0 ns/op	3379.79 MB/s
BenchmarkDecode_RandomAscii/puzpuzpuz/Decode   	 2153888	       557.3 ns/op	3674.55 MB/s
PASS
ok  	github.com/puzpuzpuz/base64	9.825s
```
