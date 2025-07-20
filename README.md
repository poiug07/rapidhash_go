# rapidhash_go

This is pure go implementation of Rapidhash using only safe code. Original work is optimized in c/c++. This is not the fastest hash implementation available, but it should be fast. Implemented more for completeness.

Previous version implementation can be found at [V1 branch](https://github.com/poiug07/rapidhash_go/tree/v1).

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
BenchmarkRapidhash/input_size_0-12              283548844                4.211 ns/op
BenchmarkRapidhash/input_size_3-12              268181721                4.473 ns/op
BenchmarkRapidhash/input_size_4-12              269034843                4.484 ns/op
BenchmarkRapidhash/input_size_5-12              268278290                4.475 ns/op
BenchmarkRapidhash/input_size_16-12             266230188                4.507 ns/op
BenchmarkRapidhash/input_size_24-12             157922896                7.553 ns/op
BenchmarkRapidhash/input_size_32-12             158540185                7.546 ns/op
BenchmarkRapidhash/input_size_64-12             100000000               10.66 ns/op
BenchmarkRapidhash/input_size_127-12            67785726                17.47 ns/op
BenchmarkRapidhash/input_size_128-12            64114759                17.65 ns/op
BenchmarkRapidhash/input_size_129-12            65093748                18.34 ns/op
BenchmarkRapidhash/input_size_256-12            51372142                25.27 ns/op
BenchmarkRapidhash/input_size_2048-12            7802228               154.7 ns/op
BenchmarkRapidhashMicro/input_size_0-12         302791624                3.979 ns/op
BenchmarkRapidhashMicro/input_size_3-12         274381105                4.341 ns/op
BenchmarkRapidhashMicro/input_size_4-12         273853056                4.454 ns/op
BenchmarkRapidhashMicro/input_size_5-12         272695909                4.431 ns/op
BenchmarkRapidhashMicro/input_size_16-12        277723452                4.329 ns/op
BenchmarkRapidhashMicro/input_size_24-12        176537156                6.707 ns/op
BenchmarkRapidhashMicro/input_size_32-12        177444130                6.720 ns/op
BenchmarkRapidhashMicro/input_size_64-12        124237419                9.675 ns/op
BenchmarkRapidhashMicro/input_size_127-12       73958985                16.32 ns/op
BenchmarkRapidhashMicro/input_size_128-12       73906606                16.31 ns/op
BenchmarkRapidhashMicro/input_size_129-12       68889540                17.78 ns/op
BenchmarkRapidhashMicro/input_size_256-12       42390596                28.73 ns/op
BenchmarkRapidhashMicro/input_size_2048-12       6009176               195.3 ns/op
BenchmarkRapidhashNano/input_size_0-12          296407465                3.989 ns/op
BenchmarkRapidhashNano/input_size_3-12          278822640                4.312 ns/op
BenchmarkRapidhashNano/input_size_4-12          276046058                4.331 ns/op
BenchmarkRapidhashNano/input_size_5-12          268430082                4.306 ns/op
BenchmarkRapidhashNano/input_size_16-12         276705790                4.344 ns/op
BenchmarkRapidhashNano/input_size_24-12         182483445                6.621 ns/op
BenchmarkRapidhashNano/input_size_32-12         183536352                6.581 ns/op
BenchmarkRapidhashNano/input_size_64-12         100000000               11.03 ns/op
BenchmarkRapidhashNano/input_size_127-12        73922395                16.48 ns/op
BenchmarkRapidhashNano/input_size_128-12        67864113                16.51 ns/op
BenchmarkRapidhashNano/input_size_129-12        67191788                18.10 ns/op
BenchmarkRapidhashNano/input_size_256-12        40040096                29.75 ns/op
BenchmarkRapidhashNano/input_size_2048-12        5677038               202.0 ns/op
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
