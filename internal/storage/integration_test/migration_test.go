//go:build dbmigrationtest

package integration_test

import (
	"testing"

	"github.com/RangelReale/debefix-sample-app/internal/testutils/dbtest"
	"github.com/stretchr/testify/require"
)

func TestDBMigration(t *testing.T) {
	err := dbtest.DBMigrationTest("debefix-sample-app")
	require.NoError(t, err)
}
