package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sebnyberg/aoc2019/day7part2/intcode"
	"github.com/sebnyberg/aoc2019/util"
)

func main() {
	inputStrs := strings.Split(util.ReadFile("input")[0], ",")
	program := make([]int, len(inputStrs))
	for idx, is := range inputStrs {
		i, err := strconv.Atoi(is)
		util.CheckErr(err)
		program[idx] = i
	}

	maxThrust := 0
	maxPhases := make([]int, 5)
	for _, phases := range util.GetAllPerms([]int{5, 6, 7, 8, 9}) {
		thrust, err := intcode.GetThrust(program, phases)
		if err != nil {
			panic(err)
		}
		if thrust > maxThrust {
			maxPhases = phases
			maxThrust = thrust
		}
	}

	fmt.Println(maxThrust)
	fmt.Println(maxPhases)
}
