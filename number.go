package main

import "fmt"

func main() {
	// int型の変数を宣言、環境依存で32bit版、64bit版に別れる（互換性なし）
	var i int = 12345

	// 64bit版に変換
	var i64 int64 = int64(i)

	// 出力
	fmt.Println(i64)
}
