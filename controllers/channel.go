package controllers

import (
	"log"
	"net/http"
	"encoding/json"
	"build_web/GoPractice/model"
	"strconv"
)

func ChannelHandler(w http.ResponseWriter, r *http.Request) {
	resp := model.Response{}
	if r.Method != "POST" {
		resp.Respcd = "1000"
		resp.Resperr = "请求方法错误"
		resp.Respmsg = "请求方法错误"
		resp.Data = nil
		str, _ := json.Marshal(resp)
		w.Write(str)
	}
	log.Println("test here end")
	page, _ := strconv.ParseInt(r.PostFormValue("page"), 10, 64)
	maxNum, _:= strconv.ParseInt(r.PostFormValue("maxnum"), 10, 64)
	log.Printf("page: %d, num: %d", page, maxNum)
	db := GetConn()
	defer db.Close()
	allChannel := QueryAllChannelInfo(db, page, maxNum)
	totalNum := QueryChannelAllTotal(db)
	type MyData struct {
		Info []model.Channel	`json:"info"`
		Num  int            `json:"num"`
	}
	data := MyData{
		allChannel,
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

func ChannelNamesHandler(w http.ResponseWriter, r *http.Request) {
	resp := model.Response{}
	if r.Method != "GET" {
		resp.Respcd = "1000"
		resp.Resperr = "请求方法错误"
		resp.Respmsg = "请求方法错误"
		resp.Data = nil
		str, _ := json.Marshal(resp)
		w.Write(str)
	}
	db := GetConn()
	defer db.Close()
	names := QueryChannelNames(db)
	resp.Respcd = "0000"
	resp.Resperr = ""
	resp.Respmsg = ""
	resp.Data = names
	str, _ := json.Marshal(resp)
	log.Println(string(str))
	w.Write(str)
}