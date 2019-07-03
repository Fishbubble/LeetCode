/*
 * @Author: Qian Qingnian
 * @Date: 2019-07-04 00:17
 */

/*
https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */

package main

import "fmt"

func main() {
	nums := []int{5, 11, 4, 15}
	target := 9
	except := []int{0, 2}
	result := twoSum(nums, target)
	fmt.Println(result)

	isOk := true
	if len(except)!= len(result) {
		isOk = false
	} else {
		for i, v := range except {
			if result[i] != v {
				isOk = false
				break
			}
		}
	}

	fmt.Println(isOk)
}

func twoSum(arr []int, target int) []int {
	m := make(map[int]int, len(arr))
	for i, v := range arr {
		m[v] = i
	}
	for k, i := range m {
		sub := target - k
		if j, ok := m[sub]; ok&&i!=j  {
			return []int{i, j}
		}
	}
	return nil
}