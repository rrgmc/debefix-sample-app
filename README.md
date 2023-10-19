# debefix-sample-app

This is a real-world sample of the [debefix](https://github.com/rrgmc/debefix) Go library.

debefix is a Go library to seed database data and/or create fixtures for DB tests.

It is a simple blog REST microservice, based on a layered architecture, with endpoint, service, and 
storage separations.

## Layout

- `cmd/server`: service entrypoint.
- `pkg/app`: app object used by `cmd/server` to do initialization and execution.
- `scripts/db-migrations`: database migrations using `golang-migrate`.
- `scripts/db-seed`: database seeding using `debefix` fixtures.
- `scripts/local-deps`: `docker-compose` scripts to run the local dependencies (PostgreSQL).
- `internal/testutils/fixtures`: fixtures for tests and seeding using `debefix`.
- `internal/testutils/dbtest`: `testcontainers-go` method to start a test PostgreSQL server, run the migrations, and 
  apply the fixtures. The database is created in a `tmpfs` to be faster and avoid disk usage.
- `internal/testutils/testdata`: extract test data objects from fixtures, with a simple query filter to mimic some SQL operations.
- `internal/storage`: storage layer, the only part that has database access.
- `internal/storage/integration_test`: storage layer tests that uses a real database.
- `internal/entity`: entities used by `storage` and `service` layers.

## Task rules

[go-task](https://github.com/go-task/task) is used for command line task running.

- `task gen`: generate mocks.
- `task local-deps-setup`: setup the local dependencies, creating, migrating and seeding the database.
- `task local-deps`: start a previously-setup local dependencies.
- `task db-migration-up`: execute pending database migrations.
- `task test`: execute only the non-database-based tests.
- `task test-db`: execute only the database-based tests.
- `task test-db-migrations`: run the database migrations (up and down) in a temporary docker container.

## Fixtures

The `debefix` fixtures are in the `internal/testutils/fixtures` folder. The directories are numbered because
`debefix` loads them in order, and the order matters. Each folder is assigned a tag by removing the number and
dash prefixes, and for inner directories concatenating all folders with a dot.

- `fixtures/01-base` [tag: **base**]: Always applied. It contains initialization of static data, in this case the countries list.
- `fixtures/05-seed` [tag: **seed**]: Only used by the database seeding using `task local-deps-setup` or `task db-seed`. 
  Meant for local developer database, not used in tests.
- `fixtures/50-tests/01-base` [tag: **tests.base**]: Always applied on tests. This data will be available for all tests.

For test-specific fixtures, create new folders in `fixtures/50-tests/`, and call 
`dbtest.DBForTest(..., dbtest.WithDBForTestFixturesTags("tests.other-test-folder"))` on the test. `DBForTest` returns
the inserted data, including any generared fields like primary keys, which can be used by `internal/testutils/testdata`
to generate expected data from the same source for tests.

## License

MIT

### Author

Rangel Reale (rangelreale@gmail.com)
