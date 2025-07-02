package config

var arithmeticCommands = map[string]struct{}{
	"add": {}, "sub": {}, "neg": {},
	"and": {}, "or": {}, "not": {},
}

func IsArithmeticCommand(cmd string) bool {
	_, ok := arithmeticCommands[cmd]
	return ok
}

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
