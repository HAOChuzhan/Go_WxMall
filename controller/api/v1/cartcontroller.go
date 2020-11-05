package v1

import (
	"net/http"
	"wx-gin-master/models/product"
	"wx-gin-master/models/user"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 初始化购物车列表
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/user/shoppingCart [get]
func IndexCart(c *gin.Context) {
	appG := app.Gin{C: c}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	data := make(map[string]interface{})
	data["products"] = user.CartsProducts(userId)
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 添加到购物车
// @Produce json
// @param pId body int true "商品ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/shoppingCart [post]
func CreateCart(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.PostForm("pId"), "pId").Message("pId 必须")
	valid.Numeric(c.PostForm("pId"), "pId").Message("pId 必须有效")

	pId := com.StrTo(c.PostForm("pId")).MustInt()
	//?是否在这刻减少库存
	product.UpdateInventory(pId, -1)
	if !user.PutInCart(userId, pId) {
		appG.Response(http.StatusInternalServerError, e.ErrorPutInCartFail, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}

// @Summary 从购物车移除
// @Produce  json
// @Param id path int true "要移除的购物车id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/shoppingCart/{id} [delete]
func DestroyCart(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.Param("id"), "cId").Message("cart Id 参数值必须")
	valid.Numeric(c.Param("id"), "cId").Message("cart Id 必须是有效数值")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	cId := com.StrTo(c.Param("id")).MustInt()
	cart := user.QueryCartById(cId)
	if cart.UserId != userId {
		appG.Response(http.StatusUnauthorized, e.Unauthorized, nil)
		return
	}
	if !user.DropFromCart(cId) {
		appG.Response(http.StatusInternalServerError, e.ErrorDropFromCartFail, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}

// @Summary 减少购物车商品数量
// @Produce  json
// @Param id path int true "要减少的购物车id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/shoppingCart/{id}/number [delete]
func DecreaseCart(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.Param("id"), "cId").Message("cart Id 参数值必须")
	valid.Numeric(c.Param("id"), "cId").Message("cart Id 必须是有效数值")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	cId := com.StrTo(c.Param("id")).MustInt()
	cart := user.QueryCartById(cId)
	if cart.UserId != userId {
		appG.Response(http.StatusUnauthorized, e.Unauthorized, nil)
		return
	}
	if !user.NumberDecrease(cId) {
		appG.Response(http.StatusInternalServerError, e.ErrorDropFromCartFail, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}
