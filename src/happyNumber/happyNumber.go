package happyNumber

import (
	"math"
)

func getDigits(n int) []int {
	var res []int

	for n > 0 {
		res = append(res, n%10)
		n = n / 10
	}

	return res
}

func getDigitsSum(n int) int {
	sum := 0
	nDigits := getDigits(n)
	for i := 0; i < len(nDigits); i++ {
		sum += int(math.Pow(float64(nDigits[i]), 2))
	}
	return sum
}

func isHappyRecursive(n int, previousValues map[int]bool) bool {
	if n == 1 {
		return true
	}

	sum := getDigitsSum(n)
	if _, ok := previousValues[sum]; ok {
		return false
	} else {
		previousValues[sum] = true
	}

	return isHappyRecursive(sum, previousValues)
}

func IsHappyRecursive(n int) bool {
	return isHappyRecursive(n, map[int]bool{})
}

func IsHappyIterative(n int) bool {
	previousValues := map[int]bool{}

	for n != 1 {
		sum := getDigitsSum(n)
		if _, ok := previousValues[sum]; ok {
			return false
		} else {
			previousValues[sum] = true
		}
		n = sum
	}

	return true
}
