package main

import (
	"os"
	"fmt"
	"encoding/csv"
)

func main() {
	// ファイルのオープン
	file, err := os.OpenFile("test.csv", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ファイルのクローズ(遅延実行)
	defer func() {
		file.Close()
	}()

	// csv.Reader
	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err != nil {
			break
		}

		fmt.Println(record)
	}
}
