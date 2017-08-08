package notify

const (
	NoError     int = 0
	CommonError int = 100
	NotifyMsg   int = 1000
)

// ReportFunc 报告状态事件定义
type ReportFunc func(sender interface{}, msg *ReportMessage)
