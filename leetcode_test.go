package main

import (
	"slices"
	"testing"
)

// hasDuplicates
func TestHasDuplicates(t *testing.T) {
	values := []int{1, 2, 3, 3}
	expected := true
	output := hasDuplicates(values)
	if expected != output {
		t.Fatalf("tester() failed")
	}

}

// isAnagram
func TestIsAnagramSuccess(t *testing.T) {
	input1 := "racecar"
	input2 := "carrace"
	expected := true
	output := isAnagram(input1, input2)
	if expected != output {
		t.Fail()
	}
}
func TestIsAnagramFailure(t *testing.T) {
	input1 := "jar"
	input2 := "jam"
	expected := false
	output := isAnagram(input1, input2)
	if expected != output {
		t.Fail()
	}
}

// Two Sum
func TestTwoSum_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	output := twoSums(nums, target)
	if slices.Compare(expected, output) != 0 {
		t.Fail()
	}
}
func TestTwoSum_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	output := twoSums(nums, target)
	if slices.Compare(expected, output) != 0 {
		t.Fail()
	}
}
func TestTwoSum_3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}
	output := twoSums(nums, target)
	if slices.Compare(expected, output) != 0 {
		t.Fail()
	}
}
