package controllers


import (
    "github.com/garyburd/redigo/redis"
	"build_web/GoPractice/dlog"
)

func SetSession(key string, val string) {
	var log = dlog.DcLog()
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	_, err = c.Do("SET", key, val, "EX", "5000")
	if err != nil {
		log.Println("redis set failed:", err)
	}
}

/*
func GetSession(key string) bool {

}
*/