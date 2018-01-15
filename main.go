package main

import (
	"log"
	"net/http"
	"build_web/GoPractice/controllers"
)

func main() {
	//新建一个处理Handle
	mux := http.NewServeMux()

	//建立一个文件服务器
	fs := http.FileServer(http.Dir("./public"))

	//静态文件
	mux.Handle("/", fs)

	//页面
	mux.Handle("/login.html", http.HandlerFunc(controllers.Login))
	mux.Handle("/overview.html", http.HandlerFunc(controllers.Overview))
	mux.Handle("/users.html", http.HandlerFunc(controllers.UserView))
	mux.Handle("/channel.html", http.HandlerFunc(controllers.ChannelView))
	mux.Handle("/rule.html", http.HandlerFunc(controllers.RuleView))

	//接口数据
	mux.Handle("/api/user/all", http.HandlerFunc(controllers.UserHandler))
	mux.Handle("/api/user/mobile", http.HandlerFunc(controllers.UserInfoByPhoneNumber))
	mux.Handle("/api/user/password/change", http.HandlerFunc(controllers.UserChangePasswordHandler))
	mux.Handle("/api/channel/all", http.HandlerFunc(controllers.ChannelHandler))
	mux.Handle("/api/channel/names/all", http.HandlerFunc(controllers.ChannelNamesHandler))
	mux.Handle("/api/rule/all", http.HandlerFunc(controllers.RuleHandler))
	mux.Handle("/api/rule/names/all", http.HandlerFunc(controllers.RuleNamesHandler))
	mux.Handle("/api/rule/create", http.HandlerFunc(controllers.RuleCreateHandler))

	//启动服务
	log.Println("Listening :8080")
	http.ListenAndServe(":8080", mux)
}
