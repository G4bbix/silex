package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetInputFile(filepath string) (*os.File, error) {
	var inputFile *os.File
	var err error

	if filepath != "" {
		inputFile, err = os.Open(filepath)
		if err != nil {
			return nil, fmt.Errorf("fatal error: %w", err)
		}
	} else {
		inputFile = os.Stdin
	}

	return inputFile, err
}

func ReadInput(inputFile *os.File) *[]string {
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	inputFile.Close()

	return &lines
}
