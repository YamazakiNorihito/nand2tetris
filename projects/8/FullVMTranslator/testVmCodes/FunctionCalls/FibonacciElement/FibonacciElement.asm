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

// function Main.fibonacci.  nVar=0 from Main
(Main.fibonacci)

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

// push constant 2
@2
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@LT_TRUE_0
D;JLT
@SP
A=M
M=0
@LT_END_0
0;JMP
(LT_TRUE_0)
@SP
A=M
M=-1
(LT_END_0)
@SP
M=M+1

// if-goto N_LT_2 from Main
@SP
M=M-1 // SP--
A=M   // A = *SP
D=M   // D = *SP
@Main$N_LT_2 // Main$N_LT_2 is a label
D;JNE // if D != 0, jump to label Main$N_LT_2

// goto N_GE_2 from Main
@Main$N_GE_2
0;JMP

(Main$N_LT_2)
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

(Main$N_GE_2)
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

// push constant 2
@2
D=A
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

// call Main.fibonacci with nArgs=1 from Main
// push return address Main$ret.1
@Main$ret.1
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
@6
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Main.fibonacci from Main
@Main.fibonacci
0;JMP

// return address label
(Main$ret.1)

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

// push constant 1
@1
D=A
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

// call Main.fibonacci with nArgs=1 from Main
// push return address Main$ret.2
@Main$ret.2
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
@6
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Main.fibonacci from Main
@Main.fibonacci
0;JMP

// return address label
(Main$ret.2)

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
(Sys$ret.3)

// function Sys.init.  nVar=0 from Sys
(Sys.init)

// nVar=0だけ0で初期化


// push constant 4
@4
D=A
@SP
A=M
M=D
@SP
M=M+1

// call Main.fibonacci with nArgs=1 from Sys
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
@6
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto Main.fibonacci from Sys
@Main.fibonacci
0;JMP

// return address label
(Sys$ret.4)

(Sys$END)
// goto END from Sys
@Sys$END
0;JMP

