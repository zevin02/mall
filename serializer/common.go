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

// DataList 待有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  item,
			Total: total,
		},
		Msg: "ok",
	}
}
