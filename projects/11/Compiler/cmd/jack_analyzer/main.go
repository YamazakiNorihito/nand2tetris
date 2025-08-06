package main

import (
	"encoding/xml"
	"fmt"
	"ny/nand2tetris/compiler/internal/compilation_engine"
	"ny/nand2tetris/compiler/internal/jack_tokenizer"
	Tokens "ny/nand2tetris/compiler/internal/tokens"
	"os"
	"path/filepath"
	"strings"
)

func handler(sourcePath string) {
	jackFiles := getSourceFiles(sourcePath)
	for _, filePath := range jackFiles {
		inputFile, closeInput := mustOpenFile(filePath)
		defer closeInput()

		tokenizer, err := jack_tokenizer.NewAnalysis(inputFile)
		if err != nil {
			panic(err)
		}

		tokens, err := Tokens.Build(tokenizer)
		if err != nil {
			panic(fmt.Errorf("failed to tokenize file %s: %w", filePath, err))
		}

		dumpTokensAsXml(tokens, filePath)

		outputXmlFile, closeOutputXml := mustCreateFile(strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".xml")
		defer closeOutputXml()

		compilationEngine, err := compilation_engine.New(tokens, xml.NewEncoder(outputXmlFile))
		if err != nil {
			panic(fmt.Errorf("failed to create compilation engine for file %s: %w", filePath, err))
		}
		if err := compilationEngine.CompileClass(); err != nil {
			panic(fmt.Errorf("failed to compile class in file %s: %w", filePath, err))
		}
	}

}

func main() {
	if len(os.Args) != 2 {
		panic("Usage: JackAnalyzer <source>")
	}
	sourcePath := os.Args[1]
	handler(sourcePath)
}

func dumpTokensAsXml(tokens []Tokens.IToken, filePath string) {
	outputFile, closeoutput := mustCreateFile(strings.TrimSuffix(filePath, filepath.Ext(filePath)) + "T.xml")
	defer closeoutput()
	Tokens.DumpTokensToXML(tokens, outputFile)
}
