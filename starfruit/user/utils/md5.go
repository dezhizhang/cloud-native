package utils

import (
	"crypto/md5"
	"encoding/hex"
	"starfruit.top/user/global"
)

// GenMd5 生成md5编码
func GenMd5(str string) string {
	fullData := str + global.SALT
	// 计算MD5摘要
	hash := md5.Sum([]byte(fullData))
	// 将摘要转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}
