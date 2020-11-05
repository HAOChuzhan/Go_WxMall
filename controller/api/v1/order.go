package v1

import (
	"net/http"
	"strconv"

	"wx-gin-master/models/order"
	"wx-gin-master/models/product"
	"wx-gin-master/models/user"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 新建订单
// @Produce json
// @param cartIds body int true "购物车IDs"
// @param addressId body int false "地址ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/order [post]
func CreateOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.PostForm("cartIds"), "cartIds").Message("cartId 必须")
	cartIds := c.PostFormArray("cartIds")
	for _, cartId := range cartIds {
		valid.Numeric(cartId, "cartId").Message("cartId 必须是有效数值")
	}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	//var address user.Address

	var total float64
	/*if addressId := com.StrTo(c.PostForm("addressId")).MustInt(); addressId != 0 {
		address = user.SelectAddress(addressId)
	} else {
		address = user.SelectDefaultAddress(userId)
	}*/

	carts := user.QueryCarts(cartIds)
	userOrder := user.QueryUserById(userId)

	// *** 可以优化
	for _, c := range carts {
		total += float64(c.Number) * product.Show(c.PId).Price
	}

	orderDate := order.Order{
		UserId: userId,
		//ProvinceName: address.ProvinceName,
		//CityName:     address.CityName,
		//CountyName:   address.CountyName,
		//DetailInfo:   address.DetailInfo,
		//PostalCode:   address.PostalCode,
		UserName:  userOrder.Nickname,
		TelNumber: userOrder.Mobile,
		Total:     total,
		SumPay:    total,
	}
	if !order.Created(&orderDate) {
		appG.Response(http.StatusInternalServerError, e.ErrorCreatedOrderFail, nil)
		return
	}
	appG.Response(http.StatusCreated, e.Created, nil)
}

// @Summary 销毁订单
// @Produce  json
// @Param id path int true "要移除的id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/order/{id} [delete]
func DestroyOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.Param("oid"), "oid").Message("oid 必须")
	valid.Numeric(c.Param("oid"), "oid").Message("oid 必须是有效数字")

	oid := com.StrTo(c.Param("oid")).MustInt()
	orderInfo := order.QueryOrderById(oid)
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	if orderInfo.UserId != userId {
		appG.Response(http.StatusUnauthorized, e.Unauthorized, nil)
		return
	}
	if !order.Destroy(oid) {
		appG.Response(http.StatusUnauthorized, e.Unauthorized, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}

// @Summary 获取所有订单信息
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/user/order [get]
func ListOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	data := make(map[string]interface{})
	data["orders"] = order.QueryOrderByUserId(userId)
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 支付订单
// @Produce json
// @Param oid path int true "要支付的订单id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/order/:id/pay [post]
func PayOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.Param("oid"), "oid").Message("oid 必须")
	valid.Numeric(c.Param("oid"), "oid").Message("oid 必须是有效数字")

	oid := com.StrTo(c.Param("oid")).MustInt()
	orderInfo := order.QueryOrderById(oid)
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	if orderInfo.UserId != userId {
		appG.Response(http.StatusUnauthorized, e.Unauthorized, nil)
		return
	}
	if !order.Settlement(userId, oid) {
		appG.Response(http.StatusForbidden, e.ErrorSettlementFail, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}

// @Summary 查看订单包含的商品
// @Produce json
// @Param oid path int true "要查询的订单id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/user/order/{id} [get]
func ViewOrderDetailsOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.Param("id"), "id").Message("id(Order Id) 必须")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	oid := com.StrTo(c.Param("id")).MustInt()
	if !order.IsOwner(userId, oid) {
		appG.Response(http.StatusForbidden, e.ErrorIsOwnerFail, nil)
		return
	}

	carts := user.QueryCartByOid(oid)
	data := make(map[string]interface{})
	for i, cart := range carts {
		data[strconv.Itoa(i)] = product.Show(cart.PId)
	}
	appG.Response(http.StatusOK, e.OK, data)
}
