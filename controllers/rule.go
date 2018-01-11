package controllers

import (
	"log"
	"net/http"
	"encoding/json"
	"build_web/GoPractice/model"
	"strconv"
)

func RuleHandler(w http.ResponseWriter, r *http.Request) {
	resp := model.Response{}
	if r.Method != "POST" {
		resp.Respcd = "1000"
		resp.Resperr = "请求方法错误"
		resp.Respmsg = "请求方法错误"
		resp.Data = nil
		str, _ := json.Marshal(resp)
		w.Write(str)
	}
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	log.Println("test here")
	val := r.PostForm
	log.Println(val)
	log.Println("test here end")
	page, _ := strconv.ParseInt(val["page"][0], 10, 32)
	maxNum, _ := strconv.ParseInt(val["maxnum"][0],10, 32)
	log.Printf("page: %d, num: %d", page, maxNum)
	db := GetConn()
	defer db.Close()
	allRule := QueryRuleInfo(db, page, maxNum)
	log.Println("allRule")
	log.Println(allRule)
	totalNum := QueryRuleAllTotal(db)
	type MyData struct {
		Info []model.Rule	`json:"info"`
		Num  int            `json:"num"`
	}
	data := MyData{
		allRule,
		totalNum,
	}
	resp.Respcd = "0000"
	resp.Resperr = ""
	resp.Respmsg = ""
	resp.Data = data
	str, _ := json.Marshal(resp)
	log.Println(string(str))
	w.Write(str)
}

func RuleNamesHandler(w http.ResponseWriter, r *http.Request)  {
	resp := model.Response{}
	if r.Method != "GET" {
		resp.Resperr = "1000"
		resp.Resperr = "请求方法错误"
		resp.Respmsg = "请求方法错误"
		resp.Data = nil
		str, _ := json.Marshal(resp)
		w.Write(str)
		return
	}
	db := GetConn()
	defer db.Close()
	names := QueryRuleNames(db)
	resp.Respcd = "0000"
	resp.Respmsg = ""
	resp.Resperr = ""
	resp.Data = names
	str, _ := json.Marshal(resp)
	w.Write(str)
	return
}