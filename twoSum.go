package main

import "fmt"

/**
* 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
* 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
 */
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, i}
		}
		m[v] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 2, 1}
	target := 4
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))

	nums = []int{2, 7, 11, 15}
	target = 9
	fmt.Println(twoSum(nums, target))
}
