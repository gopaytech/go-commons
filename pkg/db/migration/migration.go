package migration

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gopaytech/go-commons/pkg/db"
	"github.com/rs/zerolog/log"
)

type Migration interface {
	Up() error
	DownLast() error
	Close() error
}

type migrator struct {
	migrate *migrate.Migrate
}

func (m *migrator) Up() (err error) {
	err = m.migrate.Up()
	if err == migrate.ErrNoChange {
		log.Info().Msg("No new migration found, Ignoring!")
		return nil
	}
	return
}

func (m *migrator) DownLast() (err error) {
	err = m.migrate.Steps(-1)
	return
}

func (m *migrator) Close() (err error) {
	_, err = m.migrate.Close()
	return
}

func New(config db.Config, path string) (mgrn Migration, err error) {
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%v)/%s", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)
	m, err := migrate.New(path, dsn)
	if err != nil {
		return
	}

	mgrn = &migrator{
		migrate: m,
	}

	return
}

func WithInstance(db *sql.DB, path string) (mgrn Migration, err error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		path,
		"mysql",
		driver,
	)

	if err != nil {
		return
	}

	mgrn = &migrator{
		migrate: m,
	}

	return
}
