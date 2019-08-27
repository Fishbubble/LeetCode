/*
 @Author: Qian Qingnian
 @Date: 2019-07-04 22:00
*/

/*
https://leetcode.com/problems/add-two-numbers/
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
package algorithm

import "fmt"

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (r *ListNode) Print() {
	p := r
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println()
}

func NewListNode(arr []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, v := range arr {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head.Next
}
