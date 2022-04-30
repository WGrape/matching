package permute

import (
	"strings"
)

var list [][]string

func generate(start int, targetLen int, nums []string, tmpArray []string) {
	if len(tmpArray) == targetLen {
		part := make([]string, len(tmpArray))
		copy(part, tmpArray)
		list = append(list, part)
		return
	}

	for i := start; i <= len(nums)-1; i++ {
		tmpArray = append(tmpArray, nums[i])
		generate(i+1, targetLen, nums, tmpArray)

		tmpArray = tmpArray[:len(tmpArray)-1]
	}
}

func Combination(nums []string) [][]string {
	list = nil

	for targetLen := len(nums); targetLen >= 1; targetLen-- {
		generate(0, targetLen, nums, make([]string, 0))
	}

	return list
}

func CombinationAndImplode(nums []string) []string {
	tmpList := Combination(nums)
	result := make([]string, 0)
	for _, v := range tmpList {
		result = append(result, strings.Join(v, ";"))
	}
	return result
}
