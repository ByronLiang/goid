package utils

import (
	"strconv"
	"strings"
)

func SplitParseInt64(str string) (ids []int64) {
	for _, sid := range strings.Split(str, ",") {
		id, e := strconv.ParseInt(sid, 10, 64)
		if e != nil {
			continue
		}
		ids = append(ids, id)
	}
	return
}
