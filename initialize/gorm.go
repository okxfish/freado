package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormRouter() *gorm.DB {
	return GormMysql()
	//switch global.GVA_CONFIG.System.DbType {
	//case "mysql":
	//	return GormMysql()
	//default:
	//	return GormSqlite()
	//}
}

func GormMysql() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/freado?charset=utf8&parseTime=True&loc=Local"
	//dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		panic("failed to connect database")
		return nil
	}
	return db
}
