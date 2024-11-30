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
BenchmarkRapidhash/input_size_0-12              290510910                4.132 ns/op
BenchmarkRapidhash/input_size_3-12              280600263                4.264 ns/op
BenchmarkRapidhash/input_size_4-12              186932817                6.450 ns/op
BenchmarkRapidhash/input_size_5-12              187434186                6.465 ns/op
BenchmarkRapidhash/input_size_16-12             186522740                6.433 ns/op
BenchmarkRapidhash/input_size_24-12             182301670                6.690 ns/op
BenchmarkRapidhash/input_size_32-12             182571181                6.526 ns/op
BenchmarkRapidhash/input_size_64-12             100000000               10.67 ns/op
BenchmarkRapidhash/input_size_127-12            76516664                15.82 ns/op
BenchmarkRapidhash/input_size_128-12            77193223                16.05 ns/op
BenchmarkRapidhash/input_size_129-12            67860391                17.23 ns/op
BenchmarkRapidhash/input_size_256-12            40423498                28.45 ns/op
BenchmarkRapidhash/input_size_2048-12            6230944               192.9 ns/op
BenchmarkXXHash/input_size_0-12                 385302340                3.088 ns/op
BenchmarkXXHash/input_size_3-12                 278421804                4.267 ns/op
BenchmarkXXHash/input_size_4-12                 402333718                2.989 ns/op
BenchmarkXXHash/input_size_5-12                 331480569                3.636 ns/op
BenchmarkXXHash/input_size_16-12                280389565                4.308 ns/op
BenchmarkXXHash/input_size_24-12                226634413                5.254 ns/op
BenchmarkXXHash/input_size_32-12                144552816                8.109 ns/op
BenchmarkXXHash/input_size_64-12                100000000               10.20 ns/op
BenchmarkXXHash/input_size_127-12               60816213                19.96 ns/op
BenchmarkXXHash/input_size_128-12               69258024                16.71 ns/op
BenchmarkXXHash/input_size_129-12               76140750                14.98 ns/op
BenchmarkXXHash/input_size_256-12               53192498                22.49 ns/op
BenchmarkXXHash/input_size_2048-12               8659554               140.4 ns/op
PASS
ok      github.com/poiug07/rapidhash_go 38.727s
```
