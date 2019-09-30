package handler

import "net/http"

//测试服务器
func Ping (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG!"))
}