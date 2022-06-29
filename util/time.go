package util

import (
	"strings"
	"time"
)

// TimeToString 格式化time.Time输出日期字符串
/*
| 标识 | 示例 | 描述 |
| YY | 22 | 年, 两位数 |
| YYYY | 2022 | 年，四位数 |
| M | 1-12 | 月，从1开始 |
| MM | 01-12 | 月，两位数字 |
| D | 1-31 | 日 |
| DD | 01-31 | 日，两位数 |
| HH | 00-23 | 24小时，两位数 |
| h | 1-12 | 12小时 |
| hh | 01-12 | 12小时，两位数 |
| m - 0-59 | 分钟 |
| mm | 00-59 | 分钟，两位数 |
| s | 0-59 | 秒 |
| ss | 00-59 | 秒，两位数 |
| Z | -05:00 | UTC偏移 |
| ZZ | -0500 | UTC偏移，两位数 |
| A | AM / PM | 上/下午，大写 |
| a | am / pm | 上/下午，小写 |
*/

func formatToLayout(format string) string {
	layout := strings.Replace(format, "YYYY", "2006", -1)
	layout = strings.Replace(layout, "YY", "06", -1)
	layout = strings.Replace(layout, "MM", "01", -1)
	layout = strings.Replace(layout, "M", "1", -1)
	layout = strings.Replace(layout, "DD", "02", -1)
	layout = strings.Replace(layout, "D", "2", -1)
	layout = strings.Replace(layout, "HH", "15", -1)
	layout = strings.Replace(layout, "hh", "03", -1)
	layout = strings.Replace(layout, "h", "3", -1)
	layout = strings.Replace(layout, "mm", "04", -1)
	layout = strings.Replace(layout, "m", "4", -1)
	layout = strings.Replace(layout, "ss", "05", -1)
	layout = strings.Replace(layout, "s", "5", -1)
	layout = strings.Replace(layout, "ZZ", "-0700", -1)
	layout = strings.Replace(layout, "Z", "-07:00", -1)
	layout = strings.Replace(layout, "A", "PM", -1)
	layout = strings.Replace(layout, "a", "pm", -1)
	return layout
}

func TimeToString(t time.Time, format string) string {
	layout := formatToLayout(format)
	return t.Format(layout)
}

func StringToTime(str, format string) (t time.Time, err error) {
	local, err := time.LoadLocation("Local")
	if err != nil {
		return
	}
	layout := formatToLayout(format)
	t, err = time.ParseInLocation(layout, str, local)
	return
}

// Date 获取日期字符串的当地Unix时间戳
// layout s的时间格式
func Date(s, layout string) (int64, error) {
	// 使用服务器设置的时区
	local, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	t, err := time.ParseInLocation(layout, s, local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
