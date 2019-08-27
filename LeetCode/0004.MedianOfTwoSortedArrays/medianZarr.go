/*
 * @Author: Qian Qingnian
 * @Date: 2019-08-27 21:30
 */

/*
https://leetcode.com/problems/median-of-two-sorted-arrays/
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
*/
package main

import "fmt"

func main() {
	nums1 := []int{3, 4}
	nums2 := []int{-2, -1}
	ans := solution1(nums1, nums2)
	fmt.Println(ans)
}

func solution1(nums1, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	sumLen := len1 + len2
	newNums := make([]int, 0, sumLen)
	i, j, k := 0, 0, 0
	for i < len1 && j < len2 {
		if nums1[i] <= nums2[j] {
			newNums = append(newNums, nums1[i])
			i++
		} else {
			newNums = append(newNums,  nums2[j])
			j++
		}
		k++
	}

	if i == len1 {
		newNums = append(newNums, nums2[j:]...)
	} else {
		newNums = append(newNums, nums1[i:]...)
	}
	fmt.Println(sumLen, newNums)

	mid := sumLen / 2
	ans := 0.0
	if sumLen%2 == 0 {
		ans = float64(newNums[mid]+newNums[mid-1]) / 2.0
	} else {
		ans = float64(newNums[mid])
	}
	return ans
}
