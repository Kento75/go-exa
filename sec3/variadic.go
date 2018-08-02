package main

import "fmt"

func main() {
	// 可変長パラメータを持つ関数の呼び出し
	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念日")
	holiday(3, "春分の日")
}

func holiday(month int, days ...string) {
	fmt.Printf("%d月のしゅくじつは%d日あります。\n", month, len(days))

	for _, day := range days {
		fmt.Println(day)
	}
	fmt.Println()
}
