package main

/**
* 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次返回删除后数组的新长度。
*  元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 新数组的索引
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2}
	expectedNums := []int{1, 2} // 长度正确的期望答案
	checkResult(nums, expectedNums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums = []int{0, 1, 2, 3, 4}
	checkResult(nums, expectedNums)
}

func checkResult(nums []int, expectedNums []int) {
	k := removeDuplicates(nums) // 调用
	println(k == len(expectedNums))
	for i := 0; i < k; i++ {
		println(nums[i] == expectedNums[i])
	}
}
