package main

import (
	"fmt"
	"os"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: reverse <string>")
		return
	}

	fmt.Println(Reverse((os.Args[1])))
}
