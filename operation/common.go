package operation

import (
	"sync"
)

const (
	RENAME  = "rename"
	INSERT  = "insert"
	MOVE    = "move"
	REPLACE = "replace"
	HEAD    = "head"
	TAIL    = "tail"
	DELETE  = "delete"
)

var TempSlicePool = sync.Pool{
	New: func() interface{} {
		return []interface{}{}
	},
}

func makeSlice(l []interface{}, index int) []interface{} {
	temp := TempSlicePool.Get().([]interface{})
	TempSlicePool.Put([]interface{}{})
	for n, v := range l {
		if n == index {
			continue
		} else {
			temp = append(temp, v)
		}
	}
	return temp
}
