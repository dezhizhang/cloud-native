package main

import "fmt"

//func main() {
//	ok := true
//	fmt.Printf("%T,%t", ok, ok) // bool,true
//
//	var r rune = 65
//	fmt.Printf("%T,%d\n", r, r)         //int32,65
//	fmt.Printf("%T,%5d\n", r, r)        //int32,   65
//	fmt.Printf("%T,%05d\n", r, r)       //int32,00065
//	fmt.Printf("%T,%x\n", r, r)         //int32,41
//	fmt.Printf("%T,%c\n", r, r)         //int32,A
//	fmt.Printf("%T,%U\n", r, r)         //int32,U+0041
//	fmt.Printf("%T,%v\n", r, r)         //int32,65
//	fmt.Printf("%T,%s\n", r, string(r)) //int32,A
//
//}

func main() {
	var s = "面试"
	for _, item := range s {
		fmt.Printf("i:%c\n", item)
	}
}
