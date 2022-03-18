/**
* 6. Zigzag Conversion (medium)
**/

package main

import (
	"strings"
)

// first attempt
func convert(s string, numRows int) string {
	ret := make([]string, numRows)
	var (
		row = 0
		up  = numRows - 2
	)

	for _, c := range s {
		if row < numRows {
			ret[row] += string(c)
			row++
		} else {
			if up > 0 {
				ret[up] += string(c)
				up--
			} else {
				ret[0] += string(c)
				row = 1
				up = numRows - 2
			}
		}
	}

	return strings.Join(ret, "")
}
