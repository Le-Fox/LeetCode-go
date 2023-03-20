package leetcodego

import "sort"

func Average(salary []int) float64 {
	result := 0
	sort.Ints(salary)
	for i := 1; len(salary)-1 > i; i++ {
		result += salary[i]
	}
	return float64(float64(result) / float64((len(salary) - 2)))

}
