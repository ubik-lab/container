# List 

Package list implements a generic doubly linked list.
The code is based on `container/list`.

## Benchmark

```
BenchmarkStdList_Any-10              573           2046827 ns/op         2715978 B/op      99489 allocs/op
BenchmarkGenericList_Any-10          571           2060806 ns/op         2715979 B/op      99489 allocs/op

BenchmarkStdList_Int-10              631           1777636 ns/op         2387811 B/op      98467 allocs/op
BenchmarkGenericList_Int-10         1030           1083593 ns/op         1280053 B/op      40001 allocs/op
```