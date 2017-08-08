package checker

import (
	"time"

	_n "github.com/jojopoper/go-models/notify"
)

// CheckBase 线程检查基类
type CheckBase struct {
	Interval    int          // 周期检测间隔 单位：秒
	stop        chan RunFlag // 强制停止标志
	isRunning   bool         // 标示是否正在运行
	beginStart  bool         // 是否为默认启动监控
	continueRun bool         // 对于非默认监控的线程，标定是否可以结束
	keyName     string       // 关键字名称
	exe         executeFunction
	_n.ReportFunction
}

// Init 初始化
func (ths *CheckBase) Init(interval int) {
	ths.ReportFunction.Init()
	ths.Interval = interval
	ths.stop = make(chan RunFlag)
	ths.isRunning = false
	ths.beginStart = true
	ths.continueRun = false
}

// Name 获取名称
func (ths *CheckBase) Name() string {
	return ths.keyName
}

// SetName 设置名称
func (ths *CheckBase) SetName(n string) {
	ths.keyName = n
}

// SetExeFunc 设置执行函数
func (ths *CheckBase) SetExeFunc(e executeFunction) {
	ths.exe = e
}

// Start 开始
func (ths *CheckBase) Start() {
	if ths.exe != nil {
		go ths.do()
	} else {
		panic("Have to make executeFunction object!")
	}
}

func (ths *CheckBase) do() {
	if ths.isRunning {
		return
	}
	ths.isRunning = true
	t1 := time.NewTimer(time.Second * time.Duration(ths.Interval))
	for {
		select {
		case <-t1.C:
			ths.exe()
			if ths.beginStart || ths.continueRun {
				t1.Reset(time.Second * time.Duration(ths.Interval))
			} else {
				go ths.Stop()
			}
		case isStop := <-ths.stop:
			if isStop == ForceStop {
				t1.Stop()
				ths.isRunning = false
				return
			}
		}
	}
}

// Stop 开始
func (ths *CheckBase) Stop() {
	if ths.isRunning {
		ths.stop <- ForceStop
	}
}

// IsRunning 是否在运行
func (ths *CheckBase) IsRunning() bool {
	return ths.isRunning
}

// IsBeginStart 是否为默认启动监控
func (ths *CheckBase) IsBeginStart() bool {
	return ths.beginStart
}

// SetContinue 设置是否继续标识
func (ths *CheckBase) SetContinue(b bool) {
	ths.continueRun = b
}

// SetBeginStart 设置默认启动监控
func (ths *CheckBase) SetBeginStart(b bool) {
	ths.beginStart = b
}

// BeginStop 设置beginStart状态
func (ths *CheckBase) BeginStop() {
	ths.beginStart = false
}

// RegistManager 将自己注册进入管理map
func (ths *CheckBase) RegistManager(m map[string]ICheckInterface) ICheckInterface {
	if m != nil {
		m[ths.keyName] = ths
	}
	return ths
}
