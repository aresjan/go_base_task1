package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// 1. 翻转数字
	var reverse int
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}
	// 2. 判断是否相等
	return x == reverse || x == reverse/10
}

func main() {

}
