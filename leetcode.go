package main

import (
	"sort"
)

func hasDuplicates(nums []int) bool {
	type value struct{}
	set := make(map[int]value)
	for _, val := range nums {
		set[val] = value{}
	}
	return len(set) != len(nums)
}

func isAnagram(s string, t string) bool {
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

func twoSums(nums []int, target int) []int {

	for idx, _ := range nums {
		if idx == (len(nums) - 1) {
			break
		}
		for j := idx; j < len(nums); j++ {
			if nums[idx]+nums[j] == target {
				return []int{idx, j}
			}
		}
	}
	return []int{}

}
