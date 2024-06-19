package ispalindrome

func isPalindrome(x int) bool {
	// 负数或类似10、100、1000这种就直接返回false
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	// 反转后半部分的数字，和前半部分做比较
	var right int
	for x > right {
		right = right*10 + x%10
		x /= 10
	}
	// 注意前半部分和后半部分刚好相等（1221）或正好差一位（121）
	return x == right || x == right/10
}
