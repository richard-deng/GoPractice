package controllers

import (
	"log"
	"build_web/GoPractice/model"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

func QueryUsersAllTotal(db *sql.DB) int {
	rows, err := db.Query("select id, login_name, nick_name, phone_num, password, user_type, email, state, username, ctime from auth_user")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return calcRowsLen(rows)
}

func QueryChannelAllTotal(db *sql.DB) int {
	rows, err := db.Query("select id from channel")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return calcRowsLen(rows)
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

func QueryAllUsersInfo(db *sql.DB, currSize, pageSize int64) []model.User {
	var allUser []model.User
	offset, limit := genOffset(currSize, pageSize)
	rows, err := db.Query("select id, login_name, nick_name, phone_num, password, user_type, email, state, username, ctime from auth_user limit ? offset ?", limit, offset)

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
	return allUser
}

func QueryAllChannelInfo(db *sql.DB, currSize, pageSize int64) []model.Channel {
	var all_channel []model.Channel
	offset, limit := genOffset(currSize, pageSize)
	rows, err := db.Query("select id, userid, remain_times, training_amt_per, divide_percent, status, is_valid, is_prepayment, ctime, utime, channel_name from channel limit ? offset ? ", limit, offset)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var channel_name, ctime, utime sql.NullString
		var id, userid sql.NullInt64
		var remain_times, training_amt_per, status, is_valid, is_prepayment sql.NullInt64
		var divide_percent sql.NullFloat64
		var channel model.Channel
		err := rows.Scan(&id, &userid, &remain_times, &training_amt_per, &divide_percent, &status, &is_valid, &is_prepayment, &ctime, &utime, &channel_name)
        if err != nil {
        	panic(err)
		}
		channel.Id = id.Int64
		channel.Userid = userid.Int64
		channel.Remain_times = remain_times.Int64
		channel.Training_amt_per = training_amt_per.Int64
		channel.Divide_percent = divide_percent.Float64
		channel.Status = status.Int64
		channel.Is_valid = is_valid.Int64
		channel.Is_prepayment = is_prepayment.Int64
		if ctime.Valid {
			channel.Ctime = ctime.String
		} else {
			channel.Ctime = ""
		}
		if utime.Valid {
			channel.Utime = utime.String
		} else {
			channel.Utime = ""
		}
		if channel_name.Valid {
			channel.Channel_name = channel_name.String
		} else {
			channel.Channel_name = ""
		}
        all_channel = append(all_channel, channel)
	}
	return all_channel
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
