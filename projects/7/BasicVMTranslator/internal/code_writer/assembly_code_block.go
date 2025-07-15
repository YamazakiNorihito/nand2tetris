package codewriter

var (
	add = `// add
@SP
M=M-1 // SP--
A=M   // D = *SP
D=M
@SP
M=M-1 // SP--
A=M
M=D+M
@SP
M=M+1 // SP++
`
	sub = `// sub
@SP
M=M-1 // SP--
A=M   // D = *SP
D=M
@SP
M=M-1 // SP--
A=M
M=M-D
@SP
M=M+1 // SP++
`
	neg = `// neg
@SP
M=M-1 // SP--
A=M
M=-M
@SP
M=M+1 // SP++
`
	and = `// and
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M&D
@SP
M=M+1
`
	or = `// or
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=D|M
@SP
M=M+1
`
	not = `// not
@SP
M=M-1
A=M
M=!M
@SP
M=M+1
`
	eq = `// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@EQ_TRUE_%s
D;JEQ
@SP
A=M
M=0
@EQ_END_%s
0;JMP
(EQ_TRUE_%s)
@SP
A=M
M=-1
(EQ_END_%s)
@SP
M=M+1
`
	gt = `// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@GT_TRUE_%s
D;JGT
@SP
A=M
M=0
@GT_END_%s
0;JMP
(GT_TRUE_%s)
@SP
A=M
M=-1
(GT_END_%s)
@SP
M=M+1
`
	lt = `// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@LT_TRUE_%s
D;JLT
@SP
A=M
M=0
@LT_END_%s
0;JMP
(LT_TRUE_%s)
@SP
A=M
M=-1
(LT_END_%s)
@SP
M=M+1
`
	pushConstant = `// push constant %d
@%d
D=A
@SP
A=M
M=D
@SP
M=M+1
`
	pushLocal = `// push local %d
@LCL
D=M
@%d
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushArgument = `// push argument %d
@ARG
D=M
@%d
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushThis = `// push this %d
@THIS
D=M
@%d
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushThat = `// push that %d
@THAT
D=M
@%d
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushTemp = `// push temp %d
@5
D=A
@%d
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushPointer = `// push pointer %s
@%s
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	pushStatic = `// push static %s
@%s.%s // アセンブラによってアドレス16から自動で割り当てられる
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	popLocal = `// pop local %d
@LCL
D=M
@%d
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`
	popArgument = `// pop argument %d
@ARG
D=M
@%d
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`
	popThis = `// pop this %d
@THIS
D=M
@%d
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`
	popThat = `// pop that %d
@THAT
D=M
@%d
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`
	popTemp = `// pop temp %d
@5
D=A
@%d
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`
	popPointer = `// pop pointer %s
@SP
M=M-1
A=M
D=M
@%s
M=D
`
	popStatic = `// pop static %s
@SP
M=M-1
A=M
D=M
@%s.%s
M=D
`
)
