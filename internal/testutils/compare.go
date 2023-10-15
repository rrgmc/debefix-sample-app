package testutils

import (
	"fmt"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Equal(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if diff := cmp.Diff(expected, actual, defaultCmpOptions...); diff != "" {
		return assert.Fail(t, fmt.Sprintf("Not equal: \n%s", diff), msgAndArgs...)
	}
	return true
}

func RequireEqual(t require.TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if Equal(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// default go-cmp options
//
//nolint:unused
var defaultCmpOptions = []cmp.Option{
	cmp.Comparer(func(x, y time.Time) bool { return true }), // don't compare time fields
}

type tHelper interface {
	Helper()
}
