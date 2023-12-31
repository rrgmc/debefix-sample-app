//go:build dbmigrationtest

package integration_test

import (
	"testing"

	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/dbtest"
	"gotest.tools/v3/assert"
)

func TestDBMigration(t *testing.T) {
	err := dbtest.DBMigrationTest("debefix-sample-app")
	assert.NilError(t, err)
}
