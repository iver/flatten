package main

import (
	// 	"gotest.tools/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	var list1 = List(1, 2, 3)
	var list2 = List(4, 5, 6, 7, 8)
	var list3 = List(list1, list2, 9, 0)
	t.Log("List3: ", list3)
	var flat = Flatten(list3)

	var expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	//	t.Log("Flat: ", flat)
	//	t.Log("Expected: ", expected)
	AssertEqual(t, flat, expected)
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, result []int, expected []int) {
	lenResult := len(result)
	lenExpected := len(expected)

	if lenResult != lenExpected {
		t.Errorf("Received %v items in the list, expected %v items in the list", lenResult, lenExpected)
		return
	}

	for index := 0; index < lenResult; index++ {
		if result[index] != expected[index] {
			t.Errorf("Received %v value in index %v, expected %v value in index %v", result[index], index, expected[index], index)
			return
		}
	}
	// debug.PrintStack()
}
