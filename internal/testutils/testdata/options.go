package testdata

type testDataOptions struct {
	tags []string
	sort string
}

type TestDataOption func(*testDataOptions)

func WithTags(tags []string) TestDataOption {
	return func(o *testDataOptions) {
		o.tags = tags
	}
}

func WithSort(sort string) TestDataOption {
	return func(o *testDataOptions) {
		o.sort = sort
	}
}

func parseOptions(options ...TestDataOption) testDataOptions {
	var optns testDataOptions
	for _, opt := range options {
		opt(&optns)
	}
	return optns
}
