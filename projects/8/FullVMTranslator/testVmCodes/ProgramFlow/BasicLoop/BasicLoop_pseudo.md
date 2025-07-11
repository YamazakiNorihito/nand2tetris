```plaintext
入力:
  n = argument[0]  // 足し合わせる最大の値（1からnまで）

処理:

1. sum = 0 として初期化
   local[0] = 0

2. ループ開始:
   LOOP:
     sum = sum + n        // 現在のnをsumに加える
     n = n - 1            // nを1減らす
     if n > 0:
       goto LOOP          // まだnが残っていればループ継続

3. 最終結果 sum をスタックにpush
   push sum
```