/*
 * @Author: Qian Qingnian
 * @Date: 2019-07-08 11:29
 */

package main

import (
	"github.com/Fishbubble/LeetCode/0002.AddTwoNumbers/algorithm"
)

func main() {
	arr1 := []int{2, 4, 3}
	//arr1 := []int{0}
	node1 := algorithm.NewListNode(arr1)
	node1.Print()
	arr2 := []int{5, 6, 4}
	node2 := algorithm.NewListNode(arr2)
	node2.Print()

	newNode := algorithm.AddTwoNumbers(node1, node2)
	newNode.Print()
}
