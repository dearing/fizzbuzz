fizzbuzz 
========
a simple golang testing and benchmark example

try it out
----------
```
go get github.com/dearing/fizzbuzz
go test -bench=. github.com/dearing/fizzbuzz
```

on testing
----------

Just run `go test` to see the testing package at work.  Note there is no cli here for you to run directly, instead this is a complete package, y'know for all your fizzbuzzing needs.

of benchmarks
-------------
Execute `go test -bench=.` to see the testing package run benchmarks on all differant kinds of fizzbuzz functions.

conclusions
-----------
Testing and Documentation are obviously important in just about *all* things in life but the trick is getting into the habit of doing the right things off the bat.  Go makes it easy to implement testing as part of your project with benchmarking and tracing to boot.  Now if only I could find more practical uses for my time...

```
go test -bench .
goos: linux
goarch: amd64
pkg: github.com/dearing/fizzbuzz
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkType1-12       1000000000               1.204 ns/op
BenchmarkType2-12       990890618                4.961 ns/op
BenchmarkType3-12       1000000000               1.127 ns/op
BenchmarkType4-12       172770543                6.936 ns/op
BenchmarkType5-12       992789904                1.287 ns/op
BenchmarkType6-12       999073558                4.805 ns/op
BenchmarkType7-12       674727265                1.778 ns/op
BenchmarkType8-12       983222809                4.977 ns/op
PASS
ok      github.com/dearing/fizzbuzz     22.931s
```

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
