package v1

import (
	"net/http"
	"regexp"
	"wx-gin-master/models/auth"
	"wx-gin-master/models/user"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/setting"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 注册用户
// @Produce json
// @param mobile body string true "手机号"
// @param password body string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/register [post]
func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	valid := validation.Validation{}
	valid.Required(mobile, "mobile").Message("请输入手机号")
	valid.MaxSize(mobile, 11, "mobile").Message("请输入有效电话")
	valid.Phone(mobile, "mobile").Message("请输入有效电话")
	valid.Required(password, "password").Message("密码不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	password = util.EncodeMD5(password)
	newUser := user.User{
		Mobile:   mobile,
		Password: password,
	}
	// 应判断手机号是否已经注册
	if !user.CreateByPassword(&newUser) {
		appG.Response(http.StatusInternalServerError, e.ErrorCreateUserFail, nil)
		return
	}
	id := auth.CheckAndReturnId(mobile, password)
	if id <= 0 {
		appG.Response(http.StatusInternalServerError, e.ErrorCreateUserFail, nil)
		return
	}
	token, err := util.GenerateToken(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthToken, nil)
		return
	}
	data["token"] = token
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 获取个人信息
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/index [get]
func ShowUser(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	data["user"] = user.GetInfo(userId)
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 用户登录
// @Produce json
// @param mobile body string true "mobile"
// @param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/login [post]
func LoginUser(c *gin.Context) {
	appG := app.Gin{C: c}

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	valid := validation.Validation{}
	valid.Required(mobile, "mobile").Message("请输入手机号")
	valid.MaxSize(mobile, 11, "mobile").Message("请输入有效电话")
	valid.Phone(mobile, "mobile").Message("请输入有效电话")
	valid.Required(password, "password").Message("密码不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	password = util.EncodeMD5(password)
	id := auth.CheckAndReturnId(mobile, password)
	if id <= 0 {
		appG.Response(http.StatusInternalServerError, e.ErrorLoginUserFail, nil)
		return
	}
	token, err := util.GenerateToken(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthToken, nil)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 修改用户信息
// @Produce  json
// @Param nickname body string false "昵称"
// @Param avatar body string false "头像地址"
// @Param sex body int false "性别 1-男 2-女"
// @Param introduction body string false "简介"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/edit [put]
func UpdateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	nickname := c.PostForm("nickname")
	avatar := c.PostForm("avatar")
	sex := c.PostForm("sex")
	introduction := c.PostForm("introduction")
	editedData := make(map[string]interface{})
	if nickname != "" {
		err := valid.MaxSize(nickname, 10, "nickname").Message("限定10个字符").Error
		if err == nil {
			editedData["nickname"] = nickname
		}
	}
	if avatar != "" {
		reg := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
		err := valid.Match(avatar, reg, "avatar").Message("请上传正确图片").Error
		if err == nil {
			editedData["avatar"] = avatar
		}
	}
	if sex != "" {
		err := valid.Numeric(sex, "sex").Message("请传入有效数据").Error
		sexInt, _ := com.StrTo(sex).Int()
		errMin := valid.Min(sexInt, 0, "sex").Message("性别传值不正确").Error
		errMax := valid.Max(sexInt, 2, "sex").Message("性别传值不正确").Error
		if errMin == nil && errMax == nil && err != nil {
			editedData["sex"] = sexInt
		}
	}
	if introduction != "" {
		err := valid.MaxSize(introduction, 80, "introduction").Message("限定80个字符").Error
		if err == nil {
			editedData["introduction"] = introduction
		}
	}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	updateUser := user.GetInfo(userId)
	// 更新数据
	if !user.Update(&updateUser, editedData) {
		appG.Response(http.StatusInternalServerError, e.ErrorUpdateUserFail, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, nil)
}

// @Summary 获取收藏列表
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/product/favorite [get]
func FavoritesList(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	data["lists"], data["total"] = user.ShowFavorites(util.GetPage(c), setting.AppSetting.PageSize, userId)
	if _, ok := data["lists"]; !ok {
		appG.Response(http.StatusInternalServerError, e.ErrorGetFavoritesListFail, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 添加到收藏夹
// @Produce json
// @param pId query int true "商品ID"
// @Success 204 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/user/favorite [post]
func FavoritesCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	valid.Required(c.PostForm("pId"), "pId").Message("pId 必须")
	valid.Numeric(c.PostForm("pId"), "pId").Message("pId 必须是有效数值")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}
	pId := com.StrTo(c.PostForm("pId")).MustInt()
	if !user.AddFavorite(userId, pId) {
		appG.Response(http.StatusInternalServerError, e.ErrorAddFavoriteFail, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}

// @Summary 从收藏夹移除
// @Produce  json
// @Param id path int true "要移除的商品ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/favorite/{id} [delete]
func FavoritesDestroy(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	valid.Required(c.Param("id"), "pId").Message("pId 参数值必须")
	valid.Numeric(c.Param("id"), "pId").Message("pId 必须是有效数值")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}
	pId := com.StrTo(c.Param("id")).MustInt()
	if !user.DestroyFavorite(userId, pId) {
		appG.Response(http.StatusInternalServerError, e.ErrorDestroyFavoriteFail, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}
