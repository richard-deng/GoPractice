package main

import (
	"log"
	"net/http"
	"build_web/GoPractice/controllers"
	"time"
)

func logginHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Complete %s in %v", r.URL.Path, time.Since(start))
	})
}


func main() {
	//新建一个处理Handle
	mux := http.NewServeMux()

	//建立一个文件服务器
	fs := http.FileServer(http.Dir("./public"))

	//静态文件
	mux.Handle("/", fs)

	//页面
	mux.Handle("/login.html", logginHandler(http.HandlerFunc(controllers.Login)))
	mux.Handle("/overview.html", logginHandler(http.HandlerFunc(controllers.Overview)))
	mux.Handle("/users.html", logginHandler(http.HandlerFunc(controllers.UserView)))
	mux.Handle("/channel.html", logginHandler(http.HandlerFunc(controllers.ChannelView)))
	mux.Handle("/rule.html", logginHandler(http.HandlerFunc(controllers.RuleView)))

	//接口数据
	mux.Handle("/api/user/all", logginHandler(http.HandlerFunc(controllers.UserHandler)))
	mux.Handle("/api/user/mobile", logginHandler(http.HandlerFunc(controllers.UserInfoByPhoneNumber)))
	mux.Handle("/api/user/password/change", logginHandler(http.HandlerFunc(controllers.UserChangePasswordHandler)))
	mux.Handle("/api/channel/all", logginHandler(http.HandlerFunc(controllers.ChannelHandler)))
	mux.Handle("/api/channel/names/all", logginHandler(http.HandlerFunc(controllers.ChannelNamesHandler)))
	mux.Handle("/api/rule/all", logginHandler(http.HandlerFunc(controllers.RuleHandler)))
	mux.Handle("/api/rule/names/all", logginHandler(http.HandlerFunc(controllers.RuleNamesHandler)))
	mux.Handle("/api/rule/create", logginHandler(http.HandlerFunc(controllers.RuleCreateHandler)))

	//启动服务
	log.Println("Listening :8080")
	http.ListenAndServe(":8080", mux)
}
