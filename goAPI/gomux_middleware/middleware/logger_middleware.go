package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Printf("URL %v dipanggil pada jam %v\n", req.URL.Path, time.Now())
		fmt.Println(1)
		next.ServeHTTP(resp, req) // akan menjalan fungsi yang memanggilnya, setelah selesai akan melanjutkan di bawah nya
		fmt.Println(2)
		fmt.Printf("URL %v dipanggil pada jam %v\n", req.URL.Path, time.Now())
	})
}
