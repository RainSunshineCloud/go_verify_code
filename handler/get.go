package handler

import "net/http"

//获取验证码
func Get (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WORLD!"))
}