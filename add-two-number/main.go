package main

import "fmt"

func main() {
	l1_7 := &ListNode{1, nil}
	l1_6 := &ListNode{0, l1_7}
	l1_5 := &ListNode{0, l1_6}
	l1_4 := &ListNode{0, l1_5}
	l1_3 := &ListNode{0, l1_4}
	l1_2 := &ListNode{0, l1_3}
	l1_head := &ListNode{1, l1_2}

	l2_3 := &ListNode{5, nil}
	l2_2 := &ListNode{6, l2_3}
	l2_head := &ListNode{4, l2_2}

	l := addTwoNumbers(l1_head, l2_head)
	fmt.Println(l)

	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := toSlice(l1)
	fmt.Println(s1)

	s2 := toSlice(l2)
	fmt.Println(s2)

	s := add(s1, s2)
	fmt.Println(s)

	l := toList(s)
	return l
}

func toSlice(l *ListNode) []int {
	s := make([]int, numberOfNode(l), numberOfNode(l))
	i := 0
	for cur := l; cur != nil; cur = cur.Next {
		s[i] = cur.Val
		i++
	}
	return s
}

func numberOfNode(l *ListNode) int {
	i := 0
	for cur := l; cur != nil; i++ {
		tmp := cur.Next
		cur = tmp
	}
	return i
}

func add(s1, s2 []int) []int {
	total := []int{}
	l1 := len(s1)
	l2 := len(s2)
	cnt := 0
	if l1 > l2 {
		cnt = l1
	} else {
		cnt = l2
	}
	i1, i2, curry := 0, 0, 0

	for i := 0; i < cnt; i++ {
		i1, i2 = 0, 0
		if i < l1 {
			i1 = s1[i]
		}
		if i < l2 {
			i2 = s2[i]
		}
		tmp := i1 + i2 + curry
		curry = 0
		if tmp >= 10 {
			curry++
			total = append(total, tmp-10)
		} else {
			total = append(total, tmp)
		}
	}
	if curry != 0 {
		total = append(total, curry)
	}
	return total
}

func toList(s []int) *ListNode {
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
