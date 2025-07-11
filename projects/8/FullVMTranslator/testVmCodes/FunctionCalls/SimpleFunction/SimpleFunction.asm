// function SimpleFunction.test.  nVar=2 from SimpleFunction
(SimpleFunction.test)

// nVar=2だけ0で初期化
// push constant 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

// push local 0
@LCL
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// push local 1
@LCL
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// add
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

// not
@SP
M=M-1
A=M
M=!M
@SP
M=M+1

// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// add
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

// push argument 1
@ARG
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// sub
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

// return
// frame = LCL
@LCL
D=M
@R13
M=D

// return adress = *(frame - 5)
@5
A=D-A
D=M
@R14
M=D   // R14 = return adress

// *ARG = pop()
@SP
M=M-1 // SP--
A=M
D=M
@ARG
A=M
M=D

// SP = ARG + 1
@ARG
D=M
@SP
M=D+1

// THAT = *(frame - 1)
@R13
M=M-1
A=M
D=M
@THAT
M=D

// THIS = *(frame - 2)
@R13
M=M-1
A=M
D=M
@THIS
M=D

// ARG = *(frame - 3)
@R13
M=M-1
A=M
D=M
@ARG
M=D

// LCL = *(frame - 4)
@R13
M=M-1
A=M
D=M
@LCL
M=D

// goto return address
@R14
A=M
0;JMP 

