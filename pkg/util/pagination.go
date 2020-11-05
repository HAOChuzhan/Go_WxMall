package util

import (
	"wx-gin-master/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage 获取页面参数
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
