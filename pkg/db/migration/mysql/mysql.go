package mysql

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gopaytech/go-commons/pkg/db"
	"github.com/gopaytech/go-commons/pkg/db/migration"
)

func New(config db.Config, path string) (mgrn migration.Migration, err error) {
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%v)/%s", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)
	m, err := migrate.New(path, dsn)
	if err != nil {
		return
	}

	mgrn = &migration.Migrator{
		Migrate: m,
	}

	return
}

func WithInstance(db *sql.DB, path string) (mgrn migration.Migration, err error) {
	return WithInstanceConfig(db, path, &mysql.Config{})
}

func WithInstanceMigrationTable(db *sql.DB, path string, migrationTable string) (mgrn migration.Migration, err error) {
	return WithInstanceConfig(db, path, &mysql.Config{
		MigrationsTable: migrationTable,
	})
}

func WithInstanceConfig(db *sql.DB, path string, config *mysql.Config) (mgrn migration.Migration, err error) {
	driver, err := mysql.WithInstance(db, config)
	m, _ := migrate.NewWithDatabaseInstance(
		path,
		"mysql",
		driver,
	)

	if err != nil {
		return
	}

	mgrn = &migration.Migrator{
		Migrate: m,
	}

	return
}
