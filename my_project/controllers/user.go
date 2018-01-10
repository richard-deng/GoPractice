package controllers

import (
	"log"
	"net/http"
	"encoding/json"
	"build_web/my_project/model"
	"strconv"

	"github.com/satori/go.uuid"
	"fmt"
)

/*
登录处理
Flow:
1.校验方法
2.获取数据并处理
3.校验数据合法
4.数据库查询是否存在该信息
	1.有的话判断密码
        1.密码正确设置回话(session)
        2.密码不对返回密码错误信息
    2.不存在则直接返回不存在的错误信息
 */
func Login(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method, r.URL.Path)
	if r.Method == "GET" {
		renderTemplate(w, "login", "login_base", nil)
	} else if r.Method == "POST" {
		resp := model.Response{}
		log.Println("do index post request")
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		val := r.PostForm
		log.Println(val)
		mobile := val["mobile"][0]
		db := GetConn()
		defer db.Close()
		user := QueryByPhoneNumber(db, mobile)
		if user.Id == 0 {
			resp.Resperr = "该用户不存在"
			resp.Respmsg = "该用户不存在"
			resp.Respcd = "2000"
			resp.Data = nil
			log.Println(resp)
			str, _ := json.Marshal(resp)
			w.Write(str)
			return
		}
		log.Println("db_password", user.Password)
		password := val["password"][0]
		log.Printf("my mobile is: %s, password is: %s", string(mobile), string(password))
		flagCheck := CheckPassword(user.Password, password)
		log.Printf("password check result=%s", flagCheck)
		type MyData struct {
			Userid int64 `json:"userid"`
		}
		my := MyData{}
		my.Userid = user.Id
		resp.Respcd = "0000"
		resp.Resperr = ""
		resp.Respmsg = ""
		resp.Data = my
		log.Println("resp:", resp)
		responseStr, _ := json.Marshal(resp)

		u4, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
		}
		userStr := strconv.FormatInt(user.Id, 10)
        SetSession(u4.String(), userStr)
		cookie := http.Cookie{
			Name: "sessionid",
			Value: u4.String(),
			Path: "/",
		}
		w.Header().Set("Set-Cookie", cookie.String())
		w.Write(responseStr)
	} else {
		log.Fatalln("not supported method")
		return
	}
}


func UserHandler(w http.ResponseWriter, r *http.Request)  {
	resp := model.Response{}
	if r.Method != "POST" {
		resp.Respcd = "1000"	//方法不对
		resp.Resperr = "请求方法错误"
		resp.Respmsg = "请求方法错误"
		resp.Data = nil
		resp_str, _ := json.Marshal(resp)
		w.Write(resp_str)
	}
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	val := r.PostForm
	log.Println(val)
	page, _ := strconv.ParseInt(val["page"][0], 10, 32)
	maxnum, _ := strconv.ParseInt(val["maxnum"][0],10, 32)
	log.Printf("page: %d, num: %d", page, maxnum)
	db := GetConn()
	defer db.Close()
	all_user := QueryAllUsersInfo(db, page, maxnum)

	type MyData struct {
		Info []model.User	`json:"info"`
		Num  int            `json:"num"`
	}
	total_num := QueryUsersAllTotal(db)
	my_data := MyData{
		all_user,
		total_num,
	}
	resp.Respcd = "0000"
	resp.Resperr = ""
	resp.Respmsg = ""
	resp.Data = my_data
	resp_str, _ := json.Marshal(resp)
	log.Println(string(resp_str))
	w.Write(resp_str)
}

func UserInfoByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	resp := model.Response{}
	if r.Method != "POST" {
		resp.Respcd = "1000"	//方法不对
		resp.Resperr = "请求方法错误"
		resp.Respmsg = "请求方法错误"
		resp.Data = nil
		resp_str, _ := json.Marshal(resp)
		w.Write(resp_str)
	}
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	val := r.PostForm
	log.Println(val)
	mobile := val["mobile"][0]
	log.Printf("mobile %s", mobile)
	db := GetConn()
	defer db.Close()
	user := QueryByPhoneNumber(db, mobile)
	log.Println(user)
	if user.Valid() {
		resp.Data = user
		resp.Respcd = "0000"
		resp.Respmsg = ""
		resp.Resperr = ""
	} else {
		resp.Data = nil
		resp.Respcd = "1002"
		resp.Respmsg = "用户不存在"
		resp.Resperr = "用户不存在"
	}
	resp_str, _ := json.Marshal(resp)
	log.Println(string(resp_str))
	w.Write(resp_str)
}
