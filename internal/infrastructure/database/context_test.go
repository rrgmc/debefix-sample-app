package database

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"
)

func testContext(t *testing.T, testFn func(db *gorm.DB, mock sqlmock.Sqlmock)) {
	conn, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer conn.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 conn,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	assert.NilError(t, err)

	testFn(db, mock)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestUnitOfWork(t *testing.T) {
	testContext(t, func(db *gorm.DB, mock sqlmock.Sqlmock) {
		rctx := NewContext(db)

		mock.ExpectBegin()
		mock.ExpectCommit()

		ctx := context.Background()

		uow, err := rctx.StartUnitOfWork(ctx, nil)
		assert.NilError(t, err)

		err = uow.Commit(ctx)
		assert.NilError(t, err)
	})
}
