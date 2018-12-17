package message

import (
	"fmt"
	"sort"
)

func mapToString(x map[string]interface{}) string {
	var keys []string
	for k := range x {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var pairs []string
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%s:%v", k, x[k]))
	}
	return fmt.Sprintf("map%v", pairs)
}

// JSONError is returned if JSON is malformed because of a bad key.
type JSONError struct {
	Data   map[string]interface{}
	BadKey string
	Reason string
}

func (e JSONError) Error() string {
	return fmt.Sprintf(
		"key \"%s\" is bad in object %s because it is %s",
		e.BadKey, mapToString(e.Data), e.Reason)
}
