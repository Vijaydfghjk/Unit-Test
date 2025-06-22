package intersectionoftwoarrays

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {

	testcase := []struct {
		nums1  []int
		nums2  []int
		result []int
	}{

		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
	}

	for _, tt := range testcase {

		ans := Intersect(tt.nums1, tt.nums2)

		if !reflect.DeepEqual(ans, tt.result) {

			t.Errorf("Your value %v and Expected value %v not matching", ans, tt.result)
		}
	}
}
