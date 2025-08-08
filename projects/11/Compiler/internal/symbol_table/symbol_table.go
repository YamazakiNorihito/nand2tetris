package symboltable

type VariableKind string

const (
	NONE   VariableKind = "unknown"
	STATIC VariableKind = "static"
	FIELD  VariableKind = "field"
	ARG    VariableKind = "arg"
	VAR    VariableKind = "var"
)

type Record struct {
	Name  string
	Type  string
	Kind  VariableKind
	Index int
}

type classScope struct {
	staticScope map[string]Record
	fieldScope  map[string]Record
}

type subroutineScope struct {
	argScope map[string]Record
	varScope map[string]Record
}

type SymbolTable struct {
	classScope      classScope
	subroutineScope subroutineScope
	indexCounters   map[VariableKind]int
}

func New() *SymbolTable {
	return &SymbolTable{
		classScope: classScope{
			staticScope: make(map[string]Record),
			fieldScope:  make(map[string]Record),
		},
		subroutineScope: subroutineScope{
			argScope: make(map[string]Record),
			varScope: make(map[string]Record),
		},
		indexCounters: map[VariableKind]int{
			STATIC: 0,
			FIELD:  0,
			ARG:    0,
			VAR:    0,
		},
	}
}

func (st *SymbolTable) Reset() {
	st.subroutineScope = subroutineScope{
		argScope: make(map[string]Record),
		varScope: make(map[string]Record),
	}
	st.indexCounters[ARG] = 0
	st.indexCounters[VAR] = 0
}

/*
   The search order should be: subroutineScope -> classScope.
*/

// type is int, boolean, char, class
func (st *SymbolTable) Define(name, typ string, kind VariableKind) {
	if kind == NONE {
		return
	}

	if st.KindOf(name) != NONE {
		panic("SymbolTable: Attempt to redefine variable " + name)
	}

	var index int
	switch kind {
	case ARG:
		index = st.indexCounters[ARG]
		st.subroutineScope.argScope[name] = Record{Name: name, Type: typ, Kind: kind, Index: index}
		st.indexCounters[ARG]++
	case VAR:
		index = st.indexCounters[VAR]
		st.subroutineScope.varScope[name] = Record{Name: name, Type: typ, Kind: kind, Index: index}
		st.indexCounters[VAR]++
	case STATIC:
		index = st.indexCounters[STATIC]
		st.classScope.staticScope[name] = Record{Name: name, Type: typ, Kind: kind, Index: index}
		st.indexCounters[STATIC]++
	case FIELD:
		index = st.indexCounters[FIELD]
		st.classScope.fieldScope[name] = Record{Name: name, Type: typ, Kind: kind, Index: index}
		st.indexCounters[FIELD]++
	default:
		return
	}
}

func (st *SymbolTable) VarCount(kind VariableKind) int {
	if kind == NONE {
		return 0
	}
	return st.indexCounters[kind]
}

func (st *SymbolTable) KindOf(name string) VariableKind {
	if record, exists := st.subroutineScope.argScope[name]; exists {
		return record.Kind
	}
	if record, exists := st.subroutineScope.varScope[name]; exists {
		return record.Kind
	}
	if record, exists := st.classScope.staticScope[name]; exists {
		return record.Kind
	}
	if record, exists := st.classScope.fieldScope[name]; exists {
		return record.Kind
	}
	return NONE
}

func (st *SymbolTable) TypeOf(name string) string {
	if record, exists := st.subroutineScope.argScope[name]; exists {
		return record.Type
	}
	if record, exists := st.subroutineScope.varScope[name]; exists {
		return record.Type
	}
	if record, exists := st.classScope.staticScope[name]; exists {
		return record.Type
	}
	if record, exists := st.classScope.fieldScope[name]; exists {
		return record.Type
	}
	return ""
}

func (st *SymbolTable) IndexOf(name string) int {
	if record, exists := st.subroutineScope.argScope[name]; exists {
		return record.Index
	}
	if record, exists := st.subroutineScope.varScope[name]; exists {
		return record.Index
	}
	if record, exists := st.classScope.staticScope[name]; exists {
		return record.Index
	}
	if record, exists := st.classScope.fieldScope[name]; exists {
		return record.Index
	}
	return -1 // Not found
}
