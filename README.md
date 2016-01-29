fizzbuzz 
========
[![Build Status](https://dearing.link/api/badges/dearing/fizzbuzz/status.svg)](https://dearing.link/dearing/fizzbuzz)

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

Just run `go test` to see the testing package at work.  Note there is no cli here for you to run directly, instead this is a complete package, y'know for all your fizzbuzzing needs.

of benchmarks
-------------
Execute `go test -bench=.` to see the testing package run benchmarks on all differant kinds of fizzbuzz functions.

*Raspberry Pi 2 Model B - go1.5 linux/arm*
-------
```
PASS
BenchmarkType1-4        50000000                39.0 ns/op
BenchmarkType2-4        30000000                39.0 ns/op
BenchmarkType3-4        30000000                42.9 ns/op
BenchmarkType4-4         5000000               252 ns/op
BenchmarkType5-4        20000000                69.0 ns/op
BenchmarkType6-4        50000000                39.0 ns/op
BenchmarkType7-4        50000000                56.8 ns/op
BenchmarkType8-4        20000000                69.0 ns/op
BenchmarkType9-4         5000000               387 ns/op
BenchmarkType9m-4        5000000               387 ns/op
ok      github.com/dearing/fizzbuzz     19.372s
```
*OnMetal Compute v1 -  go1.5 linux/amd64*
------------------
```
PASS
BenchmarkType1-20       500000000                2.95 ns/op
BenchmarkType2-20       1000000000               4.12 ns/op
BenchmarkType3-20       500000000                3.22 ns/op
BenchmarkType4-20       50000000                20.3 ns/op
BenchmarkType5-20       200000000                6.53 ns/op
BenchmarkType6-20       500000000                3.02 ns/op
BenchmarkType7-20       1000000000               4.33 ns/op
BenchmarkType8-20       200000000                6.32 ns/op
BenchmarkType9-20       50000000                26.0 ns/op
BenchmarkType9m-20      50000000                27.2 ns/op
ok      github.com/dearing/fizzbuzz     24.183s
```

conclusions
-----------
Testing and Documentation are obviously important in just about *all* things in life but the trick is getting into the habit of doing the right things off the bat.  Go makes it easy to implement testing as part of your project with benchmarking and tracing to boot.  Now if only I could find more practical uses for my time...

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
