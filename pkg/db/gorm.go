package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"hpf/common/config"
)

type DbGorm struct {
	Type         string
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
}

func NewMysqlConn(mysql config.Mysql) *gorm.DB {
	dbGorm := DbGorm{Dsn: mysql.Dsn()}
	dbGorm.MaxIdleConns = mysql.MaxIdleConns
	dbGorm.MaxOpenConns = mysql.MaxOpenConns
	return dbGorm.GormMysql()
}

func (dg *DbGorm) GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       dg.Dsn, // DSN data source name
		DefaultStringSize:         191,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   false,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,  // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic(err)
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dg.MaxOpenConns)
	return db
}
