package app

import (
	"net/http"
	"wx-gin-master/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid 绑定并验证数据
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.BadRequest
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.BadRequest
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.BadRequest
	}

	return http.StatusOK, e.OK
}
