package seed

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"

	"github.com/gopaytech/go-commons/pkg/dir"
	"github.com/gopaytech/go-commons/pkg/zlog"
)

type Seeder func() error

func ProvideSeeder(db *sql.DB, path string) Seeder {
	return func() error {
		zlog.Debug("seed source dir: %s", path)
		files, err := dir.ListFiles(path)
		if err != nil {
			return err
		}
		for _, file := range files {
			logField := zlog.LogFields{
				"Path": path,
			}

			fullPath := filepath.Join(path, file)

			logField["SeedFile"] = fullPath
			zlog.DebugF(logField, "read seed file")

			openedFile, err := ioutil.ReadFile(fullPath)
			if err != nil {
				return err
			}

			sqlText := string(openedFile)
			zlog.DebugF(logField, "execute seed file")

			logField["SqlText"] = sqlText
			result, err := db.Exec(sqlText)
			if err != nil {
				zlog.ErrorF(logField, err, "execution failed")
				return err
			}
			rowAffected, _ := result.RowsAffected()
			zlog.DebugF(logField, "execution finished, row affected %d", rowAffected)
		}
		return nil
	}
}
