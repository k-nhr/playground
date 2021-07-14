package main

import "fmt"

func main() {
	s := convert("ab", 1)
	fmt.Println(s)
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	z := zigzag{
		rows: make([][]rune, min(numRows, len(s))),
	}

	cur := 0
	foldBack := false
	for _, c := range s {
		z.append(c, cur)
		if cur == 0 || cur == numRows-1 {
			foldBack = !foldBack
		}
		if foldBack {
			cur++
		} else {
			cur--
		}
	}
	return z.toString()
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

type zigzag struct {
	rows [][]rune
}

func (r *zigzag) append(c rune, idx int) {
	r.rows[idx] = append(r.rows[idx], c)
}

func (r *zigzag) toString() string {
	var str string
	for _, s := range r.rows {
		str = str + string(s)
	}
	return str
}
