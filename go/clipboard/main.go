package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	if err := clipboard.WriteAll("line1\nline2"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
