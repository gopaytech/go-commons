package postgresql

import (
	"fmt"

	"github.com/gopaytech/go-commons/pkg/db"
	gorm2 "github.com/gopaytech/go-commons/pkg/zlog/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config db.Config, gormConfig *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", config.Host, config.Username, config.Password, config.DatabaseName, config.Port)
	return gorm.Open(postgres.Open(dsn), gormConfig)
}

func ConnectDefault(config db.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", config.Host, config.Username, config.Password, config.DatabaseName, config.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gorm2.GormLogger,
	})
}
