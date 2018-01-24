package controllers


import(
	"time"
	"math/rand"
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"build_web/GoPractice/dlog"
)



const (
	KC_RAND_KIND_NUM   = 0	// 纯数字
	KC_RAND_KIND_LOWER = 1	// 小写字母
	KC_RAND_KIND_UPPER = 2	// 大写字母
	KC_RAND_KIND_ALL   = 3 	// 数字、大小写字母
)

// 随机字符串
func RandStr(size int, kind int) []byte {
	iKind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if isAll { // random iKind
			iKind = rand.Intn(3)
		}
		scope, base := kinds[iKind][0], kinds[iKind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}

func GenPassword(md5Password string) string {
	var log = dlog.DcLog()
	result := RandStr(5, KC_RAND_KIND_ALL)
	randomStr := string(result)
	log.Println(randomStr)
	newStr := randomStr + md5Password
	log.Println(newStr)
    r := sha1.Sum([]byte(newStr))
	sha1Hex := hex.EncodeToString(r[:])
	log.Println(sha1Hex)
	//return string(sha1Hex)
	return "sha1" + "$" + randomStr + "$" + string(sha1Hex)
}

func CheckPassword(fullPassword string, md5Password string) bool {
	var log = dlog.DcLog()
	log.Printf("full password=%s", fullPassword)
	all := strings.Split(fullPassword, "$")
	salt := all[1]
	password := all[2]
	log.Printf("salt=%s|password=%s", salt, password)
	newStr := salt + md5Password
	log.Printf("new str=%s", newStr)
	r := sha1.Sum([]byte(newStr))
	sha1Hex := hex.EncodeToString(r[:])
	log.Println(sha1Hex)
    if string(sha1Hex) != password {
    	return false
	}
	return true
}