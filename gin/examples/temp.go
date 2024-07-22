package examples

import (
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	files, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		panic(err)
	}
	// 渲染模板
	err = files.Execute(w, "晓智科技有限公司")

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
