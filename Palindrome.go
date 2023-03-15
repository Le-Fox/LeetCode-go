package leetcodego

import "strconv"

func isPalindrome(x int) bool {

	if x == reverse(x) {
		return true
	}
	return false
}

func reverse(x int) int {
	var result string
	var final int
	for _, v := range strconv.Itoa(x) {
		result = string(v) + result
	}
	final, _ = strconv.Atoi(result)
	return final
}
