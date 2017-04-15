package db

import (
	"fmt"
	"sync"

	"github.com/go-xorm/xorm"
)

// OperationBase 基础Operation定义
type OperationBase struct {
	engine       *xorm.Engine
	locker       *sync.Mutex
	addFunc      GernalFunc
	getFunc      GernalFunc
	removeFunc   GernalFunc
	updateFunc   GernalFunc
	UpdateSqlCmd string
	OperationKey string
}

// Init 定义
func (ths *OperationBase) Init(e *xorm.Engine) error {
	if e == nil {
		return fmt.Errorf("[OperationBase:Init] Xorm engine can not be null")
	}
	ths.locker = &sync.Mutex{}
	ths.engine = e
	ths.getFunc = ths.get
	ths.addFunc = ths.add
	ths.removeFunc = ths.remove
	ths.updateFunc = ths.update
	ths.OperationKey = "OperationBase"
	return nil
}

// GetEngine 获取数据库引擎，如果需要操作必须注意使用线程锁
func (ths *OperationBase) GetEngine() *xorm.Engine {
	ths.locker.Lock()
	defer ths.locker.Unlock()
	return ths.engine
}

// GetKey 获取Key值
func (ths *OperationBase) GetKey() string {
	return ths.OperationKey
}

// SetGernalFun 设置通用方法的处理方法
func (ths *OperationBase) SetGernalFun(t GernalFuncType, f GernalFunc) {
	switch t {
	case AddFuncType:
		ths.addFunc = f
	case GetFuncType:
		ths.getFunc = f
	case RemoveFuncType:
		ths.removeFunc = f
	case UpdateFuncType:
		ths.updateFunc = f
	}
}

// Query 操作
func (ths *OperationBase) Query(qtype int, v ...interface{}) error {
	if v == nil || len(v) < 1 {
		return fmt.Errorf("[OperationBase:Query()] Query function , must input typeof struct pointer as parameter")
	}

	ths.locker.Lock()
	defer ths.locker.Unlock()
	switch qtype {
	case QtAddOneRecord:
		if ths.addFunc == nil {
			return fmt.Errorf("[OperationBase:Query()] Please sign to addFunc before you call this function")
		}
		return ths.addFunc(v[0])
	case QtGetOneRecord:
		if ths.getFunc == nil {
			return fmt.Errorf("[OperationBase:Query()] Please sign to getFunc before you call this function")
		}
		return ths.getFunc(v[0])
	case QtDeleteOneRecord:
		if ths.removeFunc == nil {
			return fmt.Errorf("[OperationBase:Query()] Please sign to removeFunc before you call this function")
		}
		return ths.removeFunc(v[0])
	case QtUpdateOneRecord:
		if ths.updateFunc == nil {
			return fmt.Errorf("[OperationBase:Query()] Please sign to updateFunc before you call this function")
		}
		return ths.updateFunc(v[0])
	case QtGetLastRecords:
		if len(v) != 3 {
			return fmt.Errorf("[OperationBase:Query()] Input parameter has to 3, Parameter cnt,orderkey,typeof struct")
		}
		return ths.getLastRecord(v[0].(int), v[1].(string), v[2])
	case QtQuaryRecords:
		if len(v) != 5 {
			return fmt.Errorf("[OperationBase:Query()] Input parameter has to 5, Parameter conditions,orderkey,cnt,isdesc,typeof struct")
		}
		return ths.getRecords(v[0].(string), v[1].(string), v[2].(int), v[3].(bool), v[4])
	case QtQuaryAllRecords:
		if len(v) != 4 {
			return fmt.Errorf("[OperationBase:Query()] Input parameter has to 4, Parameter conditions,orderkey,isdesc,typeof struct")
		}
		return ths.getAllRecords(v[0].(string), v[1].(string), v[3].(bool), v[4])
	}

	return fmt.Errorf("[OperationBase:Query] Query type is not defined (%d)", qtype)
}

func (ths *OperationBase) get(v interface{}) error {
	b, err := ths.engine.Get(v)
	if err != nil {
		return fmt.Errorf("[OperationBase:get()] %v", err)
	}
	if !b {
		return fmt.Errorf("[OperationBase:get()] The data that you want to get is not exist")
	}
	return nil
}

func (ths *OperationBase) add(v interface{}) error {
	_, err := ths.engine.InsertOne(v)
	return err
}

func (ths *OperationBase) remove(v interface{}) error {
	_, err := ths.engine.Delete(v)
	return err
}

func (ths *OperationBase) update(v interface{}) error {
	_, err := ths.engine.Update(v)
	return err
}

// GetLastRecord 得到最新的cnt条记录
func (ths *OperationBase) getLastRecord(cnt int, orderKey string, v interface{}) error {
	session := ths.engine.NewSession()
	defer session.Close()
	if cnt > 0 {
		session = session.Limit(cnt)
	}
	return session.Desc(orderKey).Find(v)
}

// GetRecords 得到一定条件的数据
func (ths *OperationBase) getRecords(conditions, orderKey string, cnt int, isdesc bool, v interface{}) error {
	session := ths.engine.NewSession()
	defer session.Close()
	if len(conditions) > 0 {
		session = session.Where(conditions)
	}
	if cnt > 0 {
		session = session.Limit(cnt)
	}
	if isdesc {
		return session.Desc(orderKey).Find(v)
	}
	return session.Asc(orderKey).Find(v)
}

// getAllRecords 得到一定条件的所有数据
func (ths *OperationBase) getAllRecords(conditions, orderKey string, isdesc bool, v interface{}) error {
	return ths.getRecords(conditions, orderKey, -1, isdesc, v)
}
