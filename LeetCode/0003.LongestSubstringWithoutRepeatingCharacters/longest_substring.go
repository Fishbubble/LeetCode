/*
 * @Author: Qian Qingnian
 * @Date: 2019-07-04 22:06
 */

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
Given a string, find the length of the longest substring without repeating characters.

Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */
package main

import "fmt"

func main() {
	s1 := "abcabcbb"
	s2 := "b"
	s3 := "pwwkew"
	substring(s1)
	substring(s2)
	substring(s3)

}

func substring(s string) {
	ml := 0
	n := len(s)
	var i, j int
	set := make(map[uint8]int)
	for ; j<n ; j++ {
		k := s[j]
		if v, ok := set[k]; ok {
			if v > i {
				i = v
			}
		}
		set[k] = j+1
		if ml < j+1-i {
			ml = j+1-i
		}
	}
	fmt.Println(ml)
}
