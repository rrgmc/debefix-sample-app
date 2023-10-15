//go:build dbmigrationtest

package integration_test

import (
	"testing"

	"github.com/RangelReale/debefix-sample-app/internal/testutils/dbtest"
	"github.com/stretchr/testify/require"
)

func TestDBMigration(t *testing.T) {
	migrationsPath, err := dbtest.MigrationsPath()
	require.NoError(t, err)

	err = dbtest.DBMigrationTest("debefix-sample-app", migrationsPath)
	require.NoError(t, err)
}
