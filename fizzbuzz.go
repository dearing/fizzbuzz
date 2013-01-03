// package fizzbuzz is all about demonstrating and testing GO in various ways
package fizzbuzz

const (
	FIZZ = 3
	BUZZ = 5
)

// Most performant is the what I call the "nested cliff" approach
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

// Clearer, but this barely lags behind Type1
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

// 3rd best and only because we now fiddle with a bool
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

// Worst of the bunch this approach wastes cycles messing with immutable stings
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

// Sadly the 2nd worst of the Types but most clear; order matters in this kind of switch.
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

// A copy of Type1 but without using constants...
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
