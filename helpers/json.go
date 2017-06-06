package helpers

import "encoding/json"

// IsJSONString function returns true if the given string is in JSON format.
func IsJSONString(s string) bool {
	var jsonStr string
	return json.Unmarshal([]byte(s), &jsonStr) == nil
}
