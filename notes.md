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

5. One probably can mess with branching and loop unrolling, but I don't think it is worth doing since behavior can change with any go version. Additionally, code will lose its clarity.