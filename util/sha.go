// Package util go实用库
// 安全散列算法SHA常用函数
package util

import (
	"crypto/sha256"
	"errors"
)

// SHA256 加密
func SHA256(s []byte) (r []byte, err error) {
	h := sha256.New()
	n, err := h.Write(s)
	if err != nil {
		return
	}
	if n != len(s) {
		err = errors.New("Write length error")
		return
	}
	r = h.Sum(nil)
	return
}
