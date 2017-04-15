package validation

// OperationResult 反馈结果结构定义
type OperationResult struct {
	ErrorMsg   string      `json:"error"`
	CodeID     int         `json:"codeid"`
	ResultData interface{} `json:"data"`
	Language   string      `json:"language"`
}

// SetError 设置错误信息
func (ths *OperationResult) SetError(id int, msg string) {
	ths.CodeID = id
	ths.ErrorMsg = msg
}
