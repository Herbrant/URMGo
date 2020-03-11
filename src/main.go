package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//GlobalInstructionArray rappresents the array of instruction to be executed
var GlobalInstructionArray []Instruction

//GlobalRegisters rappresents the array of registers
var GlobalRegisters []uint

func main() {
	args := os.Args
	errLog := log.New(os.Stderr, "", 0)

	//Args check
	if len(args) < 2 {
		errLog.Fatalf("Usage:" + args[0] + " URM_FILE_SOURCE [R1, R2, ...] []")
	}

	registerString := args[2:]
	loadInstruction(errLog, args[1])
	loadRegisters(errLog, registerString)

	fmt.Println("INSTRUCTION LIST:")
	fmt.Println(GlobalInstructionArray)
	fmt.Println("INITIAL REGISTERS")
	fmt.Println(GlobalRegisters)

	if start() == true {
		fmt.Println("FINAL REGISTERS:")
		fmt.Println(GlobalRegisters)
		fmt.Println("Output (R1): " + strconv.Itoa(int(GlobalRegisters[0])))
	} else {
		fmt.Println("COMPUTATION NOT COMPLETED")
	}
}
