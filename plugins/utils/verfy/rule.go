package verfy

import (
	"regexp"
	"strings"
)

// space 当前字段的值不能是空白格
func space(s string) bool {
	return strings.TrimSpace(s) != ""
}

// tel 手机号匹配
func tel(tel string) bool {
	regex := `^1(?:3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[13578])\d{8}$`
	re := regexp.MustCompile(regex)

	// 使用正则表达式进行匹配
	return re.MatchString(tel)
}

// password 包含大小写字母、数字，长度在6到30之间
func password(password string) bool {
	// 正则表达式，规则：包含大小写字母、数字，长度在6到30之间
	regex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{6,30}$`
	re := regexp.MustCompile(regex)

	// 使用正则表达式进行匹配
	return re.MatchString(password)
}
