package main

import "fmt"

func serialize(
	name string, timestamp int64,
	field string, value float64) string {
	return fmt.Sprintf(
		"{\"Name\":\"%s\",\"Timestamp\":%d,\"%s\":%f}",
		name, timestamp, field, value)
}
