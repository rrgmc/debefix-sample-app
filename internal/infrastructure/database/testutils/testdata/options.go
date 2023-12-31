package testdata

import (
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"github.com/rrgmc/debefix/filter"
)

type testDataOptions struct {
	filterDataOptions []filter.FilterDataOption
	filterRefIDs      []string
	resolvedData      *debefix.Data
	resolveTags       []string
	sort              string
}

type TestDataOption func(*testDataOptions)

// WithResolvedData sets the resolved data. If set, tag filtering by "WithResolveTags" will be disabled.
func WithResolvedData(data *debefix.Data) TestDataOption {
	return func(o *testDataOptions) {
		o.resolvedData = data
	}
}

// WithFilterAll include all records by default, depending on other filters if they exist.
// All requested filters must return true to select the row.
func WithFilterAll(filterAll bool) TestDataOption {
	return func(o *testDataOptions) {
		o.filterDataOptions = append(o.filterDataOptions, filter.WithFilterAll(filterAll))
	}
}

// WithFilterRefIDs filters by refID.
// All requested filters must return true to select the row.
func WithFilterRefIDs(refIDs []string) TestDataOption {
	return func(o *testDataOptions) {
		o.filterDataOptions = append(o.filterDataOptions, filter.WithFilterRefIDs(refIDs))
		o.filterRefIDs = refIDs
	}
}

// WithFilterFields filters fields values.
// All requested filters must return true to select the row.
func WithFilterFields(fields map[string]any) TestDataOption {
	return func(o *testDataOptions) {
		o.filterDataOptions = append(o.filterDataOptions, filter.WithFilterFields(fields))
	}
}

// WithFilterRow filters using a callback.
// All requested filters must return true to select the row.
func WithFilterRow(filterRow func(row debefix.Row) (bool, error)) TestDataOption {
	return func(o *testDataOptions) {
		o.filterDataOptions = append(o.filterDataOptions, filter.WithFilterRow(filterRow))
	}
}

// WithOffsetLimit filters the returning array from the offset, with limit amount of records.
func WithOffsetLimit(offset int, limit int) TestDataOption {
	return func(o *testDataOptions) {
		o.filterDataOptions = append(o.filterDataOptions, filter.WithOffsetLimit(offset, limit))
	}
}

// WithResolveTags sets the tags for the data resolver.
// "base" and "tests.base" are always included automatically.
func WithResolveTags(tags []string) TestDataOption {
	return func(o *testDataOptions) {
		o.resolveTags = tags
	}
}

// WithSort sets the output sort name.
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
	optns.resolveTags = utils.EnsureSliceContains(optns.resolveTags, []string{"base", "tests.base"})
	return optns
}
