package fixtures

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix"
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
		debefix.NewDirectoryFileProvider(curDir, debefix.WithDirectoryAsTag()),
		debefix.WithLoadTaggedValueParser(
			debefix.ValueParserUUID(),
		))
	if err != nil {
		panic(fmt.Sprintf("error loading test fixtures: %s", err))
	}
}

func ResolveFixtures(options ...ResolveFixtureOption) (*debefix.Data, error) {
	optns := &resolveFixturesOptions{
		tags: []string{"01-base"},
	}
	for _, opt := range options {
		opt(optns)
	}

	sourceData := fixtures
	if optns.resolvedData != nil {
		sourceData = optns.resolvedData
		optns.tags = nil // don't filter tags if already resolved data
	}

	data, err := debefix.Resolve(sourceData, resolveCallback, debefix.WithResolveTags(optns.tags))
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

func WithMergeData(data *debefix.Data) ResolveFixtureOption {
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
		o.tags = nil
		if !slices.Contains(tags, "01-base") {
			o.tags = append(o.tags, "01-base")
		}
		o.tags = append(o.tags, tags...)
	}
}

func WithOutput(output bool) ResolveFixtureOption {
	return func(o *resolveFixturesOptions) {
		o.output = output
	}
}

type resolveFixturesOptions struct {
	mergeData    *debefix.Data
	resolvedData *debefix.Data
	tags         []string
	output       bool
}

func resolveCallback(ctx debefix.ResolveContext, fields map[string]any) error {
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
