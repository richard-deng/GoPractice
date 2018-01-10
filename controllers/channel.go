package controllers

import (
	"log"
	"net/http"
	"encoding/json"
	"build_web/my_project/model"
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
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	log.Println("test here")
	val := r.PostForm
	log.Println(val)
	log.Println("test here end")
	page, _ := strconv.ParseInt(val["page"][0], 10, 32)
	maxnum, _ := strconv.ParseInt(val["maxnum"][0],10, 32)
	log.Printf("page: %d, num: %d", page, maxnum)
	db := GetConn()
	defer db.Close()
	all_channel := QueryAllChannelInfo(db, page, maxnum)
	total_num := QueryChannelAllTotal(db)
	type MyData struct {
		Info []model.Channel	`json:"info"`
		Num  int            `json:"num"`
	}
	data := MyData{
		all_channel,
		total_num,
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