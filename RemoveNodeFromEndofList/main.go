package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	in := newList(2)
	printList(in)

	out := removeNthFromEnd(in, 2)
	printList(out)

}

func newList(n int) *ListNode {
	head := &ListNode{}

	cur := head
	for i := 0; i <= n-1; i++ {
		cur.Val = i
		if i+1 < n {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head
}

func printList(head *ListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		println(cur.Val)
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slice := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		slice = append(slice, cur.Val)
	}
	list := remove(slice, n)
	return makeList(list)
}

func remove(slice []int, n int) []int {
	l := len(slice)
	if l == 1 && n == 1 {
		return []int{}
	}
	s := l - n
	return append(slice[:s], slice[s+1:]...)
}

func makeList(slice []int) *ListNode {
	n := len(slice)
	if n == 0 {
		return nil
	}
	head := &ListNode{}
	cur := head
	for i := 0; i <= n-1; i++ {
		cur.Val = slice[i]
		if i+1 < n {
			tmp := ListNode{
				Val:  i + 1,
				Next: nil,
			}
			cur.Next = &tmp
			cur = cur.Next
		}
	}
	return head
}
