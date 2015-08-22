fizzbuzz 
========

[![Build Status](https://drone.io/github.com/dearing/fizzbuzz/status.png)](https://drone.io/github.com/dearing/fizzbuzz/latest)

a simple golang testing and benchmark example

try it out
----------
```
go get github.com/dearing/fizzbuzz
go test -bench=. github.com/dearing/fizzbuzz
```
- click the badge above and see it on drone.io

on testing
----------

Just run `go test` to see the testing package at work.  Note there is no compiled 'cmd' here for you to run directly, instead this is a complete package, y'know for all your fizzbuzzing needs.

of benchmarks
-------------
Execute `go test -bench=.` to see the testing package run benchmarks on all differant kinds of fizzbuzz functions.

on a i5 destop pc
-------
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

an old mobile celeron
-------
```
PASS
BenchmarkType1  20000000                85.4 ns/op
BenchmarkType2  20000000               101 ns/op
BenchmarkType3  20000000                83.2 ns/op
BenchmarkType4   5000000               645 ns/op
BenchmarkType5  20000000                99.5 ns/op
BenchmarkType6  20000000                85.3 ns/op
BenchmarkType7  20000000               101 ns/op
ok      github.com/dearing/fizzbuzz     15.603s
```

raspberry pi 2
-------
```
PASS
BenchmarkType1-4        50000000                39.0 ns/op
BenchmarkType2-4        30000000                39.0 ns/op
BenchmarkType3-4        30000000                42.3 ns/op
BenchmarkType4-4         5000000               254 ns/op
BenchmarkType5-4        20000000                68.0 ns/op
BenchmarkType6-4        50000000                39.0 ns/op
BenchmarkType7-4        50000000                56.8 ns/op
ok      github.com/dearing/fizzbuzz     13.241s
```

conclusions
-----------
Testing and Documentation are obviously important in just about *all* things in life but the trick is getting into the habit of doing the right things off the bat.  Go makes it easy by dropping it into he Go library and building support for it withint the toolchain.  Now if only I could find more practical uses for my time...

* [Type1] the a winner! ... ignoring the Type6 duplicate minus the constants.
* [Type5] is the most clear in my opinion but comes in 2nd to worst in the benchmarks.
* [Type2] would be the best choice in clarity and speed if you're that kinda person.

License and Authors
-------------------
Author: Jacob Dearing // jacob.dearing@gmail.com

```
The MIT License (MIT)

Copyright (c) 2013 Jacob Dearing

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

[Type1]: https://github.com/dearing/fizzbuzz/blob/master/fizzbuzz.go#L10-L26
[Type2]: https://github.com/dearing/fizzbuzz/blob/master/fizzbuzz.go#L30-L45
[Type5]: https://github.com/dearing/fizzbuzz/blob/master/fizzbuzz.go#L90-L104
