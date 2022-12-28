package desensitizedUtil

import (
	"github.com/aihohu/gotool/stringUtil"
	"strings"
)

// Hide 数据脱敏
// @param idCardNum 数据
// @param front     保留：前面的front位数；从1开始
// @param end       保留：后面的end位数；从1开始
// @return          脱敏后的数据
func Hide(idCardNum string, front int, end int) string {
	return generateDesensitized(idCardNum, front, end)
}

// IdCard 身份证号码脱敏 保留前4位和后3位
// @param idCardNum 身份证号码
// @return          脱敏后的身份证号码
func IdCard(idCardNum string) string {
	return generateDesensitized(idCardNum, 4, 3)
}

// MobilePhone  手机号码脱敏 保留前3位和后4位
// @param mobilePhone 手机号码
// @return      脱敏后的手机号码
func MobilePhone(mobilePhone string) string {
	return generateDesensitized(mobilePhone, 3, 4)
}

// Telephone 固定电话 保留前4位和后2位
// @param telephone 固定电话
// @return          脱敏后的固定电话
func Telephone(telephone string) string {
	return generateDesensitized(telephone, 4, 2)
}

// Email 电子邮箱 邮箱前缀仅显示第一个字母，其他用*号代替，@及后面的地址显示，比如：g**@163.com
// @param email 邮箱
// @return      脱敏后的邮箱
func Email(email string) string {

	if stringUtil.IsBlank(email) {
		return ""
	}

	index := strings.Index(email, "@")
	if index <= 1 {
		return email
	}

	return generateDesensitized(email, 1, index)
}

// Address 地址
// @param address 地址
// @param count   敏感信息长度
// @return        脱敏后的地址
func Address(address string, count int) string {
	// 中文占用3个字节长度
	start := len(address) - (count * 3)
	desensitized := generateDesensitized(address, start, 0)
	return desensitized[0 : len(desensitized)-(count*2)]
}

func generateDesensitized(str string, start int, end int) string {
	//str不能为空
	if stringUtil.IsBlank(str) {
		return ""
	}

	//需要截取的不能小于0
	if start < 0 || end < 0 {
		return ""
	}

	num := len(str)
	//需要截取的长度不能大于str长度
	if (start + end) > num {
		return ""
	}

	count := num - start - end
	return str[0:start] + strings.Repeat("*", count) + str[num-end:num]
}

// 生成指定数量*号
func generateAsterisk(count int) string {
	return strings.Repeat("*", count)
}
