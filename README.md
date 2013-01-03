fizzbuzz
========

a simple golang testing and benchmark example


get it!
----
```
git clone https://github.com/dearing/fizzbuzz
```

of testing
---
Just run `go test` to see the testing package at work.  Note there is no compiled 'cmd' here for you to run directly, instead this is a complete package.

of benchmarks
---
Execute `go test -bench=".*"` to see the testing package run benchmarks on all differant kinds of fizzbuzz functions.

```
BenchmarkType1  500000000                3.88 ns/op
BenchmarkType2  500000000                5.25 ns/op
BenchmarkType3  500000000                5.96 ns/op
BenchmarkType4  100000000               18.2 ns/op
BenchmarkType5  200000000                8.51 ns/op
BenchmarkType6  500000000                4.12 ns/op
```
Yea, [Type1](https://github.com/dearing/fizzbuzz/blob/master/fizzbuzz.go#L10) is best
