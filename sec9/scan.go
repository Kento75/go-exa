package main

import "fmt"

func main() {
	// 文字列入力
	var fname, name string

	fmt.Println("あなたのお名前は（姓 名）？")
	fmt.Scanln(&fname, &name)

	// 数値入力
	fmt.Println("あなたの年齢は？")

	var age int
	fmt.Scanf("%d", &age)

	// 入力結果の出力
	fmt.Printf("名前：%s %s 年齢：%d\n", fname, name, age)
}
