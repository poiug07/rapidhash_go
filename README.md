# rapidhash_go

This is 1-to-1 implementation of Rapidhash using mostly safe go code. Original work is optimized in c++, my implementation is not most efficient for golang.

You can read more info about `rapidhash` in original repo: [https://github.com/Nicoshev/rapidhash](https://github.com/Nicoshev/rapidhash)

TODO: benchmark and try to improve performance

# Output comparison

To run tests you can just run `go test`.

No fancy comparison or comprehensive test set. Just some strings.
`additional/file.csv` is output of [additional/test.cpp](additional/test.cpp). I didn't include rapidhash header file, you can do that if you want.

Output seems to be matching.
