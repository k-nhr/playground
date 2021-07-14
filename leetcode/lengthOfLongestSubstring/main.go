package main

import "fmt"

func main() {

	i := lengthOfLongestSubstring("dvdf")
	fmt.Println(i)

}

func lengthOfLongestSubstring(s string) int {
	var length int
	var subString []rune

	for _, c := range s {
		if !constain(c, subString) {
			subString = append(subString, c)
		} else {
			length = max(length, len(subString))
			pos := position(c, subString)
			subString = subString[pos+1:]
			subString = append(subString, c)
		}
	}
	return max(length, len(subString))
}

func constain(target rune, r []rune) bool {
	for _, c := range r {
		if target == c {
			return true
		}
	}
	return false
}

func position(target rune, r []rune) int {
	for i, c := range r {
		if target == c {
			return i
		}
	}
	return -1
}

func max(i1, i2 int) int {
	if i1 < i2 {
		return i2
	}
	return i1
}
