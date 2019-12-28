package day4

import (
	"strconv"
)

func Part2(from int, to int) int {
	count := 0
	for i := from; i < to; i++ {
		if CheckNumber(i) {
			count++
		}
	}

	return count
}

func CheckNumber(n int) bool {
	strNum := strconv.Itoa(n)

	if len(strNum) != 6 {
		return false
	}

	strArr := [6]string{}
	numArr := [6]int{}
	for i, ch := range strNum {
		strArr[i] = string(ch)
		numArr[i], _ = strconv.Atoi(string(ch))
	}

	numCount := [10]int{}
	prev := numArr[0]
	for _, cur := range numArr {
		if prev > cur {
			return false
		}
		numCount[cur]++

		prev = cur
	}

	for _, numCount := range numCount {
		if numCount == 2 {
			return true
		}
	}

	return false
}
