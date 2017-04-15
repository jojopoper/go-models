package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

// ManageBase 数据库管理基础定义
type ManageBase struct {
	dbInfo        DatabaseInfo
	dbEngine      *xorm.Engine
	operations    map[string]OperationInterface
	ormRegModels  RegModelsFunc
	initOperation InitOperationFunc
}

// Init 初始化数据库
func (ths *ManageBase) Init(dbConfig DatabaseInfo) {
	ths.dbInfo = dbConfig

	ths.initEngine()

	// 注册Orm的数据库表
	if ths.ormRegModels == nil {
		panic("[ManageBase:Init] Can not find 'RegModelsFunc' function point!")
	}
	ths.ormRegModels(ths.dbEngine)

	if ths.initOperation == nil {
		panic("[ManageBase:Init] Can not find 'InitOperationFunc' function point!")
	}
	ths.initOperation()
}

// SetRegModelsFunc 设置 RegModelsFunc 方法
func (ths *ManageBase) SetRegModelsFunc(f RegModelsFunc) {
	ths.ormRegModels = f
}

// SetInitOperaFunc 设置 InitOperationFunc 方法
func (ths *ManageBase) SetInitOperaFunc(f InitOperationFunc) {
	ths.initOperation = f
}

// GetEngine 获取数据库引擎
func (ths *ManageBase) GetEngine() *xorm.Engine {
	return ths.dbEngine
}

// AppendOperation 添加注册数据库Operation
func (ths *ManageBase) AppendOperation(o OperationInterface) {
	if ths.operations == nil {
		ths.operations = make(map[string]OperationInterface)
	}
	o.Init(ths.dbEngine)
	ths.operations[o.GetKey()] = o
}

func (ths *ManageBase) initEngine() {
	ths.dbEngine = nil
	var err error
	switch ths.dbInfo.DbType {
	case MySqlDriver:
		ths.dbEngine, err = ths.getMySQLEngine()
	case PostgresDriver:
		ths.dbEngine, err = ths.getPostgresEngine()
	}
	if ths.dbEngine == nil {
		panic(fmt.Errorf("[ManageBase:initEngine] Undefined db type = %s\n", ths.dbInfo.DbType))
	}
	if err != nil {
		panic(fmt.Errorf("[ManageBase:initEngine] init engine has error = \n%+v\n", err))
	}
	ths.dbEngine.ShowDebug = ths.dbInfo.IsDebug
	ths.dbEngine.ShowInfo = ths.dbInfo.IsDebug
	ths.dbEngine.ShowSQL = ths.dbInfo.IsDebug
	ths.dbEngine.ShowErr = true
	ths.dbEngine.ShowWarn = true
}

func (ths *ManageBase) getMySQLEngine() (*xorm.Engine, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", //Asia%2FShanghai
		ths.dbInfo.UserName, ths.dbInfo.Password, ths.dbInfo.Host, ths.dbInfo.Port, ths.dbInfo.AliasName)
	ret, err := xorm.NewEngine(string(MySqlDriver), dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("[ManageBase:getMySQLEngine] Create MySql has error! \r\n\t%v\r\n", err)
	}
	err = ret.Ping()
	if err != nil {
		return nil, fmt.Errorf("[ManageBase:getMySQLEngine] Create MySql Ping error! \r\n\t %v\r\n", err)
	}
	return ret, nil
}

func (ths *ManageBase) getPostgresEngine() (*xorm.Engine, error) {
	dataSourceName := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		ths.dbInfo.AliasName, ths.dbInfo.UserName, ths.dbInfo.Password, ths.dbInfo.Host, ths.dbInfo.Port)
	ret, err := xorm.NewEngine(string(PostgresDriver), dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("[Manager:getPostgresEngine] Create Postgres has error! \r\n\t%v\r\n", err)
	}
	return ret, nil
}

// GetOperation 得到对应的操作数据库控制器
func (ths *ManageBase) GetOperation(key string) OperationInterface {
	return ths.operations[key]
}

// Query 通用数据库执行接口
func (ths *ManageBase) Query(operationKey string, qtype int, val ...interface{}) error {
	return ths.operations[operationKey].Query(qtype, val...)
}

// Add 添加记录
func (ths *ManageBase) Add(operationKey string, val interface{}) error {
	return ths.operations[operationKey].Query(QtAddOneRecord, val)
}

// Get 获取记录
func (ths *ManageBase) Get(operationKey string, val interface{}) error {
	return ths.operations[operationKey].Query(QtGetOneRecord, val)
}

// Remove 删除记录
func (ths *ManageBase) Remove(operationKey string, val interface{}) error {
	return ths.operations[operationKey].Query(QtDeleteOneRecord, val)
}

// Update 更新记录
func (ths *ManageBase) Update(operationKey string, val interface{}) error {
	return ths.operations[operationKey].Query(QtUpdateOneRecord, val)
}

// GetLastRecord 获取最新cnt条记录
func (ths *ManageBase) GetLastRecord(operationKey string, val ...interface{}) error {
	return ths.operations[operationKey].Query(QtGetLastRecords, val...)
}

// GetRecords 获取cnt条记录
func (ths *ManageBase) GetRecords(operationKey string, val ...interface{}) error {
	return ths.operations[operationKey].Query(QtQuaryRecords, val...)
}

// GetAllRecords 获取所有记录
func (ths *ManageBase) GetAllRecords(operationKey string, val ...interface{}) error {
	return ths.operations[operationKey].Query(QtQuaryAllRecords, val...)
}
