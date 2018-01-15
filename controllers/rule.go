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
	page, _ := strconv.ParseInt(r.PostFormValue("page"), 10, 64)
	maxNum, _:= strconv.ParseInt(r.PostFormValue("maxnum"), 10, 64)
	var name = r.PostFormValue("name")
	log.Printf("page: %d, num: %d, name: %s", page, maxNum, name)
	db := GetConn()
	defer db.Close()
	allRule := QueryRuleInfo(db, page, maxNum, name)
	log.Println("allRule")
	log.Println(allRule)
	totalNum := QueryRuleAllTotal(db, name)
	type MyData struct {
		Info []model.Rule	`json:"info"`
		Num  int64            `json:"num"`
	}
	data := MyData{}
	if totalNum == 0 {
		data.Info = []model.Rule{}
	} else {
		data.Info = allRule
	}
	data.Num = totalNum
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