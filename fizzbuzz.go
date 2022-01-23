// Package fizzbuzz is all about demonstrating and testing GO in various ways
package fizzbuzz

// Constants to demonstrate the challenge
const (
	FIZZ     = 3  // multiples of 3 yield fizz
	BUZZ     = 5  // multiples of 5 yield buzz
	FIZZBUZZ = 15 // finally, multiples of 15 yield fizzbuzz
)

// Type1 is the most performant I could come up with
func Type1(i int) string {

	if i%FIZZ == 0 {
		if i%BUZZ == 0 {
			return "FIZZBUZZ"
		}
		return "FIZZ"
	}

	if i%BUZZ == 0 {
		return "BUZZ"
	}

	return string(rune(i))
}

// Type2 is clearer than Type1
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

	return string(rune(i))
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
		}
		return "BUZZ"
	}

	if p {
		return "FIZZ"
	}
	return string(rune(i))

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
		return string(rune(i))
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

	return string(rune(i))
}

// Type6 is exactly like Type1 but works with without the constants defined before
func Type6(i int) string {

	if i%3 == 0 {
		if i%5 == 0 {
			return "FIZZBUZZ"
		}
		return "FIZZ"
	}

	if i%5 == 0 {
		return "BUZZ"
	}

	return string(rune(i))
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

	return string(rune(i))
}

// Type8 does the same as type5 but with a named return
func Type8(i int) (result string) {
	switch {

	case i%FIZZ == 0 && i%BUZZ == 0:
		result = "FIZZBUZZ"

	case i%FIZZ == 0:
		result = "FIZZ"

	case i%BUZZ == 0:
		result = "BUZZ"

	default:
		result = string(rune(i))
	}

	return
}
