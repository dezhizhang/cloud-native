package pkg

import "fmt"

var A = "initA"

func init() {
	fmt.Println(A)
}
