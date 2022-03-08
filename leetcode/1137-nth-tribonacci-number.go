/**
* 1137. N-th Tribonacci Number (easy)
**/

package main

// fibonacci with closure
func tribonacci(n int) int {
	if n < 1 {
		return 0
	}
	if n < 3 {
		return 1
	}

	a, b, c := 0, 0, 1

	var helper func() int
	helper = func() int {
		a, b, c = b, c, a+b+c

		return c
	}

	var ret int

	for i := 1; i < n; i++ {
		ret = helper()
	}

	return ret
}
