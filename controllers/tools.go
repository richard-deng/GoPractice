package controllers

import "fmt"

func InitUserWhere(phoneNum, loginName, nickName string) string {
	whereStr := ""
	var whereArr []string
	var where = map[string]string{}
	where["phone_num"] = phoneNum
	where["login_name"] = loginName
	where["nick_name"] = nickName
	fmt.Println(where)
	for key := range where {
		if where[key] != "" {
			whereArr = append(whereArr, fmt.Sprintf("%s=\"%s\"", key, where[key]))
		}
	}
	fmt.Println(whereArr)
	for _, v := range whereArr {
		whereStr += v
		whereStr += " "
	}
	return whereStr
}
