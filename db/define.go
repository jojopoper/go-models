package db

import "github.com/go-xorm/xorm"

// DatabaseType 数据库类型
type DatabaseType string

// GernalFunc 通用方法定义
type GernalFunc func(v interface{}) error

// RegModelsFunc 数据库注册表结构方法定义
type RegModelsFunc func(eng *xorm.Engine)

// InitOperationFunc 注册数据库Operation方法定义
type InitOperationFunc func()

// GernalFuncType 通用方法类型
type GernalFuncType int

const (
	QtAddOneRecord    = iota + 1 // 添加一条记录
	QtGetOneRecord               // 获取一条记录
	QtDeleteOneRecord            // 删除一条记录
	QtUpdateOneRecord            // 更新一条记录
	QtGetLastRecords             // 获取最近的N条记录
	QtQuaryRecords               // 获取指定的N条记录
	QtQuaryAllRecords            // 获取所有记录

	MySqlDriver    DatabaseType = "mysql"
	sqliteDriver   DatabaseType = "sqlite3"
	PostgresDriver DatabaseType = "postgres"

	AddFuncType    GernalFuncType = 1
	GetFuncType    GernalFuncType = 2
	RemoveFuncType GernalFuncType = 3
	UpdateFuncType GernalFuncType = 4

	LanguageChinese = "cn"
	LanguageEnglish = "en"
)

// DatabaseInfo 数据库基本信息定义
type DatabaseInfo struct {
	DbType    DatabaseType
	AliasName string
	Host      string
	Port      string
	UserName  string
	Password  string
	IsDebug   bool
}

// OperationInterface 数据库接口定义
type OperationInterface interface {
	Init(e *xorm.Engine) error
	GetKey() string
	Query(qtype int, v ...interface{}) error
	GetEngine() *xorm.Engine
}
