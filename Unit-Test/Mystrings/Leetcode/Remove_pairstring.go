package leetcode

func Demo(s string) string {

	var index = 1

	for index < len(s) {

		if s[index] == s[index-1] {

			s = s[:index-1] + s[index+1:]

			if index != 1 {

				index--
			}

		} else {

			index++
		}

	}
	return s
}
