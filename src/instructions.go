package main

import (
	"log"
	"strconv"
)

//Instruction rappresents an URM simple instruction (Z, S, T, J)
type Instruction struct {
	TYPE       rune
	i          uint
	parameters []uint
}

var instructionSet = []rune{'Z', 'S', 'T', 'J'}

//NumberRegisters rappresents the number of registers which we can use
const NumberRegisters = 100

func loadRegisters(errLog *log.Logger, reg []string) {

	for i := 0; i < NumberRegisters; i = i + 1 {
		GlobalRegisters = append(GlobalRegisters, 0)
	}

	for i, val := range reg {
		tmp, err := strconv.ParseUint(val, 10, 16)

		if err != nil {
			errLog.Panicln("Input Registers Error")
		}

		GlobalRegisters[i] = uint(tmp)
	}
}

func resetInstruction(n uint) {
	GlobalRegisters[n] = 0
}

func incrementInstruction(reg uint) {
	GlobalRegisters[reg]++
}

func assignmentInstruction(m uint, n uint) {
	GlobalRegisters[n] = GlobalRegisters[m]
}

func conditionalJumpInstruction(m uint, n uint, q uint) {
	if GlobalRegisters[m] == GlobalRegisters[n] {
		k = q
	} else {
		k++
	}
}
