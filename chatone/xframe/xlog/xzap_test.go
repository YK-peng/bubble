package xlog

import (
	"encoding/json"
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func Test_initLogger(t *testing.T) {
	x := &Test{
		Name: "xiaoming",
		Age:  12,
	}
	data, err := json.Marshal(x)
	if err != nil {
		fmt.Println("marshal is failed,err: ", err)
	}

	//
	logger := NewLogger("./all.log", "debug")
	for i := 0; i < 6; i++ {
		logger.Info(fmt.Sprint("test log ", i), zap.Int("line", 47))
		logger.Debug(fmt.Sprint("debug log ", i), zap.ByteString("level", data))
		logger.Info(fmt.Sprint("Info log ", i), zap.String("level", `{"a":"4","b":"5"}`))
		logger.Warn(fmt.Sprint("Info log ", i), zap.String("level", `{"a":"7","b":"8"}`))
	}
}

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
