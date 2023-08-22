package digests

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 11:03
//
import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	return MD5Bytes([]byte(str))
}

func MD5Bytes(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
