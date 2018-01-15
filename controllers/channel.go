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
	isPrepayment := r.PostFormValue("is_prepayment")
	isValid := r.PostFormValue("is_valid")
	channelName := r.PostFormValue("channel_name")
	phoneNum := r.PostFormValue("phone_num")
	log.Printf("page: %d, num: %d", page, maxNum)
	db := GetConn()
	defer db.Close()
	allChannel := QueryAllChannelInfo(db, page, maxNum, isPrepayment, isValid, channelName, phoneNum)
	log.Println("allChannel", allChannel)
	totalNum := QueryChannelAllTotal(db, isPrepayment, isValid, channelName, phoneNum)
	log.Println("totalNum", totalNum)
	type MyData struct {
		Info []model.Channel	`json:"info"`
		Num  int64            `json:"num"`
	}
	data := MyData{}
	if totalNum == 0 {
		data.Info = []model.Channel{}
		data.Num = totalNum
	} else {
		data.Info = allChannel
		data.Num = totalNum
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