package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func createParametersArray(parameters []string) ([]uint, error) {
	array := []uint{}

	for _, value := range parameters {
		tmp, err := strconv.ParseUint(value, 10, 16)

		if err != nil {
			return nil, err
		}

		array = append(array, uint(tmp))
	}

	return array, nil
}

func checkInstruction(instruction rune) bool {
	for _, i := range instructionSet {
		if instruction == i {
			return true
		}
	}

	return false
}

func checkParametersNumber(instruction rune, parameters []string) bool {
	if parameters == nil {
		return false
	}

	switch instruction {
	case 'Z':
		if len(parameters) != 1 {
			return false
		}

	case 'S':
		if len(parameters) != 1 {
			return false
		}
	case 'T':
		if len(parameters) != 2 {
			return false
		}
	case 'J':
		if len(parameters) != 3 {
			return false
		}
	}

	return true
}

func cleanSplittedLine(values []string) {
	for _, el := range values {
		el = strings.TrimSpace(el)
	}
}

func createInstruction(errLog *log.Logger, line string) Instruction {
	line = strings.ReplaceAll(line, ":", "")
	line = strings.ReplaceAll(line, "(", " ")
	line = strings.ReplaceAll(line, ",", " ")
	line = strings.ReplaceAll(line, ")", "")
	splittedLine := strings.Split(line, " ")

	cleanSplittedLine(splittedLine)

	i, err := strconv.ParseUint(splittedLine[0], 10, 16)

	if err != nil {
		errLog.Panicln("Input error")
	}

	t := rune(strings.ToUpper(splittedLine[1])[0])

	if !checkInstruction(t) {
		errLog.Panicln("Input error")
	}

	parameterString := splittedLine[2:]

	if !checkParametersNumber(t, parameterString) {
		errLog.Panicln("Input error")
	}

	parameters, err := createParametersArray(parameterString)

	if err != nil {
		errLog.Panicln("Input error")
	}

	var inst Instruction
	inst.i = uint(i)
	inst.TYPE = t
	inst.parameters = parameters

	return inst
}

func loadInstruction(errLog *log.Logger, pathfile string) {
	readFile, err := os.Open(pathfile)

	if err != nil {
		errLog.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for _, eachline := range fileTextLines {
		GlobalInstructionArray = append(GlobalInstructionArray, createInstruction(errLog, eachline))
	}
}
