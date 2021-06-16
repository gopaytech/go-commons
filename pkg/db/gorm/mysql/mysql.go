package mysql

import (
	"fmt"
	"github.com/gopaytech/go-commons/pkg/db"
	"github.com/gopaytech/go-commons/pkg/zlog"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(config db.Config, gormConfig *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)
	return gorm.Open(mysqlDriver.Open(dsn), gormConfig)
}

func ConnectDefault(config db.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)
	return gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{
		Logger: zlog.GormLogger,
	})
}
