package notify

// ReportMessage 报告消息定义
type ReportMessage struct {
	StatusID int
	Err      error
	Msg      string
	Datas    interface{}
}

// SetCheckMessage 设置消息
func (ths *ReportMessage) SetCheckMessage(id int, e error, m string, d interface{}) *ReportMessage {
	ths.StatusID = id
	ths.Err = e
	ths.Msg = m
	ths.Datas = d
	return ths
}

// SetError 设置错误
func (ths *ReportMessage) SetError(id int, e error) *ReportMessage {
	return ths.SetCheckMessage(id, e, "", nil)
}

// GetError 获取错误
func (ths *ReportMessage) GetError() error {
	return ths.Err
}

// SetMessage 设置消息
func (ths *ReportMessage) SetMessage(id int, m string) *ReportMessage {
	return ths.SetCheckMessage(id, nil, m, nil)
}

// GetMessage 获取消息
func (ths *ReportMessage) GetMessage() string {
	return ths.Msg
}

// SetData 设置数据
func (ths *ReportMessage) SetData(id int, d interface{}) *ReportMessage {
	return ths.SetCheckMessage(id, nil, "", d)
}

// GetData 获取数据
func (ths *ReportMessage) GetData() interface{} {
	return ths.Datas
}
