package main

import "fmt"

func main() {
	// f1の戻り値をf2のパラメータに使用する
	f2(f1())
}

// 多値を返す関数
func f1() (int, string, float32) {
	return 0, "xyz", 3.14
}

func f2(a int, b string, c interface{}) {
	fmt.Println(a, b, c)
}
