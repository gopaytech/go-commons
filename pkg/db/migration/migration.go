package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
)

type Migration interface {
	Up() error
	DownLast() error
	Down() error
	Close() error
	Instance() *migrate.Migrate
}

type Migrator struct {
	Migrate *migrate.Migrate
}

func (m *Migrator) Up() (err error) {
	err = m.Migrate.Up()
	if err == migrate.ErrNoChange {
		log.Info().Msg("No new migration found, Ignoring!")
		return nil
	}
	return
}

func (m *Migrator) DownLast() (err error) {
	err = m.Migrate.Steps(-1)
	return
}

func (m *Migrator) Instance() *migrate.Migrate {
	return m.Migrate
}

func (m *Migrator) Down() (err error) {
	err = m.Migrate.Down()
	return
}

func (m *Migrator) Close() (err error) {
	_, err = m.Migrate.Close()
	return
}
