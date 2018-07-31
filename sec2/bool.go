package sec2

import (
	"fmt"
)

func main() {
	// bool型の変数を宣言
	var b bool

	// bool型の変数にリテラル定数TrueとFalseを代入
	b = true
	b = false

	// bool型の変数に論理演算した結果を代入
	b = true || false

	fmt.Println(b)
}