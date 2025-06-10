// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

//// Replace this comment with your code.

/*
LOOP: キーボード入力を監視し、画面色（COLOR）を決定する
      - キーが押されていれば黒（-1）、押されていなければ白（0）をCOLORにセット
      - FILLにジャンプして画面を塗りつぶす
 */
(LOOP)
@KBD
D=M

// KBD != 0 means a key is pressed
@BLACK
D;JNE

@COLOR
M=0

@FILL
0;JMP

/*
BLACK: COLORに黒（-1）をセットし、FILLにジャンプする
      - 責務は「黒色のセット」のみ
 */
(BLACK)
@COLOR
M=-1

@FILL
0;JMP


/*
FILL: 画面塗りつぶしの準備をする
      - current_address（ポインタ）をSCREENの先頭にセット
      - countを8192（画面ワード数）にセット
      - FILL_LOOPにジャンプ
 */
(FILL)
// 画面の先頭アドレスを current_address にセット（current_addressはポインタ変数）
// current_address = SCREENの先頭アドレス（ポインタとして使う）
@SCREEN
D=A
@current_address
M=D

// 画面全体のワード数（8192）を count にセット
// count = 8192
@8192
D=A
@count
M=D

/*
FILL_LOOP: 画面全体をCOLORで塗りつぶす
      - current_addressが指すアドレスにCOLORを書き込む（ポインタ間接参照）
      - countが0になるまで繰り返す
      - 終了したらLOOPに戻る
 */
(FILL_LOOP)
    // if count==0 goto LOOP  else continue
    @count
    D=M
    @LOOP
    D;JEQ

    // current_addressが指すアドレスにCOLORを書き込む（ポインタ間接参照
    // つまり、current_addressが指すアドレスにCOLORの値を書き込む）
    // *(SCREEN) = @COLOR
    @COLOR
    D=M
    @current_address
    A=M
    M=D

    // 残りワード数をデクリメント
    @count
    M=M-1

    // current_address（ポインタ）を次のアドレスへ進める
    @current_address
    M=M+1

    @FILL_LOOP
    0;JMP