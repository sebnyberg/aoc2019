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

type OpCode int

const (
	OpCode_ADD            OpCode = 1
	OpCode_MULTIPLY       OpCode = 2
	OpCode_SET_INPUT_ADDR OpCode = 3
	OpCode_DO_OUTPUT      OpCode = 4
	OpCode_JUMP_IF_TRUE   OpCode = 5
	OpCode_JUMP_IF_FALSE  OpCode = 6
	OpCode_LESS_THAN      OpCode = 7
	OpCode_EQUALS         OpCode = 8
	OpCode_EXIT           OpCode = 99
)

var opcodeMap = map[OpCode]string{
	OpCode_ADD:            "add",
	OpCode_MULTIPLY:       "mul",
	OpCode_SET_INPUT_ADDR: "input",
	OpCode_DO_OUTPUT:      "output",
	OpCode_JUMP_IF_TRUE:   "jump if true",
	OpCode_JUMP_IF_FALSE:  "jump if false",
	OpCode_LESS_THAN:      "lt",
	OpCode_EQUALS:         "eq",
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
	IP        int
}

type Instruction struct {
	Operation      OpCode
	ParameterModes [3]ParameterMode
}

func (c *IntcodeComputer) Read() int {
	val := c.Memory[c.IP]
	c.IP++
	return val
}

func (c *IntcodeComputer) ReadParam(pm ParameterMode) int {
	p := c.Read()
	if pm != ParameterMode_POSITION {
		return p
	}
	return c.Memory[p]
}

func (c *IntcodeComputer) Add(pm1 ParameterMode, pm2 ParameterMode) {
	p1 := c.ReadParam(pm1)
	p2 := c.ReadParam(pm2)
	p3 := c.Read()

	fmt.Printf("Adding numbers %v and %v together and putting it in memory at address index %v\n", p1, p2, p3)
	c.Memory[p3] = p1 + p2
}

func (c *IntcodeComputer) Multiply(pm1 ParameterMode, pm2 ParameterMode) {
	p1 := c.ReadParam(pm1)
	p2 := c.ReadParam(pm2)
	p3 := c.Read()

	fmt.Printf("Multiplying numbers %v and %v together and putting it in memory at address index %v\n", p1, p2, p3)
	c.Memory[p3] = p1 * p2
}

func (c *IntcodeComputer) Output(pm ParameterMode) {
	p := c.ReadParam(pm)

	c.Memory[c.InputAddr] = p
	fmt.Printf("Outputting number %v and putting it in memory at address index %v\n", p, c.InputAddr)
}

func (c *IntcodeComputer) SetInput() {
	p := c.Read()
	fmt.Printf("Setting input address to be %v\n", c.Memory[p])
	c.InputAddr = c.Memory[p]
}

func (c *IntcodeComputer) JumpIfTrue(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])

	fmt.Printf("Jumping to %v if %v is non-zero...\n", p1, p2)
	if p1 != 0 {
		c.IP = p2
	}
}

func (c *IntcodeComputer) JumpIfFalse(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])

	fmt.Printf("Jumping to %v if %v is zero...\n", p1, p2)
	if p1 == 0 {
		c.IP = p2
	}
}

func (c *IntcodeComputer) LessThan(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])
	p3 := c.ReadParam(pms[2])
	fmt.Printf("If %v is less than %v, stores 1 in %v, otherwise 0\n", p1, p2, p3)
	if p1 < p2 {
		c.Memory[p3] = 1
	} else {
		c.Memory[p3] = 0
	}
}

func (c *IntcodeComputer) Equals(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])
	p3 := c.ReadParam(pms[2])
	fmt.Printf("If %v equals %v, stores 1 in %v, otherwise 0\n", p1, p2, p3)
	if p1 == p2 {
		c.Memory[p3] = 1
	} else {
		c.Memory[p3] = 0
	}
}

func (c *IntcodeComputer) Run() {
	for c.IP < len(c.Memory) {
		op, parameterModes := c.readInstruction()
		switch op {
		case OpCode_EXIT:
			return
		case OpCode_ADD:
			fmt.Printf("op: %v, pms: %v\nmemory %v\n", op, parameterModes, c.Memory[c.IP:c.IP+3])
			c.Add(parameterModes[0], parameterModes[1])
		case OpCode_MULTIPLY:
			fmt.Printf("op: %v, pms: %v\nmemory %v\n", op, parameterModes, c.Memory[c.IP:c.IP+3])
			c.Multiply(parameterModes[0], parameterModes[1])
		case OpCode_DO_OUTPUT:
			fmt.Printf("op: %v, pms: %v\nmemory %v\n", op, parameterModes, c.Memory[c.IP:c.IP+1])
			c.Output(parameterModes[0])
		case OpCode_SET_INPUT_ADDR:
			fmt.Printf("op: %v, pms: %v\nmemory %v\n", op, parameterModes, c.Memory[c.IP:c.IP+1])
			c.SetInput()
		default:
			panic("dafuque")
		}
	}
}

func (c *IntcodeComputer) readInstruction() (OpCode, []ParameterMode) {
	code := c.Read()

	opCode := OpCode(code % 100)

	parameterModes := []ParameterMode{
		ParameterMode(code / 100 & 1),
		ParameterMode(code / 1000 & 1),
		ParameterMode(code / 10000 & 1),
	}

	switch opCode {
	case OpCode_ADD, OpCode_MULTIPLY, OpCode_DO_OUTPUT, OpCode_SET_INPUT_ADDR, OpCode_EXIT:
		break
	default:
		panic(fmt.Sprintf("failed to parse %v as instruction\n", code))
	}

	return opCode, parameterModes
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
