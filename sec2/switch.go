package main

import "fmt"

func main() {
	// ループ
	for i := 0; i < 5; i++ {
		// switch文、暗黙的にbreakが実行される
		switch i {
		case 0:
			fmt.Println("0")
		case 1, 2:
			fmt.Println("1または2")
		default:
			fmt.Println("その他")
		}
	}
}

