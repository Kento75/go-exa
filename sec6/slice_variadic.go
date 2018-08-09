package main

import "fmt"

func main() {
	// スライスを作成
	s := []string{"a", "b", "c"}

	// スライスをそのままに渡すには...をつける
	test(s...)

	// こうしたときに渡される値は一緒
	test("a", "b", "c")
}

// 可変長パラメータを持つ関数
func test(s ...string) {
	// スライスの長さと値を出力
	fmt.Println(len(s), s)
}
