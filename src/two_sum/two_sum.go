package two_sum

// Time: O(n^2)
func TwoSumNSquare(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// Time: O(n^2)
func TwoSum(nums []int, target int) []int {
	table := make(map[int]int)

	if len(nums) < 2 {
		return []int{}
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := table[nums[i]]; ok {
			table[nums[i]] += 1
		} else {
			table[nums[i]] = 1
		}
	}

	firstNumber := -1
	secondNumber := -1
	for key, value := range table {
		firstNumber = target - key
		if firstNumber == key && value == 1 {
			continue
		}
		secondNumber = key

		if _, ok := table[firstNumber]; ok {
			break
		}
	}

	i := -1
	j := -1
	for k := 0; k < len(nums); k++ {
		if i == -1 && nums[k] == firstNumber {
			i = k
		} else if j == -1 && nums[k] == secondNumber {
			j = k
		}
		if i != -1 && j != -1 {
			return []int{i, j}
		}
	}

	return []int{}
}

// Time: O(n)
// Size: O(n)
func TwoSum2(nums []int, target int) []int {
	table := make(map[int][]int)

	if len(nums) < 2 {
		return []int{}
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := table[nums[i]]; ok {
			table[nums[i]] = append(table[nums[i]], i)
		} else {
			table[nums[i]] = []int{i}
		}
	}

	for mayBeSecondNumber, value := range table {
		firstNumber := target - mayBeSecondNumber
		if firstNumber == mayBeSecondNumber {
			if len(value) == 1 {
				continue
			}
		}

		if _, ok := table[firstNumber]; ok {
			i := 0
			if firstNumber == mayBeSecondNumber {
				i = 1
			}
			return []int{table[firstNumber][0], table[mayBeSecondNumber][i]}
		}
	}
	return []int{}
}
