package testdata

import (
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix/filter"
)

type testDataOptions struct {
	filterDataOptions []filter.FilterDataOption
	mergeData         *debefix.Data
	resolveTags       []string
	sort              string
}

type TestDataOption func(*testDataOptions)

func WithMergeData(data *debefix.Data) TestDataOption {
	return func(o *testDataOptions) {
		o.mergeData = data
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

// WithResolveTags sets the tags for the data resolver.
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
	return optns
}
