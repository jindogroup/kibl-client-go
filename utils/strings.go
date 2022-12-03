package utils

import "strconv"

var String strUtils

type strUtils struct {
}

func (strUtils) ToInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

func (u strUtils) ToInt64(str string) int64 {
	if n, e := strconv.ParseInt(str, 10, 64); e == nil {
		return n
	}
	return 0
}

func (u strUtils) ToInt64Slice(strs []string) []int64 {
	if len(strs) == 0 {
		return []int64{}
	}
	out := make([]int64, len(strs))
	for i := range strs {
		out[i] = u.ToInt64(strs[i])
	}
	return out
}

func (u strUtils) FromInt64Slice(ints []int64) []string {
	if len(ints) == 0 {
		return []string{}
	}
	out := make([]string, len(ints))
	for i := range ints {
		out[i] = strconv.Itoa(int(ints[i]))
	}
	return out
}

func (u strUtils) ToFloat64(str string) float64 {
	if n, e := strconv.ParseFloat(str, 64); e == nil {
		return n
	}
	return 0
}