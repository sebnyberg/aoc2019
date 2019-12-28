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

type OpCode int

const (
	OpCode_ADD            OpCode = 1
	OpCode_MULTIPLY       OpCode = 2
	OpCode_SET_INPUT_ADDR OpCode = 3
	OpCode_DO_OUTPUT      OpCode = 4
	OpCode_EXIT           OpCode = 99
)

var opcodeMap = map[OpCode]string{
	OpCode_ADD:            "add",
	OpCode_MULTIPLY:       "mul",
	OpCode_SET_INPUT_ADDR: "input",
	OpCode_DO_OUTPUT:      "output",
	OpCode_EXIT:           "exit",
}

func (oc OpCode) String() string {
	return opcodeMap[oc]
}

type ParameterMode int

const (
	ParameterMode_POSITION  ParameterMode = 0
	ParameterMode_IMMEDIATE ParameterMode = 1
)

var pmMap = map[ParameterMode]string{
	ParameterMode_IMMEDIATE: "imm",
	ParameterMode_POSITION:  "pos",
}

func (pm ParameterMode) String() string {
	return pmMap[pm]
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type IntcodeComputer struct {
	Memory    []int
	InputAddr int
}

type Instruction struct {
	Operation      OpCode
	ParameterModes [3]ParameterMode
}

func (c *IntcodeComputer) Add(instr *Instruction, idx int) {
	paramVals := [2]int{}
	switch instr.ParameterModes[0] {
	case ParameterMode_POSITION:
		paramVals[0] = c.Memory[c.Memory[idx]]
	case ParameterMode_IMMEDIATE:
		paramVals[0] = c.Memory[idx]
	}

	switch instr.ParameterModes[1] {
	case ParameterMode_POSITION:
		paramVals[1] = c.Memory[c.Memory[idx+1]]
	case ParameterMode_IMMEDIATE:
		paramVals[1] = c.Memory[idx+1]
	}
	fmt.Printf("Adding numbers %v and %v together and putting it in memory at address index %v\n", paramVals[0], paramVals[1], idx+2)

	c.Memory[c.Memory[idx+2]] = paramVals[0] + paramVals[1]
}

func (c *IntcodeComputer) Multiply(instr *Instruction, idx int) {
	paramVals := [2]int{}
	switch instr.ParameterModes[0] {
	case ParameterMode_POSITION:
		paramVals[0] = c.Memory[c.Memory[idx]]
	case ParameterMode_IMMEDIATE:
		paramVals[0] = c.Memory[idx]
	}

	switch instr.ParameterModes[1] {
	case ParameterMode_POSITION:
		paramVals[1] = c.Memory[c.Memory[idx+1]]
	case ParameterMode_IMMEDIATE:
		paramVals[1] = c.Memory[idx+1]
	}

	fmt.Printf("Multiplying numbers %v and %v together and putting it in memory at address index %v\n", paramVals[0], paramVals[1], idx+2)
	c.Memory[c.Memory[idx+2]] = paramVals[0] * paramVals[1]
}

func (c *IntcodeComputer) Output(instr *Instruction, idx int) {
	var paramVal int
	switch instr.ParameterModes[0] {
	case ParameterMode_POSITION:
		paramVal = c.Memory[c.Memory[idx]]
	case ParameterMode_IMMEDIATE:
		paramVal = c.Memory[idx]
	}
	c.Memory[c.InputAddr] = paramVal
	fmt.Printf("Outputting number %v and putting it in memory at address index %v\n", paramVal, c.InputAddr)
}

func (c *IntcodeComputer) SetInput(instr *Instruction, idx int) {
	fmt.Printf("Setting input address to be %v\n", c.Memory[idx])
	c.InputAddr = c.Memory[idx]
}

func (c *IntcodeComputer) Run() {
	idx := 0
	for idx < len(c.Memory) {
		instr := NewInstruction(c.Memory[idx])
		switch instr.Operation {
		case OpCode_EXIT:
			return
		case OpCode_ADD:
			fmt.Printf("instr: %v\nmemory %v\n", instr, c.Memory[idx+1:idx+4])
			c.Add(instr, idx+1)
			idx += 4
		case OpCode_MULTIPLY:
			fmt.Printf("instr: %v\nmemory %v\n", instr, c.Memory[idx+1:idx+4])
			c.Multiply(instr, idx+1)
			idx += 4
		case OpCode_DO_OUTPUT:
			fmt.Printf("instr: %v\nmemory %v\n", instr, c.Memory[idx+1:idx+2])
			c.Output(instr, idx+1)
			idx += 2
		case OpCode_SET_INPUT_ADDR:
			fmt.Printf("instr: %v\nmemory %v\n", instr, c.Memory[idx+1:idx+2])
			c.SetInput(instr, idx+1)
			idx += 2
		default:
			panic("dafuque")
		}
	}
}

func NewInstruction(raw int) *Instruction {
	inst := new(Instruction)

	opCode := OpCode(raw % 100)

	inst.ParameterModes = [3]ParameterMode{
		ParameterMode(raw / 100 & 1),
		ParameterMode(raw / 1000 & 1),
		ParameterMode(raw / 10000 & 1),
	}

	switch opCode {
	case OpCode_ADD, OpCode_MULTIPLY, OpCode_DO_OUTPUT, OpCode_SET_INPUT_ADDR, OpCode_EXIT:
		inst.Operation = opCode
		return inst
	default:
		panic(fmt.Sprintf("failed to parse %v as instruction\n", raw))
	}
}

func main() {
	var fp = flag.String("filepath", "day5_input", "path to file")

	original := getInput(*fp)
	original[225] = 1

	ic := IntcodeComputer{
		Memory: original,
	}
	ic.Run()

	// result := RunIntegerMachine(original)
	fmt.Println(ic.Memory)
}

func RunIntegerMachine(memory []int) []int {
	// ugly fix
	memory[225] = 1
	skip := 0
	outputToAddr := -1
	for idx := range memory {
		if skip > 0 {
			skip--
			continue
		}

		instruction := memory[idx]

		// exit condition
		if OpCode(instruction) == OpCode_EXIT {
			return memory
		}

		p1 := ParameterMode(instruction / 100 & 1)
		p2 := ParameterMode(instruction / 1000 & 1)

		fmt.Printf("instruction: %v\n", instruction)
		fmt.Printf("parameter modes: p1: %v, p2: %v\n", p1, p2)
		switch OpCode(instruction % 100) {
		case OpCode_SET_INPUT_ADDR:
			fmt.Println("waiting for input...")
			outputToAddr = memory[idx+1]
			skip = 1
		case OpCode_DO_OUTPUT:
			if outputToAddr == -1 {
				panic("wth no output address???")
			}
			if p1 == ParameterMode_IMMEDIATE {
				memory[outputToAddr] = memory[idx+1]
			} else {
				memory[outputToAddr] = memory[memory[idx+1]]
			}
			skip = 1
		case OpCode_ADD:
			vals := [2]int{}
			if p1 == ParameterMode_IMMEDIATE {
				vals[0] = memory[idx+1]
			} else {
				vals[0] = memory[memory[idx+1]]
			}
			if p2 == ParameterMode_IMMEDIATE {
				vals[1] = memory[idx+2]
			} else {
				vals[1] = memory[memory[idx+2]]
			}
			fmt.Printf("adding numbers %v + %v = %v and putting in %v\n", vals[0], vals[1], vals[0]+vals[1], memory[idx+3])
			memory[memory[idx+3]] = vals[0] + vals[1]
			skip = 3
		case OpCode_MULTIPLY:
			vals := [2]int{}
			if p1 == ParameterMode_IMMEDIATE {
				vals[0] = memory[idx+1]
			} else {
				vals[0] = memory[memory[idx+1]]
			}
			if p2 == ParameterMode_IMMEDIATE {
				vals[1] = memory[idx+2]
			} else {
				vals[1] = memory[memory[idx+2]]
			}
			memory[memory[idx+3]] = vals[0] * vals[1]
			skip = 3
		default:
			panic(fmt.Sprintf("wtf are you doing here?? got opcode %v", instruction))
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
