package dbtest

import (
	"errors"
	"path/filepath"
	"runtime"
)

func MigrationsPath() (string, error) {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("failed to get path")
	}
	return "file://" + filepath.ToSlash(filepath.Dir(path)) + "/../../../scripts/db-migrations", nil
}
