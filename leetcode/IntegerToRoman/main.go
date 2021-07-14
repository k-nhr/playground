package main

import "fmt"

func main() {
	num := 1994
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	if num < 1 {
		return ""
	}

	var onesPlace, tensPlace, hundredsPlace, thousandsPlace int

	thousandsPlace = num / 1000
	hundredsPlace = (num - thousandsPlace*1000) / 100
	tensPlace = (num - thousandsPlace*1000 - hundredsPlace*100) / 10
	onesPlace = num - thousandsPlace*1000 - hundredsPlace*100 - tensPlace*10

	var roman string

	for i := 0; i < thousandsPlace; i++ {
		roman += "M"
	}

	if hundredsPlace >= 5 {
		if hundredsPlace == 9 {
			roman += "CM"
		} else {
			roman += "D"
			for i := 0; i < hundredsPlace-5; i++ {
				roman += "C"
			}
		}
	} else {
		if hundredsPlace == 4 {
			roman += "CD"
		} else {
			for i := 0; i < hundredsPlace; i++ {
				roman += "C"
			}
		}
	}

	if tensPlace >= 5 {
		if tensPlace == 9 {
			roman += "XC"
		} else {
			roman += "L"
			for i := 0; i < tensPlace-5; i++ {
				roman += "X"
			}
		}
	} else {
		if tensPlace == 4 {
			roman += "XL"
		} else {
			for i := 0; i < tensPlace; i++ {
				roman += "X"
			}
		}
	}

	if onesPlace >= 5 {
		if onesPlace == 9 {
			roman += "IX"
		} else {
			roman += "V"
			for i := 0; i < onesPlace-5; i++ {
				roman += "I"
			}
		}
	} else {
		if onesPlace == 4 {
			roman += "IV"
		} else {
			for i := 0; i < onesPlace; i++ {
				roman += "I"
			}
		}
	}

	return roman
}
