package utils

import "slices"

func EnsureSliceContains(s []string, shouldContain []string) []string {
	ret := slices.Clone(s)
	for _, c := range shouldContain {
		if !slices.Contains(ret, c) {
			ret = append(ret, c)
		}
	}
	return ret
}
