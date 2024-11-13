package serializer

// Response 团队基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"` //这个返回给到用户
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
