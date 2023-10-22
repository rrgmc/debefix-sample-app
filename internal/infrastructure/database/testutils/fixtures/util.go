package fixtures

import (
	"fmt"
	"testing/fstest"

	"github.com/rrgmc/debefix"
)

func MergeData(source *debefix.Data, mergeData []string) (*debefix.Data, error) {
	newData, err := source.Clone()
	if err != nil {
		return nil, err
	}

	fs := fstest.MapFS{}
	for idx, data := range mergeData {
		fs[fmt.Sprintf("file_%02d.dbf.yaml", idx)] = &fstest.MapFile{Data: []byte(data)}
	}

	newData, err = debefix.Load(debefix.NewFSFileProvider(fs), debefix.WithLoadInitialData(newData))
	if err != nil {
		return nil, err
	}

	return newData, nil
}
