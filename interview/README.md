# go interview
### init init() 函数是什么时候执行的
1. ##### 同一个go文件的多个init方法，按照定义的顺序执行
```go
package main

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Hello World")
}
```
2. ##### 同一个go文件的多个init方法，按照定义的顺序执行

```go
// Package pkg a.go
var A = "initA"

func init() {
	fmt.Println(A)
}

// Package pkg a.go
var B = "initB"

fmfunc main() {
    fmt.Println("main", pkg.B)
    fmt.Println("main", pkg.A)
}


func main() {
    fmt.Println("main", pkg.B)
    fmt.Println("main", pkg.A)
}

```
### 如何获取项目的根目录
```go
func main() {
    var err error
    dir, _ := os.Getwd()
    path, err := filepath.Abs(dir)
    fmt.Println(path, err)
}

```
### 输出时 %v %+v %#v 有什么区别
```go
func main() {
    ok := true
    fmt.Printf("%T,%t", ok, ok) // bool,true
    
    var r rune = 65
    fmt.Printf("%T,%d\n", r, r) //int32,65
    fmt.Printf("%T,%5d\n", r, r) //int32,   65
    fmt.Printf("%T,%05d\n", r, r) //int32,00065
    fmt.Printf("%T,%x\n", r, r) //int32,41
    fmt.Printf("%T,%c\n", r, r) //int32,A
    fmt.Printf("%T,%U\n", r, r) //int32,U+0041
    fmt.Printf("%T,%v\n", r, r) //int32,65
    fmt.Printf("%T,%s\n", r, string(r)) //int32,A

    var s = "面试"
    for _, item := range s {
    fmt.Printf("i:%c\n", item)
}
```