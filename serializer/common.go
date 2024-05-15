package serializer

type Response struct {
	Status int    `json:"status"`
	Data   any    `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type TokenData struct {
	User  any    `json:"user"`
	Token string `json:"token"`
}

type DataList struct {
	Items any  `json:"items"`
	Total uint `json:"total"`
}

func BuildListResponse(items any, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Items: items,
			Total: total,
		},
		Msg: "ok",
	}
}