package leetcodego

func strStr(haystack string, needle string) int {
	result := -1
	for i := 0; len(haystack) > i; i++ {
		if len(needle) == len(haystack) && needle == haystack {
			result = 0
		}
		if i+len(needle) <= len(haystack) && needle == haystack[i:i+len(needle)] {
			result = i
			break
		}
	}
	return result
}
