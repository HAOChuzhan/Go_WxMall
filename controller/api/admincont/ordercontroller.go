package admincont

/*
还未完善各项功能

*/
import (
	"net/http"
	"time"
	"wx-gin-master/models/order"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 修改订单快递信息
// @Produce json
// @Param id path int true "要修改的订单id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/order/{id}/express [post]
func UpdateOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	valid.Required(c.Param("id"), "oid").Message("oid")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	oid := com.StrTo(c.Param("id")).MustInt()

	orderUp := order.QueryOrderById(oid)

	expressTime := time.Now().Unix()
	orderData := order.Order{

		ExpressTime: expressTime,
	}
	if orderUp.UserId != userId {
		appG.Response(http.StatusUnauthorized, e.Unauthorized, nil)
		return
	}
	if !order.Update(oid, &orderData) {
		appG.Response(http.StatusForbidden, e.ErrorUpdateOrderFain, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}

/*
func GetOderList(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	//userId := c.MustGet("AuthData").(*util.Claims).User.ID
	valid.Required(c.Param("id"), "oid").Message("oid")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	oid := com.StrTo(c.Param("id")).MustInt()
	//orderUp := order.QueryOrderById(oid)

}*/
