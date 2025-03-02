// Solving Leetcode problems in my freetime
package main

import (
	"sort"
)

// Problem: https://leetcode.com/problems/contains-duplicate/
func HasDuplicates(nums []int) bool {
	type value struct{}
	set := make(map[int]value)
	for _, val := range nums {
		set[val] = value{}
	}
	return len(set) != len(nums)
}

// Problem: https://leetcode.com/problems/valid-anagram
func IsAnagram(s string, t string) bool {
	s_rune := []rune(s)
	t_rune := []rune(t)
	sort.Slice(s_rune, func(i, j int) bool {
		return s_rune[i] < s_rune[j]
	})
	sort.Slice(t_rune, func(i, j int) bool {
		return t_rune[i] < t_rune[j]
	})

	if string(s_rune) == string(t_rune) {
		return true
	}

	return false
}

// Problem: https://leetcode.com/problems/two-sum/
func TwoSums(nums []int, target int) []int {

	for idx := range len(nums) - 1 {
		for j := idx + 1; j < len(nums); j++ {
			if (nums[idx] + nums[j]) == target {
				return []int{idx, j}
			}
		}
	}
	return []int{}
}

// Problem: https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {

	return nil
}
