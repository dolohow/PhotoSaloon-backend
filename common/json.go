package common

import (
	"encoding/json"
	"fmt"
)

// JSONMsg returns simple json with one field "msg" created from format
// specifier.
func JSONMsg(format string, a ...interface{}) string {
	msg, _ := json.Marshal(map[string]string{
		"msg": fmt.Sprintf(format, a...),
	})

	return string(msg)
}
