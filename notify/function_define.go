package notify

import (
	"sync"
)

// ReportFunction 报告方法定义
type ReportFunction struct {
	locker   *sync.Mutex
	reporter []ReportFunc
}

// Init 初始化
func (ths *ReportFunction) Init() *ReportFunction {
	ths.locker = new(sync.Mutex)
	ths.reporter = make([]ReportFunc, 0)
	return ths
}

// AddReportFunc 添加状态报告函数
func (ths *ReportFunction) AddReportFunc(f ReportFunc) {
	ths.locker.Lock()
	defer ths.locker.Unlock()
	ths.reporter = append(ths.reporter, f)
}

// AddReportFuncs 添加状态报告函数
func (ths *ReportFunction) AddReportFuncs(f ...ReportFunc) {
	ths.locker.Lock()
	defer ths.locker.Unlock()
	ths.reporter = append(ths.reporter, f...)
}

// Notify 通知外部函数
func (ths *ReportFunction) Notify(sender interface{}, msg *ReportMessage) {
	ths.locker.Lock()
	defer ths.locker.Unlock()
	for _, itmFun := range ths.reporter {
		go itmFun(sender, msg)
	}
}

// GetReporters 获取所有的报告事件方法
func (ths *ReportFunction) GetReporters() []ReportFunc {
	ths.locker.Lock()
	defer ths.locker.Unlock()
	return ths.reporter
}
