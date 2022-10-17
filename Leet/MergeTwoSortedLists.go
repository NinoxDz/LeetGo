package Leet

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoSortedLists() {
	list1 := getList2()
	list2 := getList()

	result := mergeTwoLists(&list1, &list2)
	for nil != result.Next {
		//println("=> ", result.Val)
		result = result.Next

	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func getList() ListNode {
	rand.Seed(time.Now().UnixNano())

	p1 := ListNode{Val: 2}
	p2 := ListNode{Val: 4, Next: &p1}
	p3 := ListNode{Val: 6, Next: &p2}
	fmt.Println(p1.Val, p2.Val, p3.Val)
	return p3
}

func getList2() ListNode {
	rand.Seed(time.Now().UnixNano())

	p1 := ListNode{Val: 1}
	p2 := ListNode{Val: 3, Next: &p1}
	p3 := ListNode{Val: 5, Next: &p2}
	fmt.Println(p1.Val, p2.Val, p3.Val)
	return p3
}
