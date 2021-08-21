
# Introduction

`go-bytebuf` is a random access buffer implementation.
It can also mark/reset read/write indexes,
which can be very useful in packet encoding/decoding scenario and many others.

> go get github.com/joesonw/go-bytebuf



# Benchmark

```shell
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/joesonw/bytebuf
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkPooledBuffer/BenchmarkPooledBuffer_512_128-12             82303             13370 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_512_512-12             95144             12129 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_512_1024-12            97935             12389 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_1024_128-12            47499             25463 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_1024_512-12            51002             23594 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_1024_1024-12           51302             23434 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_2048_128-12            24166             50067 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_2048_512-12            25639             46754 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_2048_1024-12           26004             46475 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_4096_128-12            12092            100358 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_4096_512-12            10000            110779 ns/op
BenchmarkPooledBuffer/BenchmarkPooledBuffer_4096_1024-12           12300             93992 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_512_2-12               91994             13662 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_512_4-12               82226             13287 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_1024_2-12              42008             26302 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_1024_4-12              48309             26050 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_2048_2-12              22138             59846 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_2048_4-12              20758             54824 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_4096_2-12               8649            117953 ns/op
BenchmarkDirectBuffer/BenchmarkDirectBuffer_4096_4-12              10000            106439 ns/op
BenchmarkBytesBuffer/BenchmarkBytesBuffer_512-12                   99037             12921 ns/op
BenchmarkBytesBuffer/BenchmarkBytesBuffer_1024-12                  47634             24565 ns/op
BenchmarkBytesBuffer/BenchmarkBytesBuffer_2048-12                  24796             48782 ns/op
BenchmarkBytesBuffer/BenchmarkBytesBuffer_4096-12                  12783             92306 ns/op
PASS
ok      github.com/joesonw/bytebuf      37.231s
```