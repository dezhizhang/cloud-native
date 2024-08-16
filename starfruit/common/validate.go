package common

import "regexp"

// VerifyMobile 验证手机号合法性
func VerifyMobile(mobile string) bool {

	if len(mobile) != 11 {
		return false
	}

	regular := "/^(((13[0-9]{1})|(14[0-9]{1})|(17[0-9]{1})|(15[0-9]{1})|(18[0-9]{1}))+\\d{8})$/"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
