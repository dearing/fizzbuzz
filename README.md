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
Just run `go test` to see the testing package at work.  Note there is no compiled 'cmd' here for you to run directly, instead this is a complete package, y'know for all your fizzbuzzing needs.

of benchmarks
---
Execute `go test -bench=".*"` to see the testing package run benchmarks on all differant kinds of fizzbuzz functions.

```
PASS
BenchmarkType1  500000000                3.70 ns/op
BenchmarkType2  500000000                5.47 ns/op
BenchmarkType3  500000000                5.99 ns/op
BenchmarkType4  100000000               18.3 ns/op
BenchmarkType5  200000000                8.73 ns/op
BenchmarkType6  500000000                3.68 ns/op
BenchmarkType7  500000000                5.46 ns/op
ok      github.com/dearing/fizzbuzz     19.142s
```

conclusions
----
Testing and Documentation are obviously important in just about *all* things in life but the trick is getting into the habit of doing the right things off the bat.  Go makes it easy by dropping it into he Go library and building support for it withint the toolchain.  Now if only I could find more practical uses for my time...

[Type1](https://github.com/dearing/fizzbuzz/blob/master/fizzbuzz.go#L10-L26) the a winner! Ignoreing the Type6 duplicate minus the constants.

todo
---
Create up to 50 variations of a fizzbuzzing algo to demonstrate just how crazy I can be when I cannot sleep.
