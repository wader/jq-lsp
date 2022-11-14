package gojqparser

import "encoding/json"

func toNumber(v string) interface{} {
	return normalizeNumber(json.Number(v))
}
