package piscine

func DescendAppendRange(max, min int) []int {
	var array []int
	if min < max {
		numb := max - min
		for i := 0; i < numb; i++ {
			if i == 0 {
				array = append(array, max)
			} else {
				array = append(array, max-i)
			}
		}
		return array
	} else {
		return []int{}
	}
	// test
}
