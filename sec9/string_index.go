package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("The Go Programming Language", "Go"))

	// 存在しない場合はindex値ではなく-1が返ってくる
	fmt.Println(strings.Index("The Go Programming Language", "xxx"))
}
