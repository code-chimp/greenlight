package models

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("%d mins", r)
	quoted := strconv.Quote(json)
	return []byte(quoted), nil
}
