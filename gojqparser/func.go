package gojqparser

import "encoding/json"

func toNumber(v string) any {
	return normalizeNumber(json.Number(v))
}
