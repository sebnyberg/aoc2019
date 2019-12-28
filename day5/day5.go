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
	var debug = flag.Bool("debug", false, "whether to debug")

	flag.Parse()

	original := getInput(*fp)
	// original[225] = 5

	// original[225] = 5
	original = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	// original = []int{
	// 	3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9,
	// }

	// original[21] = 1
	ic := IntcodeComputer{
		Memory: original,
	}
	ic.Run(*debug)

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

func (c *IntcodeComputer) Add(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])
	p3 := c.Read()
	fmt.Printf("ADD [%v, %v]: (%v + %v) -> %v\n", pms[0], pms[1], p1, p2, p3)

	c.Memory[p3] = p1 + p2
}

func (c *IntcodeComputer) Multiply(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])
	p3 := c.Read()

	fmt.Printf("MULTIPLY [%v, %v]: (%v * %v) -> %v\n", pms[0], pms[1], p1, p2, p3)
	c.Memory[p3] = p1 * p2
}

func (c *IntcodeComputer) Output(pms []ParameterMode) {
	p := c.ReadParam(pms[0])
	fmt.Println()
	fmt.Println("OUTPUT:", p)
	fmt.Println()
}

func (c *IntcodeComputer) SetInput() {
	p := c.Read()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	text, err := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	checkErr(err)
	intVal, err := strconv.Atoi(text)
	checkErr(err)

	fmt.Printf("INPUT: putting %v in memory at %v\n", intVal, p)

	c.Memory[p] = intVal
}

func (c *IntcodeComputer) JumpIfTrue(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.Read()

	fmt.Printf("JUMP IF TRUE (!=0) [%v]: (%v)", pms[0], p1)
	if p1 != 0 {
		fmt.Printf(" jumping to @%v\n", p2)
		c.IP = p2
	} else {
		fmt.Printf(" staying..\n")
	}
}

func (c *IntcodeComputer) JumpIfFalse(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.Read()

	fmt.Printf("JUMP IF FALSE (==0) [%v]: (%v)", pms[0], p1)
	if p1 == 0 {
		fmt.Printf(" jumping to @%v\n", p2)
		c.IP = p2
	} else {
		fmt.Printf(" staying..\n")
	}
}

func (c *IntcodeComputer) LessThan(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])
	p3 := c.Read()
	fmt.Printf("LESS THAN [%v, %v]: (%v < %v) -> @%v = ", pms[0], pms[1], p1, p2, p3)
	if p1 < p2 {
		c.Memory[p3] = 1
		fmt.Printf("1\n")
	} else {
		c.Memory[p3] = 0
		fmt.Printf("0\n")
	}
}

func (c *IntcodeComputer) Equals(pms []ParameterMode) {
	p1 := c.ReadParam(pms[0])
	p2 := c.ReadParam(pms[1])
	p3 := c.Read()
	fmt.Printf("EQUALS [%v, %v]: (%v == %v) -> @%v = ", pms[0], pms[1], p1, p2, p3)
	if p1 == p2 {
		c.Memory[p3] = 1
		fmt.Printf("1\n")
	} else {
		c.Memory[p3] = 0
		fmt.Printf("0\n")
	}
}

func (c *IntcodeComputer) Run(debug bool) {
	for c.IP < len(c.Memory) {
		code := c.Read()
		op := OpCode(code % 100)
		pms := []ParameterMode{
			ParameterMode(code / 100 & 1),
			ParameterMode(code / 1000 & 1),
			ParameterMode(code / 10000 & 1),
		}
		if debug {
			fmt.Printf("%d\n", code)
			fmt.Print(c.Memory[:c.IP-1])
			fmt.Printf(" [%v] ", c.IP)
			fmt.Println(c.Memory[c.IP-1 : len(c.Memory)])
		}

		switch op {
		case OpCode_EXIT:
			return
		case OpCode_ADD:
			c.Add(pms)
		case OpCode_MULTIPLY:
			c.Multiply(pms)
		case OpCode_DO_OUTPUT:
			c.Output(pms)
		case OpCode_SET_INPUT_ADDR:
			c.SetInput()
		case OpCode_JUMP_IF_FALSE:
			c.JumpIfFalse(pms)
		case OpCode_JUMP_IF_TRUE:
			c.JumpIfTrue(pms)
		case OpCode_LESS_THAN:
			c.LessThan(pms)
		case OpCode_EQUALS:
			c.Equals(pms)
		default:
			panic("dafuque")
		}
	}
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
