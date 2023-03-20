package leetcodego

func isIsomorphic(s string, t string) bool {

	sMap := make(map[string]string)
	tMap := make(map[string]string)

	if len(s) != len(t) {
		return false
	}

	for i := 0; len(s) > i; i++ {
		sValue, sIsTrue := sMap[string(s[i])]
		tValue, tIsTrue := tMap[string(t[i])]

		if sIsTrue == true && sValue != string(t[i]) || tIsTrue == true && tValue != string(s[i]) {
			return false
		} else {
			sMap[string(s[i])] = string(t[i])
			tMap[string(t[i])] = string(s[i])
		}

	}

	return true

}
