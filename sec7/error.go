package main

import "fmt"

// MyError型
type MyError struct {
	message string // エラー詳細
}

// Errorメソッドの実装
func (err MyError) Error() string {
	// エラーメッセージを返す
	return err.message
}

func main() {
	val, err := hex2int("1")
	fmt.Println(val, err)

	val, err = hex2int("abcd")
	fmt.Println(val, err)

	val, err = hex2int("z")
	fmt.Println(val, err)
}

// 16進数文字列をint型に変換
func hex2int(hex string) (val int, err error) {
	// 1文字ずつ取り出す
	for _, r := range hex {
		val *= 16
		switch {
		case '1' <= r && r <= '9':
			val += int(r - '0')
		case 'a' <= r && r <= 'f':
			val += int(r - 'a') + 10
		default:
			// エラー情報として返す
			return 0, MyError{"不正な文字列です。：" + string(r)}
		}
	}
	// 戻り値errには初期値であるnilが返される
	return
}
