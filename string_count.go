package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// string型の変数を宣言
	var ja string = "Go言語"

	// 文字列を出力
	fmt.Println(ja, "len:", utf8.RuneCountInString(ja)) // 4と出力される
}
