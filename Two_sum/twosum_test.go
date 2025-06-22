package twosum

import (
	"reflect"
	"testing"
)

func TestTwosum(t *testing.T) {

	tests := []struct {
		array         []int
		target        int
		expectedvalue []int
	}{

		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		ans := Twosum(tt.array, tt.target)
		if !reflect.DeepEqual(ans, tt.expectedvalue) {
			t.Errorf("For input %v and target %d, expected %v but got %v", tt.array, tt.target, tt.expectedvalue, ans)
		}

	}

}
