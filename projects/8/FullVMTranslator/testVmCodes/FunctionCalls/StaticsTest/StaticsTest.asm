// Bootstrap code
@256
D=A
@SP
M=D

// call Sys.init with nArgs=0 from Bootstrap
// push return address Bootstrap$ret.0
@Bootstrap$ret.0
D=A
@SP
A=M
M=D
@SP
M=M+1 // SP++

// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// ARG = SP - nArgs - 5
@SP
D=M
@5
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Sys.init from Bootstrap
@Sys.init
0;JMP

// return address label
(Bootstrap$ret.0)

// function Class1.set.  nVar=0 from Class1
(Class1.set)

// nVar=0だけ0で初期化


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

// pop static 0
@SP
M=M-1
A=M
D=M
@Class1.0
M=D

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

// pop static 1
@SP
M=M-1
A=M
D=M
@Class1.1
M=D

// push constant 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

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

// function Class1.get.  nVar=0 from Class1
(Class1.get)

// nVar=0だけ0で初期化


// push static 0
@Class1.0 // アセンブラによってアドレス16から自動で割り当てられる
D=M
@SP
A=M
M=D
@SP
M=M+1

// push static 1
@Class1.1 // アセンブラによってアドレス16から自動で割り当てられる
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

// Bootstrap code
@256
D=A
@SP
M=D

// call Sys.init with nArgs=0 from Sys
// push return address Sys$ret.1
@Sys$ret.1
D=A
@SP
A=M
M=D
@SP
M=M+1 // SP++

// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// ARG = SP - nArgs - 5
@SP
D=M
@5
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Sys.init from Sys
@Sys.init
0;JMP

// return address label
(Sys$ret.1)

// function Sys.init.  nVar=0 from Sys
(Sys.init)

// nVar=0だけ0で初期化


// push constant 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 8
@8
D=A
@SP
A=M
M=D
@SP
M=M+1

// call Class1.set with nArgs=2 from Sys
// push return address Sys$ret.2
@Sys$ret.2
D=A
@SP
A=M
M=D
@SP
M=M+1 // SP++

// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// ARG = SP - nArgs - 5
@SP
D=M
@7
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Class1.set from Sys
@Class1.set
0;JMP

// return address label
(Sys$ret.2)

// pop temp 0
@5
D=A
@0
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

// push constant 23
@23
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 15
@15
D=A
@SP
A=M
M=D
@SP
M=M+1

// call Class2.set with nArgs=2 from Sys
// push return address Sys$ret.3
@Sys$ret.3
D=A
@SP
A=M
M=D
@SP
M=M+1 // SP++

// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// ARG = SP - nArgs - 5
@SP
D=M
@7
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Class2.set from Sys
@Class2.set
0;JMP

// return address label
(Sys$ret.3)

// pop temp 0
@5
D=A
@0
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

// call Class1.get with nArgs=0 from Sys
// push return address Sys$ret.4
@Sys$ret.4
D=A
@SP
A=M
M=D
@SP
M=M+1 // SP++

// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// ARG = SP - nArgs - 5
@SP
D=M
@5
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Class1.get from Sys
@Class1.get
0;JMP

// return address label
(Sys$ret.4)

// call Class2.get with nArgs=0 from Sys
// push return address Sys$ret.5
@Sys$ret.5
D=A
@SP
A=M
M=D
@SP
M=M+1 // SP++

// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// ARG = SP - nArgs - 5
@SP
D=M
@5
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Class2.get from Sys
@Class2.get
0;JMP

// return address label
(Sys$ret.5)

(Sys$END)
// goto END from Sys
@Sys$END
0;JMP

// function Class2.set.  nVar=0 from Class2
(Class2.set)

// nVar=0だけ0で初期化


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

// pop static 0
@SP
M=M-1
A=M
D=M
@Class2.0
M=D

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

// pop static 1
@SP
M=M-1
A=M
D=M
@Class2.1
M=D

// push constant 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

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

// function Class2.get.  nVar=0 from Class2
(Class2.get)

// nVar=0だけ0で初期化


// push static 0
@Class2.0 // アセンブラによってアドレス16から自動で割り当てられる
D=M
@SP
A=M
M=D
@SP
M=M+1

// push static 1
@Class2.1 // アセンブラによってアドレス16から自動で割り当てられる
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

