package main

import "fmt"

func main() {
	// Printlnを遅延指定
	defer fmt.Println("defer指定はpanic時でも実行される")

	f1()
}

func f1() {
	panic("パニック発生")
}
