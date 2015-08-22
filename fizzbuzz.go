// Package fizzbuzz is all about demonstrating and testing GO in various ways
package fizzbuzz

// Constants to demonstrate the challenge
const (
	FIZZ     = 3  // multiples of 3 yield fizz
	BUZZ     = 5  // multiples of 5 yield buzz
	FIZZBUZZ = 15 // finally, multiples of 15 yield fizzbuzz
)

var memoized map[int]string

func init() {
	memoized = make(map[int]string)
}

// Type1 is the most performant I could come up with
func Type1(i int) string {

	if i%FIZZ == 0 {
		if i%BUZZ == 0 {
			return "FIZZBUZZ"
		} else {
			return "FIZZ"
		}

	} else {
		if i%BUZZ == 0 {
			return "BUZZ"
		}
	}

	return string(i)
}

// Type2 is clearer that Type1
func Type2(i int) string {

	if i%FIZZ == 0 && i%BUZZ == 0 {
		return "FIZZBUZZ"
	}

	if i%FIZZ == 0 {
		return "FIZZ"
	}

	if i%BUZZ == 0 {
		return "BUZZ"
	}

	return string(i)
}

// Type3 tosses in a bool for working
func Type3(i int) string {
	p := false
	if i%FIZZ == 0 {
		p = true
	}
	if i%BUZZ == 0 {
		if p {
			return "FIZZBUZZ"
		} else {
			return "BUZZ"
		}
	} else {
		if p {
			return "FIZZ"
		}
	}
	return string(i)

}

// Type4 demonstrates how messing with the immutable strings is costly
func Type4(i int) string {
	p := false
	r := ""

	if i%FIZZ == 0 {
		r += "FIZZ"
		p = true
	}

	if i%BUZZ == 0 {
		r += "BUZZ"
		p = true
	}

	if !p {
		return string(i)
	}
	return r
}

// Type5 uses a switch though appears less performant than Type2
func Type5(i int) string {
	switch {

	case i%FIZZ == 0 && i%BUZZ == 0:
		return "FIZZBUZZ"

	case i%FIZZ == 0:
		return "FIZZ"

	case i%BUZZ == 0:
		return "BUZZ"
	}

	return string(i)
}

// Type6 is exactly like Type1 but works with without the constants defined before
func Type6(i int) string {

	if i%3 == 0 {
		if i%5 == 0 {
			return "FIZZBUZZ"
		} else {
			return "FIZZ"
		}

	} else {
		if i%5 == 0 {
			return "BUZZ"
		}
	}

	return string(i)
}

// Type7 is like Type1 but without the nested start
func Type7(i int) string {

	if i%FIZZBUZZ == 0 {
		return "FIZZBUZZ"
	}

	if i%FIZZ == 0 {
		return "FIZZ"
	}

	if i%BUZZ == 0 {
		return "BUZZ"
	}

	return string(i)
}

func Type8(i int) (result string) {
	switch {

	case i%FIZZ == 0 && i%BUZZ == 0:
		result = "FIZZBUZZ"

	case i%FIZZ == 0:
		result = "FIZZ"

	case i%BUZZ == 0:
		result = "BUZZ"

	default:
		result = string(i)
	}

	return
}

func Type9(i int) (result string) {

	if memoized[i] != "" {
		result = memoized[i]
		return
	}

	switch {

	case i%FIZZ == 0 && i%BUZZ == 0:
		result = "FIZZBUZZ"

	case i%FIZZ == 0:
		result = "FIZZ"

	case i%BUZZ == 0:
		result = "BUZZ"

	default:
		result = string(i)
	}

	memoized[i] = result

	return
}
