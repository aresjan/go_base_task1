package main

import "fmt"

/**
* 查找字符串数组中的最长公共前缀
 */
func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	prefix := str[0]
	for i := 1; i < len(str); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(str[i]) || str[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
