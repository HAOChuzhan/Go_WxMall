package app

import (
	"wx-gin-master/pkg/logging"

	"github.com/astaxie/beego/validation"
)

// MarkErrors 记录错误日志
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
