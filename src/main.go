package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	errLog := log.New(os.Stderr, "", 0)

	//Args check
	if len(args) < 2 {
		errLog.Fatalf("Usage:" + args[0] + " [URM_FILE_SOURCE] [R1, R2, ...]")
	}

	registers := args[2:]

	fmt.Println(registers)

	InstructionArray := loadInstruction(errLog, args[1])

	fmt.Println(InstructionArray)

}
