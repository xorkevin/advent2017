package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	puzzleInput = "day18/input.txt"
)

const (
	isnd = iota
	iset
	iadd
	imul
	imod
	ircv
	ijgz
)

type (
	// Instr is an instruction
	Instr struct {
		instr int
		args  []int
		regs  []bool
	}
)

// NewInstr creates a new Instr
func NewInstr(instr int, args []int, regs []bool) *Instr {
	return &Instr{
		instr: instr,
		args:  args,
		regs:  regs,
	}
}

// Reg returns the value of the register or the constant value if not a register
func (i *Instr) Reg(n int, c *Compute) int {
	if i.regs[n] {
		return c.Reg(i.args[n])
	}
	return i.args[n]
}

// Arg returns the value of the nth arg
func (i *Instr) Arg(n int) int {
	return i.args[n]
}

// ParseRegister encodes register to int
func ParseRegister(reg string) int {
	return int(reg[0] - 'a')
}

// ParseArgs encodes multiple args to an int and bool slice
func ParseArgs(regs string) ([]int, []bool) {
	rs := strings.Split(regs, " ")
	k := make([]int, 0, len(rs))
	kb := make([]bool, 0, len(rs))
	for _, i := range rs {
		var a int
		var b bool
		num, err := strconv.Atoi(i)
		if err == nil {
			a = num
			b = false
		} else if len(i) == 1 {
			a = ParseRegister(i)
			b = true
		} else {
			log.Fatal(err)
		}
		k = append(k, a)
		kb = append(kb, b)
	}
	return k, kb
}

// Parse encodes instruction to Instr
func Parse(line string) *Instr {
	s := strings.SplitN(line, " ", 2)
	a, b := ParseArgs(s[1])
	switch s[0] {
	case "snd":
		return NewInstr(isnd, a, b)
	case "set":
		return NewInstr(iset, a, b)
	case "add":
		return NewInstr(iadd, a, b)
	case "mul":
		return NewInstr(imul, a, b)
	case "mod":
		return NewInstr(imod, a, b)
	case "rcv":
		return NewInstr(ircv, a, b)
	case "jgz":
		return NewInstr(ijgz, a, b)
	}
	log.Fatalf("Parse error: %s does not match an instruction\n", s[0])
	return nil
}

type (
	// Compute is a construction that executes Instrs
	Compute struct {
		mode2        bool
		counter      int
		registers    map[int]int
		instructions []*Instr
		sndChan      chan<- int
		rcvChan      <-chan int
		sentMessages int
		lastPlayed   int
	}
)

// NewCompute creates a new Compute
func NewCompute(mode2 bool, programNum int, sndChan chan<- int, rcvChan <-chan int, instrs []*Instr) *Compute {
	c := &Compute{
		mode2:        mode2,
		counter:      0,
		registers:    map[int]int{},
		instructions: instrs,
		sndChan:      sndChan,
		rcvChan:      rcvChan,
		sentMessages: 0,
		lastPlayed:   0,
	}
	if mode2 {
		c.registers[ParseRegister("p")] = programNum
	}
	return c
}

// Reg returns the value of the register
func (c *Compute) Reg(regid int) int {
	if val, ok := c.registers[regid]; ok {
		return val
	}
	c.registers[regid] = 0
	return 0
}

// WriteReg writes the value of the register
func (c *Compute) WriteReg(regid int, val int) {
	c.registers[regid] = val
}

// Execute executes one instruction
func (c *Compute) Execute() bool {
	if c.mode2 {
		return c.executeMode2()
	}
	return c.executeMode1()
}

func (c *Compute) executeMode2() bool {
	if c.counter < 0 || c.counter >= len(c.instructions) {
		return true
	}
	instr := c.instructions[c.counter]
	nextInstr := c.counter + 1
	programEnd := false

	switch instr.instr {
	case isnd:
		c.sndChan <- instr.Reg(0, c)
		c.sentMessages++
	case iset:
		c.WriteReg(instr.Arg(0), instr.Reg(1, c))
	case iadd:
		c.WriteReg(instr.Arg(0), instr.Reg(0, c)+instr.Reg(1, c))
	case imul:
		c.WriteReg(instr.Arg(0), instr.Reg(0, c)*instr.Reg(1, c))
	case imod:
		c.WriteReg(instr.Arg(0), instr.Reg(0, c)%instr.Reg(1, c))
	case ircv:
		select {
		case val := <-c.rcvChan:
			c.WriteReg(instr.Arg(0), val)
		default:
			nextInstr = c.counter
			programEnd = true
		}
	case ijgz:
		if instr.Reg(0, c) > 0 {
			nextInstr = c.counter + instr.Reg(1, c)
		}
	}

	c.counter = nextInstr

	return programEnd
}

func (c *Compute) executeMode1() bool {
	if c.counter < 0 || c.counter >= len(c.instructions) {
		return true
	}
	instr := c.instructions[c.counter]
	nextInstr := c.counter + 1
	programEnd := false

	switch instr.instr {
	case isnd:
		c.lastPlayed = instr.Reg(0, c)
	case iset:
		c.WriteReg(instr.Arg(0), instr.Reg(1, c))
	case iadd:
		c.WriteReg(instr.Arg(0), instr.Reg(0, c)+instr.Reg(1, c))
	case imul:
		c.WriteReg(instr.Arg(0), instr.Reg(0, c)*instr.Reg(1, c))
	case imod:
		c.WriteReg(instr.Arg(0), instr.Reg(0, c)%instr.Reg(1, c))
	case ircv:
		if instr.Reg(0, c) != 0 {
			programEnd = true
		}
	case ijgz:
		if instr.Reg(0, c) > 0 {
			nextInstr = c.counter + instr.Reg(1, c)
		}
	}

	c.counter = nextInstr

	return programEnd
}

func main() {
	file, err := os.Open(puzzleInput)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	instrs := []*Instr{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instrs = append(instrs, Parse(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	compute := NewCompute(false, 0, nil, nil, instrs)

	for !compute.Execute() {
	}

	fmt.Println(compute.lastPlayed)

	chan0 := make(chan int, 1024)
	chan1 := make(chan int, 1024)
	program0 := NewCompute(true, 0, chan0, chan1, instrs)
	program1 := NewCompute(true, 1, chan1, chan0, instrs)

	executing1 := false
	for {
		if executing1 {
			if waiting := program1.Execute(); waiting {
				executing1 = false
				if len(chan1) == 0 {
					break
				}
			}
		} else {
			if waiting := program0.Execute(); waiting {
				executing1 = true
				if len(chan0) == 0 {
					break
				}
			}
		}
	}

	fmt.Println(program1.sentMessages)
}
