package main

import (
	"fmt"
	"log"
	"os"

	"github.com/G4bbix/go-pcre"
)

func main() {
	config, err := ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatalln("An Error occured during parsing of command line arguments -", err)
	}

	inputFile, err := GetInputFile(config.inputFilePath)
	if err != nil {
		log.Fatalln("An Error occured during the reading of the input file -", err)
	}

	inputData := ReadInput(inputFile)
	defer inputFile.Close()

	regex, err := pcre.Compile(config.regex, 0)
	if err != nil {
		log.Fatalln("An Error occured during the compilation of the regex -", err)
	}

	BuildCastStruct(config.casts, config.regex)

	for _, j := range *inputData {
		matcher := regex.MatcherString(j, 0)
		matches, captureGroupIndicies := matcher.ExtractAllString()
		fmt.Println(captureGroupIndicies)
		for _, k := range matches {
			fmt.Printf("%s \n", k)
		}
	}
	for i, j := range config.casts {
		fmt.Printf("%d - %s\n", i, j)
	}
}
