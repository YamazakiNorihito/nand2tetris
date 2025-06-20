package symboltable

import "maps"

type ISymbolTable interface {
	AddEntry(symbol string, address int)
	AddVariable(symbol string)
	Contains(symbol string) bool
	GetAddress(symbol string) int
}

type SymbolTable struct {
	table          map[string]int
	nextRamAddress int
}

const variableStartAddr = 16

func NewSymbolTable() ISymbolTable {
	tableCopy := make(map[string]int, len(predefinedSymbols))
	maps.Copy(tableCopy, predefinedSymbols)
	return &SymbolTable{
		table:          tableCopy,
		nextRamAddress: variableStartAddr,
	}
}

func (s *SymbolTable) AddEntry(symbol string, address int) {
	s.table[symbol] = address
}

// AddVariable adds a new variable symbol to the table and assigns it the next available RAM address.
// If the symbol already exists, it does nothing.
func (s *SymbolTable) AddVariable(symbol string) {
	if !s.Contains(symbol) {
		s.AddEntry(symbol, s.nextRamAddress)
		s.nextRamAddress++
	}
}

func (s *SymbolTable) Contains(symbol string) bool {
	_, ok := s.table[symbol]
	return ok
}

func (s *SymbolTable) GetAddress(symbol string) int {
	return s.table[symbol]
}

var (
	// predefinedSymbols contains the predefined symbols for the Hack assembly language.
	predefinedSymbols = map[string]int{
		// Data Register
		"R0":  0,
		"R1":  1,
		"R2":  2,
		"R3":  3,
		"R4":  4,
		"R5":  5,
		"R6":  6,
		"R7":  7,
		"R8":  8,
		"R9":  9,
		"R10": 10,
		"R11": 11,
		"R12": 12,
		"R13": 13,
		"R14": 14,
		"R15": 15,

		"SCREEN": 16384,
		"KBD":    24576,

		"SP":   0,
		"LCL":  1,
		"ARG":  2,
		"THIS": 3,
		"THAT": 4,
	}
)
