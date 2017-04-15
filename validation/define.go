package validation

import "github.com/astaxie/beego"

const (
	NoError            = 0
	CommomError        = 1
	UnknownFormatError = 2
	DataBaseError      = 3
	ParamInvalid       = 100
)

// IOperation 操作接口定义
type IOperation interface {
	GetResultData() *OperationResult
	QueryExecute() *OperationResult
	DecodeContext(ctl *beego.Controller)
}
