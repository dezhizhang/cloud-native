package main

import (
	"fmt"
	"os"
)

func main() {
	//path, err := filepath.Abs("./")
	//fmt.Println(path, "path")
	//fmt.Println(err)

	//fmt.Println(path, err, "-----")

	path, err := os.Getwd()
	fmt.Println(path, err)
}
