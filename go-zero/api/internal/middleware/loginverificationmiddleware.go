package middleware

import (
	"fmt"
	"net/http"
)

type LoginVerificationMiddleware struct {
}

func NewLoginVerificationMiddleware() *LoginVerificationMiddleware {
	return &LoginVerificationMiddleware{}
}

func (m *LoginVerificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("token"))
		if r.Header.Get("token") == "go-zero" {
			next(w, r)
			return
		}
		_, err := w.Write([]byte("您没有权限"))
		if err != nil {
			panic(err)
		}
		return
	}
}
