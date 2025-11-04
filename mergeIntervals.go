package main

import (
	"fmt"
	"sort"
)

/**
* 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
* 然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
* 如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
 */
func merge(intervals [][]int) [][]int {
	// 边界情况处理
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，将第一个区间加入
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := merged[len(merged)-1]

		if current[0] <= lastMerged[1] {
			lastMerged[1] = max(lastMerged[1], current[1])
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

func main() {
	// 测试用例1
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("输入: %v\n", intervals1)
	fmt.Printf("输出: %v\n", merge(intervals1))
	// 预期输出: [[1 6] [8 10] [15 18]]

	// 测试用例2
	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Printf("输入: %v\n", intervals2)
	fmt.Printf("输出: %v\n", merge(intervals2))
	// 预期输出: [[1 5]]

	// 测试用例3：空输入
	intervals3 := [][]int{}
	fmt.Printf("输入: %v\n", intervals3)
	fmt.Printf("输出: %v\n", merge(intervals3))
	// 预期输出: []

	// 测试用例4：单个区间
	intervals4 := [][]int{{1, 4}}
	fmt.Printf("输入: %v\n", intervals4)
	fmt.Printf("输出: %v\n", merge(intervals4))
	// 预期输出: [[1 4]]
}
