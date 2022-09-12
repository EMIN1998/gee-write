package main

import (
	"fmt"
	"gee-rewrite/gee"
	"net/http"
)

func main() {
	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/hello", helloHandler)
	r := gee.New()
	r.GET("/hello", helloHandler)
	//fmt.r.RUN("2500")
	fmt.Errorf("======= %v", r.RUN("127.0.0.1:2500"))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(ctx *gee.Context) {
	for k, v := range ctx.Request.Header {
		fmt.Fprintf(ctx.Writer, "Header[%q] = %q\n", k, v)
	}

	fmt.Print("request income")
}
