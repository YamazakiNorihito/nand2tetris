```plaintext
入力:
  n    = argument[0]  // 計算したいフィボナッチ数列の要素数
  addr = argument[1]  // 数列の格納先メモリアドレス

処理:
1. THATポインタをaddrに設定する（数列の格納先のベースアドレスを決定）
   THAT = argument[1]

2. フィボナッチ数列の最初の2要素を初期化
   *that[0] = 0
   *that[1] = 1

3. 残り(n - 2)個の要素数をargument[0]に格納し直す
   n = n - 2

4. ループ開始:
   LOOP:
     if n > 0:
       // 次のフィボナッチ数を計算して保存
       that[2] = that[0] + that[1]

       // THATポインタを1つ進めて、次のメモリアドレスを指すようにする
       THAT = THAT + 1

       // n = n - 1（あと何個要素を生成すべきか更新）
       n = n - 1

       goto LOOP
     else:
       goto END

5. END:
   // 終了
```