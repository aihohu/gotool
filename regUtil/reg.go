package regUtil

import (
	"github.com/dlclark/regexp2"
	"regexp"
)

// Email 校验邮箱地址
func Email(email string) bool {
	matched, err := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", email)
	if err != nil {
		return false
	}

	return matched
}

// Phone 校验手机号码
func Phone(phone string) bool {
	matched, err := regexp.MatchString("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", phone)
	if err != nil {
		return false
	}
	return matched
}

// Password 校验密码强度
// 至少包含：数字,英文,字符中的两种以上，长度6-20
func Password(password string) (string, bool) {
	expr := `^(?![0-9]+$)(?![a-z]+$)(?![A-Z]+$)(?!([^(0-9a-zA-Z)])+$).{6,20}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(password)
	if m != nil {
		return "", true
	}
	return "至少包含：数字,英文,字符中的两种以上，长度6-20", false
}

// Domain 域名校验
func Domain(domain string) bool {
	matched, err := regexp.MatchString("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", domain)
	if err != nil {
		return false
	}
	return matched
}
