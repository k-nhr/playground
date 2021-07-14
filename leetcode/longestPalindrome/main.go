package main

import "fmt"

func main() {
	str := longestPalindrome("abba")
	fmt.Println(str)
}

func longestPalindrome(s string) string {
	max := len(s)
	if max == 0 {
		return ""
	}

	var left, right, palindromeLen int
	var palindrome string

	for i := 0; i < max; i++ {
		left, right = i, i
		for left >= 0 && right < max && s[left] == s[right] {
			if (right - left + 1) > palindromeLen {
				palindrome = s[left : right+1]
				palindromeLen = right - left + 1
			}
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < max && s[left] == s[right] {
			if (right - left + 1) > palindromeLen {
				palindrome = s[left : right+1]
				palindromeLen = right - left + 1
			}
			left--
			right++
		}
	}
	return palindrome
}

// func fanOut(s string, left, right int) (int, int) {
// 	for left >= 0 && right < len(s) {
// 		if s[left] != s[right] {
// 			return left + 1, right - 1
// 		}
// 		left--
// 		right++
// 	}

// 	return left + 1, right - 1
// }

// func longestPalindrome(s string) string {
// 	l := len(s)
//     if l == 0 {
//         return ""
//     }
// 	var result string
// 	var left int
// 	var right int
// 	for i := 0; i < l-1; i++ {
// 		left1, right1 := fanOut(s, i, i+1)
// 		left2, right2 := fanOut(s, i, i)
// 		if right1-left1 > right2-left2 {
// 			left = left1
// 			right = right1
// 		} else {
// 			left = left2
// 			right = right2
// 		}

// 		if left < right && right-left+1 > len(result) {
// 			result = s[left : right+1]
// 		}
// 	}
// 	if len(result) == 0 {
// 		return string(s[0])
// 	}
// 	return result
// }

// func longestPalindrome(s1 string) string {
// 	l := len(s)
//     if l == 0 {
//         return ""
//     }
// 	var sub, ret string

// 	s2 := reverse(s1)
// 	for i, c1 := range s1 {
// 		for j, c2 := range s2 {
// 			if c1 == c2 {
// 				sub = getCommonSubstring(s1[i:], s2[j:])
// 				if isPalindrome(sub) {
// 					ret = longest(sub, ret)
// 				}
// 			}
// 		}
// 	}
// 	return ret
// }

// func getCommonSubstring(s1, s2 string) string {
// 	var i int

// 	n := min(len(s1), len(s2))
// 	for i = 0; i < n; i++ {
// 		if s1[i] != s2[i] {
// 			break
// 		}
// 	}
// 	return s1[:i]
// }

// func isPalindrome(s1 string) bool {
// 	return s1 == reverse(s1)
// }

// func reverse(s string) string {
// 	rns := []rune(s)
// 	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
// 		rns[i], rns[j] = rns[j], rns[i]
// 	}
// 	return string(rns)
// }

// func min(i1, i2 int) int {
// 	if i1 < i2 {
// 		return i1
// 	}
// 	return i2
// }

// func longest(s1, s2 string) string {
// 	if len(s1) > len(s2) {
// 		return s1
// 	}
// 	return s2
// }
