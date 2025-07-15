package config

type CommandType int

const (
	C_UNKNOWN    CommandType = iota - 1 // -1, 初期値/エラー状態
	C_ARITHMETIC                        // 0
	C_PUSH                              // 1
	C_POP                               // 2
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

var commandInfo = map[string]struct {
	Type      CommandType
	ArgLength int
}{
	"push":     {C_PUSH, 2},
	"pop":      {C_POP, 2},
	"label":    {C_LABEL, 1},
	"goto":     {C_GOTO, 1},
	"if-goto":  {C_IF, 1},
	"function": {C_FUNCTION, 2},
	"call":     {C_CALL, 2},
	"return":   {C_RETURN, 0},
	"add":      {C_ARITHMETIC, 0},
	"sub":      {C_ARITHMETIC, 0},
	"neg":      {C_ARITHMETIC, 0},
	"and":      {C_ARITHMETIC, 0},
	"or":       {C_ARITHMETIC, 0},
	"not":      {C_ARITHMETIC, 0},
	"eq":       {C_ARITHMETIC, 0},
	"gt":       {C_ARITHMETIC, 0},
	"lt":       {C_ARITHMETIC, 0},
}

func GetCommandInfo(command string) (CommandType, int) {
	if info, ok := commandInfo[command]; ok {
		return info.Type, info.ArgLength
	}
	return C_UNKNOWN, 0
}
