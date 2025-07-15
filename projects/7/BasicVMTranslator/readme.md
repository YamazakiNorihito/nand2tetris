`SimpleAdd`からもわかる通り、VMからアセンブリへ変換するとコードが冗長になりがち。

たとえば、同じ処理を直接アセンブリで書くと非常に短くなります。

```asm
// 7 と 8 を足してスタックにプッシュ
@7
D=A
@8
D=D+A
@SP
A=M
M=D
@SP
M=M+1
```

このように、VMコマンド一つを変換すると多くのアセンブリ命令が必要になりますが、直接書く場合は数行で済みます。

```bash
go build -o BasicVMTranslator ./cmd

$ find ./testVmCodes -name "*.vm" | sort | sed 's/^/\.\/BasicVMTranslator /'
./BasicVMTranslator ./testVmCodes/MemoryAccess/BasicTest/BasicTest.vm
./BasicVMTranslator ./testVmCodes/MemoryAccess/PointerTest/PointerTest.vm
./BasicVMTranslator ./testVmCodes/MemoryAccess/StaticTest/StaticTest.vm
./BasicVMTranslator ./testVmCodes/StackArithmetic/SimpleAdd/SimpleAdd.vm
./BasicVMTranslator ./testVmCodes/StackArithmetic/StackTest/StackTest.vm
```
