package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		fmt.Println("path:", r.URL.Path)

		a, _ := strconv.Atoi(r.Form.Get("a"))
		b, _ := strconv.Atoi(r.Form.Get("b"))

		w.Header().Set("content-type", "application/json")
		_, _ = w.Write([]byte(fmt.Sprintf("%d", a+b)))
	})
	_ = http.ListenAndServe(":8080", nil)
}
