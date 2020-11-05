package admincont

import (
	"net/http"

	"wx-gin-master/models/auth"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Summary 管理员登录
// @Produce json
// @param username body string true "用户名"
// @param password body string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /adminUserLogin [post]
func LoginAdmin(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.Param("username")
	password := c.Param("password")
	valid.Required(username, "username").Message("请输入用户名")
	valid.Required(password, "password").Message("请输入密码")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	id := auth.CheckAdmin(username, password)
	if id <= 0 {
		appG.Response(http.StatusUnauthorized, e.ErrorCheckAdminFail, valid.Errors)
		return
	}
	token, err := util.GenerateAdmin(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthToken, nil)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 获取AuthData
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /admin/info [get]
func ShowAdmin(c *gin.Context) {
	appG := app.Gin{C: c}
	data := c.MustGet("AuthData").(*util.AdminClaims)
	appG.Response(http.StatusOK, e.OK, data)
}
