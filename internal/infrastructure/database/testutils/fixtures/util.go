package fixtures

import (
	"github.com/rrgmc/debefix"
)

func MergeData(source *debefix.Data, mergeData []string) (*debefix.Data, error) {
	newData, err := source.Clone()
	if err != nil {
		return nil, err
	}

	newData, err = debefix.Load(debefix.NewStringFileProvider(mergeData),
		debefix.WithLoadInitialData(newData),
		debefix.WithLoadRowsSetIgnoreTags(true))
	if err != nil {
		return nil, err
	}

	return newData, nil
}
