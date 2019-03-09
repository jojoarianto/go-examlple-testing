package main

import (
	"fmt"
)

// Addtional between 2 variable
func Additional(x int, y int) (result int) {
	result = x + y
	return result
}

// Subtraction between 2 variabel (first - second)
func Subtraction(x int, y int) (result int) {
	result = x - y
	return result
}

// angkaToRomawi function to convert int to romawi number
func AngkaToRomawi(i int) string {
	romawi := ""

	for i > 0 {
		if i >= 5000 {
			romawi += "IƆƆ"
			i -= 5000
			continue
		}

		if i >= 1000 {
			romawi += "M"
			i -= 1000
			continue
		}

		if i >= 900 {
			romawi += "CM"
			i -= 900
			continue
		}

		if i >= 500 {
			romawi += "D"
			i -= 500
			continue
		}

		if i >= 400 {
			romawi += "CD"
			i -= 400
			continue
		}

		if i >= 100 {
			romawi += "C"
			i -= 100
			continue
		}

		if i >= 90 {
			romawi += "XC"
			i -= 90
			continue
		}

		if i >= 50 {
			romawi += "L"
			i -= 50
			continue
		}

		if i >= 40 {
			romawi += "XL"
			i -= 40
			continue
		}

		if i >= 10 {
			romawi += "X"
			i -= 10
			continue
		}

		if i >= 9 {
			romawi += "IX"
			i -= 9
			continue
		}

		if i >= 5 {
			romawi += "V"
			i -= 5
			continue
		}

		if i >= 4 {
			romawi += "IV"
			i -= 4
			continue
		}

		if i >= 1 {
			romawi += "I"
			i -= 1
			continue
		}
	}

	return romawi
}

func main() {
	fmt.Println("Hello World")
}
