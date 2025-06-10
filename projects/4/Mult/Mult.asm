// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// The algorithm is based on repetitive addition.

//// Replace this comment with your code.

/*
R0とR1に保存された値を使って R0*R1を計算してR2に保存する。
前提として R0>0 , R1 >0 , R0*R1<32768
*/

@R2
M=0

@count
M=0

(LOOP)
    @count
    D=M

    // D = R1 - count
    @R1
    D=M-D

    // if count == R1 then END else continue
    @END
    D;JEQ

    // R2 += R0
    @R0
    D=M
    @R2
    M=D+M

    // count++
    @count
    M=M+1

    @LOOP
    0;JMP

(END)
    @END
    0;JMP