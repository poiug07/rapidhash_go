# Notes on improving performance using safe code

Code is already quite fast, so not much luck. Unless..., you want to write assembly or go c(pun not intended).

1. Replacing handwritten multiplication with bits library function improves performance by 4-5 times.

2. Moving slice instead of calculating offset each time using datapos improved performance. By 3-10% on all inputs except smallest ones. But it is most likely noise, since that flow is out of path.

3. Moving array elements of seed to variables improved performance for small inputs, however on larger inputs sizes there was decline. Decided that this tradeoff is not worth making. Probably optimizer does it, when necessary, anyway.

4. Remove unnecessary assignment. In golang values set to 0 when initialized anyway. This improved performance by 4-6% in the in corresponding inputs. But I don't like that code starts looking more implicit, so will leave it as comments.
```
if bufferlen > 0 {
    a = readSmall(data, bufferlen)
    /* 	b = 0 // initialized to 0 by default
    } else {
        a = 0
        b = 0 */
}   
```

5. One probably can mess with branching and more of loop unrolling, but I implemented only the ones in original implementation for completeness. Could be removed using AI quickly.

6. Tests should have tested longer inputs and across boundaries, that would have prevented some bugs that were present.

7. I used to have this:
```go
func rapid_mum(A *uint64, B *uint64) {
	hi, lo := bits.Mul64(*A, *B)
	*A = lo
	*B = hi
}
```

Changing to values, which fascilitates register use since it is pure computation and there is no memory side effects required. There is no guarantee or official mechanism. Running benchmark confirmed 15-30% improvement and assembly has no memory access in new version. Go compiler suppose to automatically recognize when pointer points to local stack variables and change it, but didn't for some reason.

```go
func rapid_mum(a, b uint64) (hi, lo uint64) {
	return bits.Mul64(a, b)
}
```

8. Every slice access incured bounds check cost. Even in manually unrolled loops, compiler was not smart enough to determine this due to indirect indexing. Adding manual BCE hints (_ = buffer[idx] ) help improve performance on bigger inputs by around 40%. Other implementations use some kind of unsafe pointer access, however, I would like to keep code readable and safe.