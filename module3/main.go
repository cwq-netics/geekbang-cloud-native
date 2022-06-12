package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	VERSION := os.Getenv("VERSION")
	w.Header().Add(`VERSION`, VERSION)
	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ip, _ := GetIP(r)
	fmt.Printf(`recv: ip %v, status code %v`, ip, http.StatusOK)
	// 当访问 localhost/healthz 时，应返回 200
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 当访问 localhost/healthz 时，应返回 200
	fmt.Fprintf(w, "Healthz! %s", time.Now())
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8080", nil)
}

func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
