package main

//Instruction rappresents an URM simple instruction (Z, S, T, J)
type Instruction struct {
	TYPE       rune
	i          uint16
	parameters []uint16
}

var instructionSet = []rune{'Z', 'S', 'T', 'J'}

func resetInstruction() {

}

func incrementInstruction() {

}

func assignmentInstruction() {

}

func conditionalJumpInstruction() {

}
