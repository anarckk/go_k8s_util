package util

import (
	"bytes"
	"fmt"
)

func ComposeMap(matchLabelsMap map[string]string) string {
	size := len(matchLabelsMap)
	i := 0
	var bt bytes.Buffer
	for k, v := range matchLabelsMap {
		bt.WriteString(fmt.Sprintf("%s=%s", k, v))
		if i < size-1 {
			bt.WriteString(",")
		}
		i++
	}
	return bt.String()
}
