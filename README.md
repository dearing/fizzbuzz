fizzbuzz 
========
a simple golang testing and benchmark example

try it out
----------
```
go get github.com/dearing/fizzbuzz
go test -bench=. github.com/dearing/fizzbuzz
```
- click the badge above and see it on drone

on testing
----------

Just run `go test` to see the testing package at work.  Note there is no cli here for you to run directly, instead this is a complete package, y'know for all your fizzbuzzing needs.

of benchmarks
-------------
Execute `go test -bench=.` to see the testing package run benchmarks on all differant kinds of fizzbuzz functions.

of github actions
-----------------
Click on the little ✓ or ✗ next to the latest commit within the banner above this view to see Github actions at work running the tests and benchmarks at need.

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
