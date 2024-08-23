package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// GenMd5 生成md5编码
func GenMd5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}
