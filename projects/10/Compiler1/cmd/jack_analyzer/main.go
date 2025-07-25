package main

import (
	"encoding/xml"
	"fmt"
	"ny/nand2tetris/compiler1/internal/compilation_engine"
	"ny/nand2tetris/compiler1/internal/jack_tokenizer"
	Tokens "ny/nand2tetris/compiler1/internal/tokens"
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

func dumpTokensAsXml(tokens []Tokens.Token, filePath string) {
	type node struct {
		XMLName xml.Name
		Value   string `xml:",chardata"`
	}
	type root struct {
		XMLName xml.Name `xml:"tokens"`
		Tokens  []node   `xml:",any"`
	}
	paddingFunc := func(v string, padding string) string {
		return fmt.Sprintf("%s%s", padding, v)
	}

	outputFile, closeoutput := mustCreateFile(strings.TrimSuffix(filePath, filepath.Ext(filePath)) + "T.xml")
	defer closeoutput()

	xmlTokens := root{
		Tokens: make([]node, len(tokens)),
	}
	for i, token := range tokens {
		xmlTokens.Tokens[i] = node{
			XMLName: xml.Name{Local: token.GetName()},
			Value:   paddingFunc(token.GetValue(), " "),
		}
	}

	encoder := xml.NewEncoder(outputFile)
	encoder.Indent("", "  ")
	if err := encoder.Encode(xmlTokens); err != nil {
		panic(fmt.Errorf("failed to write XML: %w", err))
	}
}

func mustOpenFile(filePath string) (*os.File, func()) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to open input file %s: %w", filePath, err))
	}
	closeFunc := func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close input file %s: %v\n", filePath, cerr)
		}
	}
	return f, closeFunc
}

func mustCreateFile(filePath string) (*os.File, func()) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to create output file %s: %w", filePath, err))
	}
	closeFunc := func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close output file %s: %v\n", filePath, cerr)
		}
	}
	return f, closeFunc
}

func getSourceFiles(sourcePath string) []string {
	fi, err := os.Stat(sourcePath)
	if err != nil {
		panic(fmt.Errorf("failed to stat input path: %w", err))
	}
	var jackFiles []string
	if fi.IsDir() {
		entries, err := os.ReadDir(sourcePath)
		if err != nil {
			panic(fmt.Errorf("failed to read directory: %w", err))
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if strings.HasSuffix(strings.ToLower(name), ".jack") {
				jackFiles = append(jackFiles, filepath.Join(sourcePath, name))
			}
		}
		if len(jackFiles) == 0 {
			panic(fmt.Errorf("no .jack files found in directory: %s", sourcePath))
		}
	} else {
		if !strings.HasSuffix(strings.ToLower(sourcePath), ".jack") {
			panic(fmt.Errorf("input file must have .jack extension: %s", sourcePath))
		}
		jackFiles = append(jackFiles, sourcePath)
	}
	return jackFiles
}
