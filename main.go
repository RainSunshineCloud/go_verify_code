package main

import (
	"fmt"
	//"github.com/fvbock/endless"
	config "github.com/RainSunshineCloud/go-iniconfig"
	"runtime"
	"log"
	"net/http"
	"os"
)
//启动服务
var Config *config.Config;
func main () {
	Config := config.New("./config.ini",true)
	if !Config.Load() {
		panic(Config.LastErr())
	}
	a := Config.Get("global.group")
	fmt.Println(a,Config.LastErr())
	// 添加服务器
	if runtime.GOOS == "windows" {
		newWindowsServ()
	} else {
		//newLinuxServ()
	}

}

//func newLinuxServ() {
//	err := endless.ListenAndServe("localhost:4242", GetRouter())
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println("verify_code stop")
//
//	os.Exit(0)
//}

func newWindowsServ() {
	err := http.ListenAndServe("localhost:8000", GetRouter())
	if err != nil {
		log.Println(err)
	}
	log.Println("verify_code stop")

	os.Exit(0)
}

