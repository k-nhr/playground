package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	s := make([]string, 0)
	s = solve(n, n, s, "")
	return s
}

func solve(open, closed int, result []string, output string) []string {
	if open == 0 && closed == 0 {
		result = append(result, output)
		return result
	}

	if closed > open {
		op2 := output + ")"
		result = solve(open, closed-1, result, op2)
	}
	if open > 0 {
		op1 := output + "("
		result = solve(open-1, closed, result, op1)
	}
	return result
}
