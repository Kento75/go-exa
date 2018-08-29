package main

import (
	"os"
	"fmt"
	"encoding/csv"
)

func main() {
	// ファイルのオープン
	file, err := os.OpenFile("test.csv", os.O_CREATE | os.O_WRONLY, 0666)

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ファイルのクローズ
	defer func() {
		file.Close()
	}()

	// csv.writerを作成
	w := csv.NewWriter(file)

	w.Write([]string{"No", "商品", "価格", "備考"})
	w.Write([]string{"1", "キャベツ", "150", "とれたて新鮮"})
	w.Write([]string{"2", "にんじん", "100", ""})
	w.Write([]string{"3", "サンマ", "99", "今日の\"特価品\"です"})

	w.Flush()
}