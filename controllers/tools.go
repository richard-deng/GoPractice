package controllers

import (
	"fmt"
	"strings"
	"strconv"
	"build_web/GoPractice/dlog"
)

func BuildWhere(whereMap map[string]string, intArr []string) string {
	var whereArr []string
	for key := range whereMap {
		if whereMap[key] != "" {
			if len(intArr) > 0 && golangIn(key, intArr) {
				dlog.Warning.Println("int arr")
				v, _ := strconv.ParseInt(whereMap[key], 10, 64)
				whereArr = append(whereArr, fmt.Sprintf("%s=%d", key, v))
			} else {
				whereArr = append(whereArr, fmt.Sprintf("%s=\"%s\"", key, whereMap[key]))
			}
		}
	}
	whereStr := strings.Join(whereArr, " and ")
	dlog.Info.Println("whereStr", whereStr)
	return whereStr
}

func golangIn(name string, arr []string) bool {
	var flag = false
	for _, v := range arr {
		if name == v {
			flag = true
			break
		}
	}
	dlog.Info.Println("name", name, "arr", arr, "flag", flag)
	return flag
}