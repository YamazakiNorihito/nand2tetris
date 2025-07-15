// push constant 111
@111
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 333
@333
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 888
@888
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop static 8
@SP
M=M-1
A=M
D=M
@StaticTest.8
M=D
// pop static 3
@SP
M=M-1
A=M
D=M
@StaticTest.3
M=D
// pop static 1
@SP
M=M-1
A=M
D=M
@StaticTest.1
M=D
// push static 3
@StaticTest.3 // アセンブラによってアドレス16から自動で割り当てられる
D=M
@SP
A=M
M=D
@SP
M=M+1
// push static 1
@StaticTest.1 // アセンブラによってアドレス16から自動で割り当てられる
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
// push static 8
@StaticTest.8 // アセンブラによってアドレス16から自動で割り当てられる
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
