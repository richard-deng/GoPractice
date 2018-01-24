package model

import (
	"build_web/GoPractice/dlog"
)

type Response struct {
	Respcd 	string	`json:"respcd"`
	Resperr string	`json:"resperr"`
	Respmsg string	`json:"respmsg"`
	Data    interface{} `json:"data"`
}


//struct定义的属性是小写开头的，不是public的，这样是不能跨包调用
type User struct {
    Id    			int64	`json:"id"`
    Login_name		string	`json:"login_name"`
    Nick_name		string	`json:"nick_name"`
    Phone_num		string	`json:"phone_num"`
    Password		string	`json:"password"`
    User_type		int64	`json:"user_type"`
    User_role       int64   `json:"user_role"`
    Email           string	`json:"email"`
    State           int64	`json:"state"`
    Username		string	`json:"username"`
    Ctime           string  `json:"ctime"`
}

type Channel struct {
	Id				int64   `json:"id"`
	Userid          int64   `json:"userid"`
	Remain_times    int64   `json:"remain_times"`
	Training_amt_per  int64  `json:"training_amt_per"`
	Divide_percent  float64 `json:"divide_percent"`
	Status          int64    `json:"status"`
	Is_valid        int64   `json:"is_valid"`
	Is_prepayment   int64    `json:"is_prepayment"`
	Ctime           string  `json:"ctime"`
	Utime           string  `json:"utime"`
	Channel_name    string  `json:"channel_name"`

}

type Rule struct {
	Id    			int64     `json:"id"`
	Name            string    `json:"name"`
	TotalAmt        int64     `json:"total_amt"`
    TrainingTimes   int64     `json:"training_times"`
    Description     string    `json:"description"`
    CTime           string    `json:"ctime"`
    UTime           string    `json:"utime"`
    IsValid         int64     `json:"is_valid"`
    Status          string    `json:"state"`
}

func (u User) Valid() bool {
	dlog.Info.Printf("user id=%d", u.Id)
	if u.Id != 0 {
		return true
	}
	return false
}