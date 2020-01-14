package advance

import (
	"fmt"
	"net/http"
)

func Logger2(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Header-x", "XXXXXXXX")
		fmt.Println(3)
		next.ServeHTTP(resp, req) // akan menjalan fungsi yang memanggilnya, setelah selesai akan melanjutkan di bawah nya
		fmt.Println(4)
		resp.Write([]byte("Done Testing"))
	})
}
