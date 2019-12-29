package day7

func NextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func GetPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func GetAllPerms(of []int) [][]int {
	results := make([][]int, 0)
	swapArray := make([]int, len(of))
	for swapArray[0] < len(swapArray) {
		results = append(results, GetPerm(of, swapArray))

		NextPerm(swapArray)
	}

	return results
}
