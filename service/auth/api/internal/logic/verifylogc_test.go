package logic

import "testing"

func Test_KeyMatch(t *testing.T) { // 功能测试以 `Test` 前缀命名
	if key2, matched, err := KeyMatch("/service/auth", "/service/:name"); err == nil {
		// t.Errorf 类似fmt.Printf()
		if !matched {
			t.Errorf("'/service/auth/config/1/apply' matched %s: %v, want true", key2, matched)
		}
	}

	if key2, matched, err := KeyMatch("/service/auth/config/1/apply", "/service/:name"); err == nil {
		// t.Errorf 类似fmt.Printf()
		if matched {
			t.Errorf("'/service/auth/config/1/apply' matched %s: %v, want false", key2, matched)
		}
	}
}
