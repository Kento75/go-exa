package main

import "fmt"

func main() {
	// 数値の足し算
	fmt.Println("1 + 2 =", 1 + 2)

	// 文字列の足し算
	fmt.Println("\"abc\" + \"XYZ\" =", "abc" + "XYZ")

	// 引き算
	fmt.Println("3 - 2 =", 3 - 2)

	// 掛け算
	fmt.Println("2 * 3 =", 2 * 3)

	// 割り算（丸められる）
	fmt.Println("5 / 2 =", 5/2)

	// 余り
	fmt.Println("5 % 2 =", 5 % 2)
}
