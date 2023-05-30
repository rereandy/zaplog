package log

import (
	"fmt"
)

func init() {
	c := NewConfig()
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}
