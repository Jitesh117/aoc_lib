package aoc_lib

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(a, a%b)
}

func LCM(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	result := nums[0] * nums[1] / GCD(nums[1], nums[0])

	for i := 1; i < len(nums); i++ {
		result = LCM([]int{result, nums[i]})
	}
	return result
}

func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
