package main

import "fmt"

var (
	Version  = "unknown"
	Revision = "unknown"
)

func main() {
	fmt.Printf("Hello World: %s %s\n", Version, Revision)
}
