package controllers

import (
	"log"
	"build_web/GoPractice/model"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func GetConn() *sql.DB{
	db, err := sql.Open("mysql", "root:@/uyu")
	if err != nil {
		panic(err)
	}
	return db
}

func QueryUserById(userId int64) model.User {
	db := GetConn()
	defer db.Close()
	var user = model.User{}
	user.Id = 0
	var loginName, nickName, phoneNum, password, email, username, cTime sql.NullString
	var id sql.NullInt64
	var userType, state sql.NullInt64
	rows, err := db.Query("select id, login_name, nick_name, phone_num, password, user_type, email, state, username, ctime from auth_user where id=? limit 1", userId)
	if err != nil {
		panic(err)
	}

	for rows.Next(){
		err := rows.Scan(&id, &loginName, &nickName, &phoneNum, &password, &userType, &email, &state, &username, &cTime)
		if err != nil{
			panic(err)
			return user
		}
	}
	if id.Valid {
		user.Id = id.Int64
	} else {
		user.Id = 0
	}
	if loginName.Valid {
		user.Login_name = loginName.String
	} else {
		user.Login_name = ""
	}
	user.Login_name = loginName.String
	if nickName.Valid {
		user.Nick_name = nickName.String
	} else {
		user.Nick_name = ""
	}
	if phoneNum.Valid {
		user.Phone_num = phoneNum.String
	} else {
		user.Phone_num = ""
	}
	if password.Valid {
		user.Password = password.String
	} else {
		user.Password = ""
	}
	if userType.Valid {
		user.User_type = userType.Int64
	} else {
		user.User_type = 0
	}
	if email.Valid {
		user.Email = email.String
	} else {
		user.Email = ""
	}
	if state.Valid {
		user.State = state.Int64
	} else {
		user.State = 0
	}
	if username.Valid {
		user.Username = username.String
	} else {
		user.Username = ""
	}
	if cTime.Valid {
		user.Ctime = cTime.String
	} else {
		user.Ctime = ""
	}
	return user
}

func QueryByPhoneNumber(db *sql.DB, mobile string) model.User {
	rows, err := db.Query("select id, login_name, nick_name, phone_num, password, user_type, email, state, username, ctime from auth_user where phone_num=? limit 1", mobile)
	if err != nil {
		panic(err)
	}
	log.Println(rows)

	var loginName, nickName, phoneNum, password, email, username, cTime sql.NullString
	var id sql.NullInt64
	var userType, state sql.NullInt64
	for rows.Next() {
		err := rows.Scan(&id, &loginName, &nickName, &phoneNum, &password, &userType, &email, &state, &username, &cTime)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(loginName, nickName, phoneNum)
	}


	log.Println("create_time", cTime)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	var user = model.User{}
	if id.Valid {
		user.Id = id.Int64
	} else {
		user.Id = 0
	}

	if loginName.Valid {
		user.Login_name = loginName.String
	} else {
		user.Login_name = ""
	}

	if nickName.Valid {
		user.Nick_name = nickName.String
	} else {
		user.Nick_name = ""
	}

	if phoneNum.Valid {
		user.Phone_num = phoneNum.String
	} else {
		user.Phone_num = ""
	}

	if password.Valid {
		user.Password = password.String
	} else {
		user.Password = ""
	}

	if userType.Valid {
		user.User_type = userType.Int64
		user.User_role = userType.Int64
	} else {
		user.User_type = 0
		user.User_role = 0
	}

	if email.Valid {
		user.Email = email.String
	} else {
		user.Email = ""
	}

	if state.Valid {
		user.State = state.Int64
	} else {
		user.State = 0
	}

	if username.Valid {
		user.Username = username.String
	} else {
		user.Username = ""
	}

	if cTime.Valid {
		user.Ctime = cTime.String
	} else {
		user.Ctime = ""
	}
	return user
}

func genOffset(page, maxNum int64) (offset, limit int64 ) {
	limit = maxNum
	offset = (page - 1) * maxNum
	return offset, limit
}

func QueryUsersAllTotal(db *sql.DB, phoneNum, loginName, nickName string) int64 {
	var total sql.NullInt64
	if phoneNum == "" && loginName == "" && nickName == "" {
		db.QueryRow("select count(*) as total from auth_user").Scan(&total)
	} else {
		var intArr []string
		var whereMap = map[string]string{}
		whereMap["phone_num"] = phoneNum
		whereMap["login_name"] = loginName
		whereMap["nick_name"] = nickName
		where := BuildWhere(whereMap, intArr)
		if where != "" {
			var querySql = "select count(*) as total from auth_user where %s"
			querySql = fmt.Sprintf(querySql, where)
			//db.QueryRow("select count(*) as total from auth_user where ?", where).Scan(&total)
			db.QueryRow(querySql).Scan(&total)
		}
	}
	return total.Int64
}

func QueryChannelAllTotal(db *sql.DB, isPrepayment, isValid, channelName, phoneNum string) int64 {
	var total sql.NullInt64
	if isPrepayment == "" && isValid == "" && channelName == "" && phoneNum == "" {
		db.QueryRow("select count(*) as total from channel").Scan(&total)
	} else {
		var intArr []string
		intArr = append(intArr, "is_prepayment")
		intArr = append(intArr, "is_valid")
		var whereMap = map[string]string{}
		whereMap["is_prepayment"] = isPrepayment
		whereMap["is_valid"] = isValid
		whereMap["channel_name"] = channelName
		whereMap["phone_num"] = phoneNum
		where := BuildWhere(whereMap, intArr)
		if where != "" {
			log.Println("query channel all ")
			var querySql = "select count(*) as total from channel where %s"
			querySql = fmt.Sprintf(querySql, where)
			log.Println("sql where ", querySql)
			//db.QueryRow("select count(*) as total from channel where ?", where).Scan(&total)
			db.QueryRow(querySql).Scan(&total)
		}
	}
	return total.Int64
	/*
	rows, err := db.Query("select id from channel")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return calcRowsLen(rows)
	*/
}

func QueryRuleAllTotal(db *sql.DB) int {
	rows, err := db.Query("select id from rules")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return calcRowsLen(rows)
}

func calcRowsLen(rows *sql.Rows) (int) {
	count := 0
	for rows.Next() {
		count ++
	}
	return count
}

func QueryAllUsersInfo(db *sql.DB, currSize, pageSize int64, phoneNum, loginName, nickName string) []model.User {
	var allUser []model.User
	var rows *sql.Rows
	var err error
	var querySql string
	offset, limit := genOffset(currSize, pageSize)
    if phoneNum == "" && loginName == "" && nickName == "" {
		querySql = "select id, login_name, nick_name, phone_num, password, user_type, email, state, username, ctime from auth_user limit ? offset ?"
		rows, err = db.Query(querySql, limit, offset)
	} else {
		var intArr []string
		var whereMap = map[string]string{}
		whereMap["phone_num"] = phoneNum
		whereMap["login_name"] = loginName
		whereMap["nick_name"] = nickName
		where := BuildWhere(whereMap, intArr)
		if where != "" {
			var querySql = "select id, login_name, nick_name, phone_num, password, user_type, email, state, username, ctime from auth_user where %s limit %d offset %d"
			querySql = fmt.Sprintf(querySql, where, limit, offset)
			log.Println("sql", querySql)
			rows, err = db.Query(querySql)
		}
	}
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var login_name, nick_name, phone_num, password, email, username, ctime sql.NullString
		var id sql.NullInt64
		var user_type, state sql.NullInt64
		var user model.User
		err := rows.Scan(&id, &login_name, &nick_name, &phone_num, &password, &user_type, &email, &state, &username, &ctime)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("create_time", ctime)

		user.Id = id.Int64
		user.Login_name = login_name.String
		user.Password = password.String
		user.User_type = user_type.Int64
		user.User_role = user_type.Int64

		if nick_name.Valid {
			user.Nick_name = nick_name.String
		} else {
			user.Nick_name = ""
		}

        if phone_num.Valid {
			user.Phone_num = phone_num.String
		} else {
			user.Phone_num = ""
		}

		if email.Valid {
			user.Email = email.String
		}

		user.State = state.Int64
		if username.Valid {
			user.Username = username.String
		} else {
			user.Username = ""
		}

		if ctime.Valid {
			user.Ctime = ctime.String
		} else {
			user.Ctime = ""
		}

		log.Println(user)
		allUser = append(allUser, user)
	}
	log.Println("allUser:", allUser)
	return allUser
}

func QueryAllChannelInfo(db *sql.DB, currSize, pageSize int64, isPrepayment, isValid, channelName, phoneNum string) []model.Channel {
	var allChannel []model.Channel
	var rows *sql.Rows
	var err error
	var querySql string
	offset, limit := genOffset(currSize, pageSize)
	if isPrepayment == "" && isValid == "" && channelName == "" && phoneNum == "" {
		querySql = "select id, userid, remain_times, training_amt_per, divide_percent, status, is_valid, is_prepayment, ctime, utime, channel_name from channel limit %d offset %d "
		querySql = fmt.Sprintf(querySql, limit, offset)
		log.Println("sql", querySql)
		rows, err = db.Query(querySql)
		if err != nil {
			panic(err)
		}
	} else {
		var intArr []string
		intArr = append(intArr, "is_prepayment")
		intArr = append(intArr, "is_valid")
		var whereMap = map[string]string{}
		whereMap["is_prepayment"] = isPrepayment
		whereMap["is_valid"] = isValid
		whereMap["channel_name"] = channelName
		whereMap["phone_num"] = phoneNum
		where := BuildWhere(whereMap, intArr)
		if where != "" {
			querySql = "select id, userid, remain_times, training_amt_per, divide_percent, status, is_valid, is_prepayment, ctime, utime, channel_name from channel where %s limit %d offset %d "
			querySql = fmt.Sprintf(querySql, where, limit, offset)
			log.Println("sql where", querySql)
			rows, err = db.Query(querySql)
			if err != nil {
				panic(err)
			}
		}
	}

	for rows.Next() {
		var channelName, cTime, uTime sql.NullString
		var id, userId sql.NullInt64
		var remainTimes, trainingAmtPer, status, isValid, isPrepayment sql.NullInt64
		var dividePercent sql.NullFloat64
		var channel model.Channel
		err := rows.Scan(&id, &userId, &remainTimes, &trainingAmtPer, &dividePercent, &status, &isValid, &isPrepayment, &cTime, &uTime, &channelName)
        if err != nil {
        	panic(err)
		}
		channel.Id = id.Int64
		channel.Userid = userId.Int64
		channel.Remain_times = remainTimes.Int64
		channel.Training_amt_per = trainingAmtPer.Int64
		channel.Divide_percent = dividePercent.Float64
		channel.Status = status.Int64
		channel.Is_valid = isValid.Int64
		channel.Is_prepayment = isPrepayment.Int64
		if cTime.Valid {
			channel.Ctime = cTime.String
		} else {
			channel.Ctime = ""
		}
		if uTime.Valid {
			channel.Utime = uTime.String
		} else {
			channel.Utime = ""
		}
		if channelName.Valid {
			channel.Channel_name = channelName.String
		} else {
			channel.Channel_name = ""
		}
        allChannel = append(allChannel, channel)
	}
	return allChannel
}

func QueryChannelNames(db *sql.DB) []string {
	var names []string
	var channel_name sql.NullString
	rows, err := db.Query("select channel_name from channel")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&channel_name)
		if err != nil {
			panic(err)
		}
		if channel_name.Valid {
			names = append(names, channel_name.String)
		}
	}
	return names
}

func QueryRuleInfo(db *sql.DB, currSize, pageSize int64) []model.Rule {
    var allRule []model.Rule
	offset, limit := genOffset(currSize, pageSize)
	rows, err := db.Query("select id, name, total_amt, training_times, description, ctime, utime, is_valid from rules limit ? offset ? ", limit, offset)
	if err != nil {
		panic(err)
	}
	var id, totalAmt, trainingTimes, isValid sql.NullInt64
	var name, description, cTime, uTime sql.NullString
	for rows.Next() {
		var rule model.Rule
		err := rows.Scan(&id, &name, &totalAmt, &trainingTimes, &description, &cTime, &uTime, &isValid)
		if err != nil {
			panic(err)
		}
		rule.Id = id.Int64
		rule.Name = name.String
		rule.TotalAmt = totalAmt.Int64
		rule.TrainingTimes = trainingTimes.Int64
		rule.Description = description.String
		rule.IsValid = isValid.Int64
		if isValid.Int64 == 0 {
			rule.Status = "启用"
		} else {
			rule.Status = "关闭"
		}
		if cTime.Valid {
			rule.CTime = cTime.String
		} else {
			rule.CTime = ""
		}
		if uTime.Valid {
			rule.UTime = uTime.String
		} else {
			rule.UTime = ""
		}
		allRule = append(allRule, rule)
	}
	return allRule
}

func QueryRuleNames(db *sql.DB) []string {
	var names []string
	var name sql.NullString
	rows, err := db.Query("select name from rules")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		names = append(names, name.String)
	}
	return names
}

func UpdatePasswordById(userId int64, password string) {
	db := GetConn()
	defer db.Close()
	var updateSql = "update auth_user set password = ? where id = ?"
	_, err := db.Exec(updateSql, password, userId)
	if err != nil {
		panic(err)
		return
	}
}
