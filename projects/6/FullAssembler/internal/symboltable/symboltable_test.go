package symboltable

import "testing"

func TestNewSymbolTable(t *testing.T) {
	st := NewSymbolTable()

	if st == nil {
		t.Fatal("NewSymbolTable() returned nil")
	}

	expectedSymbols := map[string]int{
		"R0": 0, "R1": 1, "R2": 2, "R3": 3, "R4": 4, "R5": 5, "R6": 6, "R7": 7,
		"R8": 8, "R9": 9, "R10": 10, "R11": 11, "R12": 12, "R13": 13, "R14": 14, "R15": 15,
		"SCREEN": 16384, "KBD": 24576,
		"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4,
	}

	for symbol, expectedAddress := range expectedSymbols {
		t.Run(symbol, func(t *testing.T) {
			if !st.Contains(symbol) {
				t.Errorf("expected symbol table to contain '%s', but it doesn't", symbol)
			}

			if gotAddress := st.GetAddress(symbol); gotAddress != expectedAddress {
				t.Errorf("expected address of '%s' to be %d, but got %d", symbol, expectedAddress, gotAddress)
			}
		})
	}
}

func TestAddEntry(t *testing.T) {
	t.Run("add a new symbol", func(t *testing.T) {
		st := NewSymbolTable()
		symbol := "loop"
		address := 100

		if st.Contains(symbol) {
			t.Fatalf("Symbol table should not contain '%s' before adding it", symbol)
		}

		st.AddEntry(symbol, address)

		if !st.Contains(symbol) {
			t.Errorf("Expected symbol table to contain '%s' after adding it", symbol)
		}

		if gotAddress := st.GetAddress(symbol); gotAddress != address {
			t.Errorf("Expected address of '%s' to be %d, but got %d", symbol, address, gotAddress)
		}
	})

	t.Run("overwrite a predefined symbol", func(t *testing.T) {
		st := NewSymbolTable()
		st.AddEntry("R5", 999)
		if gotAddress := st.GetAddress("R5"); gotAddress != 999 {
			t.Errorf("Expected address of 'R5' to be overwritten to 999, but got %d", gotAddress)
		}
	})
}

func TestAddVariable(t *testing.T) {
	t.Run("adds new variables with incrementing addresses", func(t *testing.T) {
		st := NewSymbolTable()
		st.AddVariable("var1")
		st.AddVariable("var2")

		if addr := st.GetAddress("var1"); addr != 16 {
			t.Errorf("Expected address for var1 to be 16, got %d", addr)
		}
		if addr := st.GetAddress("var2"); addr != 17 {
			t.Errorf("Expected address for var2 to be 17, got %d", addr)
		}
	})

	t.Run("does not add a variable if it already exists", func(t *testing.T) {
		st := NewSymbolTable()
		st.AddEntry("my_label", 100)
		st.AddVariable("my_label") // should not overwrite

		if addr := st.GetAddress("my_label"); addr != 100 {
			t.Errorf("AddVariable should not overwrite existing symbol. Expected 100, got %d", addr)
		}

		// Check that next available address is still 16
		st.AddVariable("new_var")
		if addr := st.GetAddress("new_var"); addr != 16 {
			t.Errorf("Expected next variable address to be 16, got %d", addr)
		}
	})

	t.Run("does not add a predefined symbol", func(t *testing.T) {
		st := NewSymbolTable()
		st.AddVariable("R5")

		if addr := st.GetAddress("R5"); addr != 5 {
			t.Errorf("AddVariable should not overwrite predefined symbol R5. Expected 5, got %d", addr)
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("predefined symbol", func(t *testing.T) {
		st := NewSymbolTable()
		if !st.Contains("R0") {
			t.Errorf("Contains('R0') should be true for a predefined symbol")
		}
	})

	t.Run("non-existent symbol", func(t *testing.T) {
		st := NewSymbolTable()
		if st.Contains("nonexistent") {
			t.Errorf("Contains('nonexistent') should be false")
		}
	})

	t.Run("user-added symbol", func(t *testing.T) {
		st := NewSymbolTable()
		st.AddEntry("my_var", 123)
		if !st.Contains("my_var") {
			t.Errorf("Contains('my_var') should be true after adding it")
		}
	})
}

func TestGetAddress(t *testing.T) {
	t.Run("predefined symbol", func(t *testing.T) {
		st := NewSymbolTable()
		if addr := st.GetAddress("R15"); addr != 15 {
			t.Errorf("Expected address of 'R15' to be 15, but got %d", addr)
		}
	})

	t.Run("user-added symbol", func(t *testing.T) {
		st := NewSymbolTable()
		st.AddEntry("new_label", 256)
		if addr := st.GetAddress("new_label"); addr != 256 {
			t.Errorf("Expected address of 'new_label' to be 256, but got %d", addr)
		}
	})

	t.Run("non-existent symbol", func(t *testing.T) {
		st := NewSymbolTable()
		if addr := st.GetAddress("not_found"); addr != 0 {
			t.Errorf("Expected address of non-existent symbol 'not_found' to be 0, but got %d", addr)
		}
	})
}

func TestSymbolTableIsolation(t *testing.T) {
	st1 := NewSymbolTable()
	st2 := NewSymbolTable()

	st1.AddEntry("new_symbol", 100)

	if st2.Contains("new_symbol") {
		t.Error("Modifying one symbol table instance affected another one. They should be isolated.")
	}

	st1.AddEntry("R1", 999)
	if st2.GetAddress("R1") != 1 {
		t.Errorf("Overwriting a predefined symbol in one table affected another. Expected R1 address in st2 to be 1, got %d", st2.GetAddress("R1"))
	}
}
