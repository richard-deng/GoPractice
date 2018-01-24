package dlog

import (
	"os"
	"log"
	"fmt"
	//"time"
)

func DcLog() *log.Logger {
	//创建输出日志文件
	//logFile, err := os.Open("20180123.txt")
	logFile, err := os.OpenFile("20180123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//logFile, err := os.Create("./" + time.Now().Format("20060102") + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	loger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	//SetFlags设置输出选项
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//返回输出前缀
	fmt.Println(loger.Prefix())

	//设置输出前缀
	//loger.SetPrefix("test_")
	return loger
}