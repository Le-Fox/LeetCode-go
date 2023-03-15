package leetcodego

func romanToInt(s string) int {

	result := 0
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			result -= romanMap[string(s[i])]
		} else {
			result += romanMap[string(s[i])]
		}
	}

	return result

}
