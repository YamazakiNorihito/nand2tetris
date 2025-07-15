package main

import (
	"fmt"
	"io"
	codewriter "ny/nand2tetris/basicvmtranslator/internal/code_writer"
	"ny/nand2tetris/basicvmtranslator/internal/config"
	"ny/nand2tetris/basicvmtranslator/internal/parser"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: VMTranslator <source.vm>")
		os.Exit(1)
	}

	inputFilePath := os.Args[1]

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFilePath := inputFilePath[:len(inputFilePath)-len(".vm")] + ".asm"
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	var reader io.Reader = inputFile
	var writer io.Writer = outputFile

	parser, err := parser.NewParser(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create parser: %v\n", err)
		os.Exit(1)
	}

	codewriter, err := codewriter.NewCodeWriter(writer)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create code writer: %v\n", err)
		os.Exit(1)
	}

	fileName := strings.TrimSuffix(filepath.Base(inputFilePath), ".vm")
	codewriter.SetFileName(fileName)

	for parser.HasMoreLines() {
		ok, err := parser.Advance()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to advance parser: %v\n", err)
			os.Exit(1)
		}
		if !ok {
			break
		}

		switch parser.CommandType() {
		case config.C_ARITHMETIC:
			err = codewriter.WriteArithmetic(parser.Arg1())
		case config.C_PUSH, config.C_POP:
			err = codewriter.WritePushPop(parser.CommandType(), parser.Arg1(), parser.Arg2())
		default:
			fmt.Fprintf(os.Stderr, "Unknown command type: %v\n", parser.CommandType())
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing command: %v\n", err)
			os.Exit(1)
		}
	}
}
