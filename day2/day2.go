package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type opcode int

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var fp = flag.String("filepath", "day2_input", "path to file")

	original := getInput(*fp)

	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			cp := make([]int, len(original))
			copy(cp, original)

			cp[1] = i
			cp[2] = j

			result := RunIntegerMachine(cp)

			if result[0] == 19690720 {
				fmt.Println("EUREKA!!", cp[1], " ", cp[2])
			}
		}
	}

	fmt.Println("Couldn't find the instructions you're looking for")
}

func RunIntegerMachine(memory []int) []int {
	skip := 0
	for idx, inputcode := range memory {
		if skip > 0 {
			skip--
			continue
		}
		switch inputcode {
		case 99:
			return memory
		case 1:
			memory[memory[idx+3]] = memory[memory[idx+1]] + memory[memory[idx+2]]
			skip = 3
		case 2:
			memory[memory[idx+3]] = memory[memory[idx+1]] * memory[memory[idx+2]]
			skip = 3
		default:
			panic("wtf are you doing here")
		}
	}

	return memory
}

func getInput(filepath string) []int {
	f, err := os.Open(filepath)
	defer f.Close()

	var input string

	checkErr(err)
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		checkErr(err)

		input = string(line)
	}

	if input == "" {
		panic("failed to parse file...")
	}

	inputStrs := strings.Split(strings.Trim(input, " "), ",")
	res := []int{}
	for _, i := range inputStrs {
		smth, err := strconv.Atoi(i)
		checkErr(err)

		res = append(res, smth)
	}

	return res
}
