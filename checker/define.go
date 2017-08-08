package checker

import (
	_n "github.com/jojopoper/go-models/notify"
)

// executeFunction 执行函数的定义
type executeFunction func()

// RunFlag 运行标识
type RunFlag int

const (
	ForceStop RunFlag = 1
)

// ICheckInterface 线程运行接口定义
type ICheckInterface interface {
	Name() string
	Start()
	Stop()
	IsRunning() bool
	IsBeginStart() bool
	RegistManager(m map[string]ICheckInterface) ICheckInterface
	AddReportFunc(f _n.ReportFunc)
	AddReportFuncs(f ..._n.ReportFunc)
}
