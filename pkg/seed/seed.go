package seed

import (
	"database/sql"
	"github.com/gopaytech/go-commons/pkg/dir"
	"github.com/gopaytech/go-commons/pkg/zlog"
	"io/ioutil"
	"path/filepath"
)

type Seeder func(path string) error

func ProvideSeeder(db *sql.DB) Seeder {
	return func(path string) error {
		zlog.Debug("seed source dir: %s", path)
		files, err := dir.ListFiles(path)
		if err != nil {
			return err
		}
		for _, file := range files {
			fullPath := filepath.Join(path, file)
			zlog.Debug("seed file: %s", fullPath)
			openedFile, err := ioutil.ReadFile(fullPath)
			if err != nil {
				return err
			}
			sqlText := string(openedFile)
			result, err := db.Exec(sqlText)
			if err != nil {
				zlog.Error(err, "execute: %s failed", sqlText)
				return err
			}
			rowAffected, _ := result.RowsAffected()
			zlog.Debug("execute: %s, row affected %s", sqlText, rowAffected)
		}
		return nil
	}
}
