package main

import "fmt"

var k uint = 1

const maxNumberInstructions = 10000

func executeInstruction() {

	switch GlobalInstructionArray[k-1].TYPE {
	case 'Z':
		resetInstruction(GlobalInstructionArray[k-1].parameters[0] - 1)
		k++
	case 'S':
		incrementInstruction(GlobalInstructionArray[k-1].parameters[0] - 1)
		k++
	case 'T':
		assignmentInstruction(GlobalInstructionArray[k-1].parameters[0]-1, GlobalInstructionArray[k-1].parameters[1]-1)
		k++
	case 'J':
		conditionalJumpInstruction(GlobalInstructionArray[k-1].parameters[0]-1, GlobalInstructionArray[k-1].parameters[1]-1, GlobalInstructionArray[k-1].parameters[2])
	}
}

func start() bool {
	executionCounter := 0
	var selected rune

	for {
		if k == 0 || k > uint(len(GlobalInstructionArray)) {
			return true
		}

		if executionCounter > maxNumberInstructions {
			fmt.Printf("Executed %d instructions, do you want to continue? (Y/N): ", maxNumberInstructions)
			fmt.Scanf("%c\n", &selected)

			if selected == 'y' || selected == 'Y' {
				executionCounter = 0
			} else {
				return false
			}
		}

		executeInstruction()
		executionCounter++
	}
}
