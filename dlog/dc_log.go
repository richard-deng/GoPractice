package dlog

import (
	"os"
	"log"
	"fmt"
	"time"
)

var (
	Debug *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init() {
	//创建输出日志文件
	//logFile, err := os.Open("20180123.txt")
	//logFile, err := os.OpenFile("20180123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logFile, err := os.OpenFile("./" + time.Now().Format("20060102") + ".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//logFile, err := os.Create("./" + time.Now().Format("20060102") + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性

	//SetFlags设置输出选项
	//loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//返回输出前缀
	//fmt.Println(loger.Prefix())
	//设置输出前缀
	//loger.SetPrefix("DEBUG")

	Debug = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}