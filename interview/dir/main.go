package main

import (
	"fmt"
	"os"
)

func main() {
	ex, _ := os.Executable()
	fmt.Println(ex)
}
