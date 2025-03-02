package main

import (
	"slices"
	"testing"
)

// HasDuplicates
func TestHasDuplicates(t *testing.T) {
	values := []int{1, 2, 3, 3}
	expected := true
	output := HasDuplicates(values)
	if expected != output {
		t.Fail()
	}

}

// IsAnagram
func TestIsAnagramSuccess(t *testing.T) {
	input1 := "racecar"
	input2 := "carrace"
	expected := true
	output := IsAnagram(input1, input2)
	if expected != output {
		t.Fail()
	}
}
func TestIsAnagramFailure(t *testing.T) {
	input1 := "jar"
	input2 := "jam"
	expected := false
	output := IsAnagram(input1, input2)
	if expected != output {
		t.Fail()
	}
}

// Two Sum

func TestTwoSum_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	output := TwoSums(nums, target)
	if slices.Compare(expected, output) != 0 {
		t.Fail()
	}
}
func TestTwoSum_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	output := TwoSums(nums, target)
	if slices.Compare(expected, output) != 0 {
		t.Fail()
	}
}
func TestTwoSum_3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}
	output := TwoSums(nums, target)
	if slices.Compare(expected, output) != 0 {
		t.Fail()
	}
}

func TestGroupAnagrams_1(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		[]string{"bat"},
		[]string{"nat", "tan"},
		[]string{"ate", "eat", "tea"},
	}
	output := GroupAnagrams(strs)
	if len(expected) != len(output) {
		t.FailNow()
	}
	for i, _ := range expected {
		if len(expected[i]) != len(output[i]) {
			t.FailNow()
		}
	}
}
