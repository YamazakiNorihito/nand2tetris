package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
