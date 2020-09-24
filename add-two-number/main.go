package main

import "fmt"

func main() {
	l1_3 := &ListNode{3, nil}
	l2_3 := &ListNode{9, nil}

	l1_2 := &ListNode{2, l1_3}
	l2_2 := &ListNode{8, l2_3}

	l1_head := &ListNode{1, l1_2}
	l2_head := &ListNode{7, l2_2}

	l := addTwoNumbers(l1_head, l2_head)
	fmt.Println(l)

	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := listToSlice(l1)
	s2 := listToSlice(l2)
	s := addition(s1, s2)
	l := sliceToList(s)
	return l
}

func listToSlice(l *ListNode) []int {
	s := make([]int, getListLen(l), getListLen(l))
	i := 0
	for cur := l; cur != nil; cur = cur.Next {
		s[i] = cur.Val
		i++
	}
	return s
}

func getListLen(l *ListNode) int {
	i := 0
	for cur := l; cur != nil; i++ {
		tmp := cur.Next
		cur = tmp
	}
	return i
}

func addition(s1, s2 []int) []int {
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

func sliceToList(s []int) *ListNode {
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
