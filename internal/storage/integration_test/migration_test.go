//go:build dbmigrationtest

package integration_test

import (
	"testing"

	"github.com/RangelReale/debefix-sample-app/internal/testutils/dbtest"
	"gotest.tools/v3/assert"
)

func TestDBMigration(t *testing.T) {
	err := dbtest.DBMigrationTest("debefix-sample-app")
	assert.NilError(t, err)
}
