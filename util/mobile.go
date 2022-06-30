package util

import "strings"

// MobileEncode 手机号加星星, 显示前3位开头和4位尾号
func MobileEncode(mobile string) string {
	if len(mobile) < 8 {
		return mobile
	}
	prefix := mobile[0:3]
	sufffix := mobile[len(mobile)-4:]
	star := strings.Repeat("*", len(mobile)-7)
	return prefix + star + sufffix
}
