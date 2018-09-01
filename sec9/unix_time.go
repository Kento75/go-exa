package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	unix := t.Unix()
	fmt.Println("Normal：", t)
	fmt.Println("Unix：", unix)
}
