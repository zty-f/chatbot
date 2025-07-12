package service

import "strings"

// SeniorFriendly 对AI回复进行适老化处理
func SeniorFriendly(reply string) string {
	reply = strings.ReplaceAll(reply, "哈哈", "很开心呢")
	return "亲爱的，" + reply
}
