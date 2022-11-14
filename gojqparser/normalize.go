package gojqparser

import (
	"encoding/json"
	"math"
	"math/big"
	"strings"
)

func normalizeNumber(v json.Number) interface{} {
	if i, err := v.Int64(); err == nil && math.MinInt <= i && i <= math.MaxInt {
		return int(i)
	}
	if strings.ContainsAny(v.String(), ".eE") {
		if f, err := v.Float64(); err == nil {
			return f
		}
	}
	if bi, ok := new(big.Int).SetString(v.String(), 0); ok {
		return bi
	}
	if strings.HasPrefix(v.String(), "-") {
		return math.Inf(-1)
	}
	return math.Inf(1)
}
