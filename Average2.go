package leetcodego

func average(salary []int) float64 {
	var min, max, sum int
	for i := range salary {
		if max < salary[i] {
			max = salary[i]
		}
		sum += salary[i]
	}
	for i, e := range salary {
		if i == 0 || e < min {
			min = e
		}
	}
	sum = sum - (min + max)
	answer := float64(sum) / float64(len(salary)-2)
	return answer

}
