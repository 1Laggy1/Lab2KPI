// main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	lab2 "github.com/1Laggy1/Lab2KPI"
)

func main() {
	exprPtr := flag.String("e", "", "Postfix expression to evaluate")
	filePtr := flag.String("f", "", "File containing postfix expression")
	outputPtr := flag.String("o", "", "Output file for result")
	flag.Parse()

	var reader lab2.InputReader
	var writer lab2.OutputWriter

	if *exprPtr != "" {
		reader = lab2.NewStringInputReader(*exprPtr)
	} else if *filePtr != "" {
		file, err := os.Open(*filePtr)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader = lab2.NewFileInputReader(file)
	} else {
		log.Fatal("No input provided")
	}

	if *outputPtr != "" {
		file, err := os.Create(*outputPtr)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		writer = lab2.NewFileOutputWriter(file)
	} else {
		writer = lab2.NewStdOutputWriter()
	}

	handler := lab2.ComputeHandler{
		InputReader:  reader,
		OutputWriter: writer,
	}

	err := handler.Compute()
	if err != nil {
		log.Fatal(err)
	}
}