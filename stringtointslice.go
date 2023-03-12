package piscine

func StringToIntSlice(str string) []int {
	var nums []int
	for _, char := range str {
		num := int(char)
		nums = append(nums, num)
	}
	return nums
}
