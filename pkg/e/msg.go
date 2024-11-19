package e

var MsgFlags = map[int]string{
	SUCCESS:            "ok",
	ERROR:              "fail",
	InvalidParams:      "请求参数错误",
	ErrorExistName:     "已存在用户名",
	ErrorFailEncoding:  "密码加密失败",
	ErrorExistUserName: "用户已存在",
	ErrorNotCompare:    "密码错误",
	ErrorAuthToken:     "token验证失败",
	ErrorExpiredToken:  "token过期",
	ErrorUploadFail:    "上传文件失败",
	ErrorExistFavorite: "收藏夹已存在",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
