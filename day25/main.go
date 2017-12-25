package main

import (
	"fmt"
)

const (
	totalSteps = 12683008
	memSize    = 100000
)

const (
	stateA = iota
	stateB
	stateC
	stateD
	stateE
	stateF
)

type (
	machine struct {
		pos   int
		state int
	}
)

func newMachine(pos int, initState int) *machine {
	return &machine{
		pos:   pos,
		state: initState,
	}
}

func (m *machine) step(mem []int) {
	nextState := m.state
	nextPos := m.pos
	switch m.state {
	case stateA:
		switch mem[m.pos] {
		case 0:
			mem[m.pos] = 1
			nextPos = m.pos + 1
			nextState = stateB
		case 1:
			mem[m.pos] = 0
			nextPos = m.pos - 1
			nextState = stateB
		}

	case stateB:
		switch mem[m.pos] {
		case 0:
			mem[m.pos] = 1
			nextPos = m.pos - 1
			nextState = stateC
		case 1:
			mem[m.pos] = 0
			nextPos = m.pos + 1
			nextState = stateE
		}

	case stateC:
		switch mem[m.pos] {
		case 0:
			mem[m.pos] = 1
			nextPos = m.pos + 1
			nextState = stateE
		case 1:
			mem[m.pos] = 0
			nextPos = m.pos - 1
			nextState = stateD
		}

	case stateD:
		switch mem[m.pos] {
		case 0:
			mem[m.pos] = 1
			nextPos = m.pos - 1
			nextState = stateA
		case 1:
			mem[m.pos] = 1
			nextPos = m.pos - 1
			nextState = stateA
		}

	case stateE:
		switch mem[m.pos] {
		case 0:
			mem[m.pos] = 0
			nextPos = m.pos + 1
			nextState = stateA
		case 1:
			mem[m.pos] = 0
			nextPos = m.pos + 1
			nextState = stateF
		}

	case stateF:
		switch mem[m.pos] {
		case 0:
			mem[m.pos] = 1
			nextPos = m.pos + 1
			nextState = stateE
		case 1:
			mem[m.pos] = 1
			nextPos = m.pos + 1
			nextState = stateA
		}
	}

	m.pos = nextPos
	m.state = nextState
}

func main() {
	mem := make([]int, memSize)
	m := newMachine(memSize/2, stateA)
	for i := 0; i < totalSteps; i++ {
		m.step(mem)
	}
	checksum := 0
	for _, i := range mem {
		if i == 1 {
			checksum++
		}
	}
	fmt.Println(checksum)
}
