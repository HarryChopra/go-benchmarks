# Native sort vs Bubble sort

### Native Sort `sort.Ints()`

| func                 | #op     | time/op     |
| -------------------- | ------- | ----------- |
| BenchmarkSort_10     | 8180247 | 135.3 ns/op |
| BenchmarkSort_1000   | 27494   | 41 ms/op    |
| BenchmarkSort_50000  | 310     | 3.3 s/op    |
| BenchmarkSort_100000 | 170     | 6.8 s/op    |

### Custom BubbleSort

| func                 | #op  | time/op    |
| -------------------- | ---- | ---------- |
| BenchmarkSort_10     | 100M | 10.7 ns/op |
| BenchmarkSort_1000   | 2M   | 570 ns/op  |
| BenchmarkSort_50000  | 1    | 3.3 s/op   |
| BenchmarkSort_100000 | 1    | 12.6 s/op  |

For smaller length slices, custom bubble sort performs remarkably better.
