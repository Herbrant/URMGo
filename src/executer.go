package main

var k uint = 1

func executeInstruction() bool {

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

	return true
}

func start() {

	for {
		if k == 0 || k > uint(len(GlobalInstructionArray)) {
			break
		}

		executeInstruction()
	}
}
