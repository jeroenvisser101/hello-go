package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("hello, %v\n", strings.TrimSpace(os.Args[1]))
}
