package util

import (
	"reflect"
	"testing"
)

func allDistinct(combinations [][]int) bool {
	isSeen := make([][]int, 0)
	for i := 0; i < len(combinations); i++ {
		for j := 0; j < len(isSeen); j++ {
			if reflect.DeepEqual(combinations[i], isSeen[j]) {
				return false
			}
		}
		isSeen = append(isSeen, combinations[i])
	}
	return true
}

func TestMakeCombinations(t *testing.T) {
	res := MakeCombinations(5, 3)

	if !reflect.DeepEqual(len(res), 10) {
		t.Errorf("Error: Incorrect length!")
	}

	if !allDistinct(res) {
		t.Errorf("Error: Not all distinct!")
	}
}

func TestMakeCombinations2(t *testing.T) {
	res := MakeCombinations(4, 2)

	if !reflect.DeepEqual(len(res), 6) {
		t.Errorf("Error")
	}

	if !allDistinct(res) {
		t.Errorf("Error: Not all distinct!")
	}
}

func TestAllNotDistinct(t *testing.T) {
	res := make([][]int, 0)
	res = append(res, []int{1, 2, 3})
	res = append(res, []int{1, 2, 4})
	res = append(res, []int{1, 2, 4})

	if allDistinct(res) {
		t.Errorf("Error: Not all distinct!")
	}
}
