package intersectionoftwoarrays

func Intersect(nums1 []int, nums2 []int) []int {

	temp := []int{}
	check := make(map[int]int)
	for _, v := range nums1 {

		check[v]++

	}

	for _, v := range nums2 {

		if val, ok := check[v]; ok && val != 0 {

			temp = append(temp, v)
			check[v]--
		}
	}

	return temp
}
