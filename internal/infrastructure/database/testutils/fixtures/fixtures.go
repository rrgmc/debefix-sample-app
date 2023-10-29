package fixtures

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"github.com/rrgmc/debefix/value"
)

var (
	fixtures *debefix.Data
)

func init() {
	curDir, err := currentSourceDirectory()
	if err != nil {
		panic(err)
	}

	fixtures, err = debefix.Load(
		debefix.NewDirectoryFileProvider(curDir,
			debefix.WithDirectoryTagFunc(debefix.StripNumberPunctuationPrefixDirectoryTagFunc)),
		debefix.WithLoadValueParser(
			value.ValueUUID{},
		))
	if err != nil {
		panic(fmt.Sprintf("error loading test fixtures: %s", err))
	}
}

func ResolveFixtures(options ...ResolveFixtureOption) (*debefix.Data, error) {
	var optns resolveFixturesOptions
	for _, opt := range options {
		opt(&optns)
	}
	optns.tags = utils.EnsureSliceContains(optns.tags, []string{"base"})

	sourceData := fixtures
	if optns.resolvedData != nil {
		sourceData = optns.resolvedData
		optns.tags = nil // don't filter tags if already resolved data
	}

	data, err := debefix.Resolve(sourceData, resolveFixturesCallback, debefix.WithResolveTags(optns.tags))
	if err != nil {
		return nil, fmt.Errorf("error resolving fixtures with tags '%s': %w", strings.Join(optns.tags, ", "), err)
	}
	return data, nil
}

func MustResolveFixtures(options ...ResolveFixtureOption) *debefix.Data {
	data, err := ResolveFixtures(options...)
	if err != nil {
		panic(err)
	}
	return data
}

type ResolveFixtureOption func(*resolveFixturesOptions)

func WithMergeData(data []string) ResolveFixtureOption {
	return func(o *resolveFixturesOptions) {
		o.mergeData = data
	}
}

func WithResolvedData(data *debefix.Data) ResolveFixtureOption {
	return func(o *resolveFixturesOptions) {
		o.resolvedData = data
	}
}

func WithTags(tags []string) ResolveFixtureOption {
	return func(o *resolveFixturesOptions) {
		o.tags = tags
	}
}

func WithOutput(output bool) ResolveFixtureOption {
	return func(o *resolveFixturesOptions) {
		o.output = output
	}
}

type resolveFixturesOptions struct {
	resolvedData *debefix.Data
	mergeData    []string
	tags         []string
	output       bool
}

// resolveFixturesCallback is used for in-memory fixture resolving, so we don't have a database to generate values.
func resolveFixturesCallback(ctx debefix.ResolveContext, fields map[string]any) error {
	for fn, fv := range fields {
		if fresolve, ok := fv.(debefix.ResolveValue); ok {
			switch fresolve.(type) {
			case *debefix.ResolveGenerate:
				// we know that all our generated fields are UUID
				ctx.ResolveField(fn, uuid.New())
			}
		}
	}
	return nil
}

func currentSourceDirectory() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current source filename")
	}
	return filepath.Dir(filename), nil
}
