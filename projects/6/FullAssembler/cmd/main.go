package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"ny/nand2tetris/fullassembler/internal/assembler"
	"ny/nand2tetris/fullassembler/internal/code"
	"ny/nand2tetris/fullassembler/internal/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: FullAssembler <inputfile.asm>")
		os.Exit(1)
	}

	inputFilePath := os.Args[1]

	p, err := parser.NewParser(inputFilePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	outputFilePath := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath)) + ".hack"
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Error: Could not create output file '%s': %v", outputFilePath, err)
	}
	defer outputFile.Close()

	c := code.NewCode()
	a := assembler.NewAssembler(c, p)

	if err := a.Assemble(outputFile); err != nil {
		log.Fatalf("Error: Assembly failed: %v", err)
	}

	fmt.Printf("Successfully assembled to '%s'.\n", outputFilePath)
}
