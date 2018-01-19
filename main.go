package main

import (
	"log"
	"net/http"
	"build_web/GoPractice/controllers"
	"time"
)

func logHandler(next http.Handler) http.Handler {
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
	mux.Handle("/login.html", logHandler(http.HandlerFunc(controllers.Login)))
	mux.Handle("/overview.html", logHandler(http.HandlerFunc(controllers.Overview)))
	mux.Handle("/users.html", logHandler(http.HandlerFunc(controllers.UserView)))
	mux.Handle("/channel.html", logHandler(http.HandlerFunc(controllers.ChannelView)))
	mux.Handle("/rule.html", logHandler(http.HandlerFunc(controllers.RuleView)))

	//接口数据
	mux.Handle("/api/user/all", logHandler(http.HandlerFunc(controllers.UserHandler)))
	mux.Handle("/api/user/mobile", logHandler(http.HandlerFunc(controllers.UserInfoByPhoneNumber)))
	mux.Handle("/api/user/password/change", logHandler(http.HandlerFunc(controllers.UserChangePasswordHandler)))
	mux.Handle("/api/channel/all", logHandler(http.HandlerFunc(controllers.ChannelHandler)))
	mux.Handle("/api/channel/names/all", logHandler(http.HandlerFunc(controllers.ChannelNamesHandler)))
	mux.Handle("/api/rule/all", logHandler(http.HandlerFunc(controllers.RuleHandler)))
	mux.Handle("/api/rule/names/all", logHandler(http.HandlerFunc(controllers.RuleNamesHandler)))
	mux.Handle("/api/rule/create", logHandler(http.HandlerFunc(controllers.RuleCreateHandler)))

	//启动服务
	log.Println("Listening :8080")
	http.ListenAndServe(":8080", mux)
}
