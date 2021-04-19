package main

import (
	"fmt"
	"os"
	"strings"
)

func Caps(s string) string {
	return strings.ToUpper(s)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: capitalize <string>")
		return
	}

	fmt.Println(Caps((os.Args[1])))
}
