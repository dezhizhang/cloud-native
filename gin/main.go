package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	files, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		panic(err)
	}

	//user := User{
	//	Name:   "Jack",
	//	Gender: "男",
	//	Age:    18,
	//}

	m1 := map[string]interface{}{
		"Name":   "Jack",
		"Gender": "男",
		"Age":    18,
	}

	// 渲染模板
	err = files.Execute(w, m1)

	if err != nil {
		panic(err)
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
