package comment

type Response struct {}

const CODE_PARMAS_ERROR  = 9000
const CODE_SUCCESS  = 200
const CODE_SYSTEM_ERROR  = 500
const CODE_NOT_FOUND  = 404

func (*Response) JsonFormat(code int, data interface{}, msg string) (json map[string]interface{}) {
	json = map[string]interface{}{
		"code": code,
		"data": data,
		"msg":  msg,
	}
	return json
}
