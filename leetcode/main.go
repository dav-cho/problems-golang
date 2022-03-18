package main

import "fmt"

func main() {
	type param struct {
		s       string
		numRows int
		ans     string
	}

	tests := []param{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for _, v := range tests {
		ret := convert(v.s, v.numRows)
		if ret != v.ans {
			fmt.Printf("Wrong Answer: %s != %s\n", ret, v.ans)
			return
		}
	}

	fmt.Println("~ OK")
}
