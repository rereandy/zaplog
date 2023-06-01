package log

import (
	"fmt"
	"go.uber.org/zap"
)

// init 初始化
func init() {
	c := NewConfig()
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}

/*
封装相关方法后续通过包名直接调用
*/

// Debugf func
func Debugf(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// Infof info func
func Infof(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

// Warnf info func
func Warnf(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

// Errorf info func
func Errorf(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}
