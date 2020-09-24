package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l1_31 := &ListNode{1, nil}
	l1_30 := &ListNode{0, l1_31}
	l1_29 := &ListNode{0, l1_30}
	l1_28 := &ListNode{0, l1_29}
	l1_27 := &ListNode{0, l1_28}
	l1_26 := &ListNode{0, l1_27}
	l1_25 := &ListNode{0, l1_26}
	l1_24 := &ListNode{1, l1_25}
	l1_23 := &ListNode{0, l1_24}
	l1_22 := &ListNode{0, l1_23}
	l1_21 := &ListNode{0, l1_22}
	l1_20 := &ListNode{0, l1_21}
	l1_19 := &ListNode{0, l1_20}
	l1_18 := &ListNode{0, l1_19}
	l1_17 := &ListNode{0, l1_18}
	l1_16 := &ListNode{0, l1_17}
	l1_15 := &ListNode{0, l1_16}
	l1_14 := &ListNode{0, l1_15}
	l1_13 := &ListNode{0, l1_14}
	l1_12 := &ListNode{0, l1_13}
	l1_11 := &ListNode{0, l1_12}
	l1_10 := &ListNode{0, l1_11}
	l1_9 := &ListNode{0, l1_10}
	l1_8 := &ListNode{0, l1_9}
	l1_7 := &ListNode{0, l1_8}
	l1_6 := &ListNode{0, l1_7}
	l1_5 := &ListNode{0, l1_6}
	l1_4 := &ListNode{0, l1_5}
	l1_3 := &ListNode{0, l1_4}
	l1_2 := &ListNode{0, l1_3}
	l1_head := &ListNode{1, l1_2}

	l2_3 := &ListNode{4, nil}
	l2_2 := &ListNode{6, l2_3}
	l2_head := &ListNode{5, l2_2}

	addTwoNumbers(l1_head, l2_head)

	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1 := reverse(slice2string(list2slice(l1)))
	fmt.Println("r1: ", r1)

	r2 := reverse(slice2string(list2slice(l2)))
	fmt.Println("r2: ", r2)

	total := add(r1, r2)
	fmt.Println("total: ", total)

	str := reverse(total)
	fmt.Println("str: ", str)

	s := string2slice(str)
	fmt.Println("s: ", s)

	l := slice2list(s)
	return l
}

func list2slice(l *ListNode) []int {
	s := make([]int, numberOfNode(l), numberOfNode(l))
	i := 0
	for cur := l; cur != nil; cur = cur.Next {
		s[i] = cur.Val
		i++
	}
	return s
}

func slice2list(s []int) *ListNode {
	head := &ListNode{}
	cur := head
	len := len(s)

	for i := 0; i < len; i++ {
		cur.Val = s[i]
		if i+1 < len {
			tmp := &ListNode{}
			cur.Next = tmp
			cur = tmp
		}
	}
	return head
}

func numberOfNode(l *ListNode) int {
	i := 0
	for cur := l; cur != nil; i++ {
		tmp := cur.Next
		cur = tmp
	}
	return i
}

func slice2string(s []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), ""), "[]")
}

func string2slice(s string) []int {
	n := len(s)
	b := make([]int, n)
	for i, v := range s {
		b[i] = int(v - '0')
	}
	return b
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func add(s1, s2 string) string {
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	i1, _ := strconv.ParseUint(s1, 10, 64)
	fmt.Println("i1: ", i1)
	i2, _ := strconv.ParseUint(s2, 10, 64)
	fmt.Println("i2: ", i2)
	return strconv.FormatUint(i1+i2, 10)
}
