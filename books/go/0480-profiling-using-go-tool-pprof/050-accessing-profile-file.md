---
Title: Accessing Profile File
Id: 263
Score: 0
SOId: 25549
---
once a prof file has been generated, one can access the prof file using **go tools**:

> go tool pprof cpu.prof

This will enter into a command line interface for exploring the `profile`

Common commands include:

    (pprof) top

lists top processes in memory

    (pprof) peek

Lists all processes, use *regex* to narrow search.

    (pprof) web

Opens an graph (in svg format) of the process.

An example of the `top` command:

```text
69.29s of 100.84s total (68.71%)
Dropped 176 nodes (cum <= 0.50s)
Showing top 10 nodes out of 73 (cum >= 12.03s)
        flat  flat%   sum%        cum   cum%
    12.44s 12.34% 12.34%     27.87s 27.64%  runtime.mapaccess1
    10.94s 10.85% 23.19%     10.94s 10.85%  runtime.duffcopy
        9.45s  9.37% 32.56%     54.61s 54.16%  github.com/tester/test.(*Circle).Draw
        8.88s  8.81% 41.36%      8.88s  8.81%  runtime.aeshashbody
        7.90s  7.83% 49.20%     11.04s 10.95%  runtime.mapaccess1_fast64
        5.86s  5.81% 55.01%      9.59s  9.51%  github.com/tester/test.(*Circle).isCircle
        5.03s  4.99% 60.00%      8.89s  8.82%  github.com/tester/test.(*Circle).openCircle
        3.14s  3.11% 63.11%      3.14s  3.11%  runtime.aeshash64
        3.08s  3.05% 66.16%      7.85s  7.78%  runtime.mallocgc
        2.57s  2.55% 68.71%     12.03s 11.93%  runtime.memhash
```
