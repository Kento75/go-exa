package main

import (
	"os"
	"fmt"
)

func main() {
	write()
	read()
}


// 書き込み
func write() {
	// ファイルのオープン
	file, err := os.OpenFile("test.txt", os.O_CREATE | os.O_WRONLY, 0666)

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file.WriteString("test\n")

	fmt.Fprintf(file, "test\n")
}

// 読み込み
func read() {
	// ファイルのオープン
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	var str string
	fmt.Fscanf(file, "%s", &str)

	fmt.Println(str)

	buf := make([]byte, 16)
	size, _ := file.Read(buf)

	fmt.Println(size, buf)
}