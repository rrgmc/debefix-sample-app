package testdata

type testDataOptions struct {
	tags []string
}

type TestDataOption func(*testDataOptions)

func WithTags(tags []string) TestDataOption {
	return func(o *testDataOptions) {
		o.tags = tags
	}
}

func parseOptions(options ...TestDataOption) testDataOptions {
	var optns testDataOptions
	for _, opt := range options {
		opt(&optns)
	}
	return optns
}
