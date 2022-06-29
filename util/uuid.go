package util

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

// UUID 获取uuid
func UUID() (uuid string, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	id := node.Generate()
	uuid = reverseString(id.Base36())
	return
}

// ParseTimeByUUID 根据uuid, 获取生成该uuid的时间
func ParseTimeByUUID(uuid string) (dateTime string, err error) {
	str := reverseString(uuid)
	id, err := snowflake.ParseBase36(str)
	if err != nil {
		return
	}
	timestamp := id.Time()
	sec := timestamp / 1000
	nec := (timestamp - sec*1000) * 1000000
	tm := time.Unix(sec, nec)
	dateTime = tm.Format("2006-01-02 15:04:05")
	return
}

// ReverseString 反转字符串
// abc 反转后结果为cba
func reverseString(s string) string {
	r := []rune(s)
	length := len(r)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
