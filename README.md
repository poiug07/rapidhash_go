# rapidhash_go

This is pure go implementation of Rapidhash using only safe code. Original work is optimized in c/c++. This is not the fastest hash implementation available, but it should be fast.

You can read more info about `rapidhash` in original repo: [https://github.com/Nicoshev/rapidhash](https://github.com/Nicoshev/rapidhash). It is fastest hash function without problems on smhasher benchmark.

# Output comparison

To run tests you can just run `go test`.

No fancy comparison or comprehensive test set. Just some strings.
`additional/file.csv` is output of [additional/test.cpp](additional/test.cpp). I didn't include rapidhash header file, you can do that if you want.

Output seems to be matching.

# Benchmarks

This is not the fastest hash implementation, but it's ok. If you need fastest, you should probably use something else. Comparison against [xxhash](https://github.com/cespare/xxhash):
```
goos: linux
goarch: amd64
pkg: github.com/poiug07/rapidhash_go
cpu: AMD Ryzen 5 4600H with Radeon Graphics
BenchmarkRapidhash/input_size_0-12              286829593                4.234 ns/op
BenchmarkRapidhash/input_size_3-12              268587536                4.501 ns/op
BenchmarkRapidhash/input_size_4-12              260718156                4.548 ns/op
BenchmarkRapidhash/input_size_5-12              266941742                4.487 ns/op
BenchmarkRapidhash/input_size_16-12             266520846                4.475 ns/op
BenchmarkRapidhash/input_size_24-12             150312818                7.973 ns/op
BenchmarkRapidhash/input_size_32-12             151597726                7.952 ns/op
BenchmarkRapidhash/input_size_64-12             100000000               10.76 ns/op
BenchmarkRapidhash/input_size_127-12            64625989                17.91 ns/op
BenchmarkRapidhash/input_size_128-12            66186228                17.84 ns/op
BenchmarkRapidhash/input_size_129-12            64386932                18.69 ns/op
BenchmarkRapidhash/input_size_256-12            49929309                23.97 ns/op
BenchmarkRapidhash/input_size_2048-12            7851744               149.7 ns/op
BenchmarkXXHash/input_size_0-12                 383870833                3.051 ns/op
BenchmarkXXHash/input_size_3-12                 282406237                4.340 ns/op
BenchmarkXXHash/input_size_4-12                 391303644                3.092 ns/op
BenchmarkXXHash/input_size_5-12                 328387610                3.625 ns/op
BenchmarkXXHash/input_size_16-12                279238797                4.323 ns/op
BenchmarkXXHash/input_size_24-12                231869424                5.181 ns/op
BenchmarkXXHash/input_size_32-12                148488446                8.110 ns/op
BenchmarkXXHash/input_size_64-12                100000000               10.20 ns/op
BenchmarkXXHash/input_size_127-12               59162182                20.10 ns/op
BenchmarkXXHash/input_size_128-12               65578489                17.30 ns/op
BenchmarkXXHash/input_size_129-12               80339310                14.97 ns/op
BenchmarkXXHash/input_size_256-12               52242411                22.26 ns/op
BenchmarkXXHash/input_size_2048-12               8708263               136.8 ns/op
PASS
ok      github.com/poiug07/rapidhash_go 38.336s
```
