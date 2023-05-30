package log

import (
	"fmt"
)

// init 初始化
func init() {
	c := NewConfig()
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}
