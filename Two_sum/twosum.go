package twosum

func Twosum(arr []int, target int) []int {

	check := make(map[int]int)

	for i, val := range arr {

		index, ok := check[target-val]

		if ok {

			return []int{i, index}
		}

		check[val] = i
	}

	return nil
}
