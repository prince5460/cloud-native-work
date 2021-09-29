package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 80
	http.HandleFunc("/healthz", HealthZ)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		fmt.Println("ListenAndServe err:", err)
		return
	}
}

func HealthZ(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	reqHeader, err := json.Marshal(r.Header)
	//将request的header写入response header
	w.Header().Set("Header", string(reqHeader))
	//将环境变量写入response header
	w.Header().Set("VERSION", os.Getenv("VERSION"))

	//返回200
	_, err = io.WriteString(w, "200")
	if err != nil {
		return
	}

	for k, v := range w.Header() {
		fmt.Printf("%s:%s\n", k, v)
	}

	//记录访问日志包括客户端IP等信息
	fmt.Printf("%s | %s | %s | %d", r.RemoteAddr, r.Method, r.RequestURI, http.StatusOK)
}
