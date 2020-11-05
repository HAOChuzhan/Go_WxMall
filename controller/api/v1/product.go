package v1

import (
	"net/http"

	"wx-gin-master/models/product"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/setting"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 商品列表
// @Produce json
// @param brand query string false "brand品牌"
// @param series query string false "series系列"
// @param page query string false "page页数"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/product/index [get]
func IndexProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	where := make(map[string]interface{})
	data := make(map[string]interface{})

	if title := c.Query("title"); title != "" {
		where["title"] = title
	}
	if producttype := c.Query("product_type"); producttype != "" {
		where["producttype"] = producttype
	}

	data["lists"], data["total"] = product.List(util.GetPage(c), setting.AppSetting.PageSize, where)
	if _, ok := data["lists"]; !ok {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 搜索
// @Produce json
// @param tag query string false "根据标签搜索"
// @param title query string false "根据商品名搜索"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/product/search [get]

func SearchProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	/*
		if tag := c.Query("tag"); tag != "" {
			data["lists"], data["total"] = product.SearchInTag(util.GetPage(c), setting.AppSetting.PageSize, tag)

		} else
	*/
	if title := c.Query("title"); title != "" {
		data["lists"], data["total"] = product.SearchInTitle(util.GetPage(c), setting.AppSetting.PageSize, title)
	}

	if _, ok := data["lists"]; !ok {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, data)
}

// @Summary 根据id获取商品信息
// @Produce json
// @param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/product/info/:id [get]
func ShowProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}

	productInfo := product.Show(id)
	appG.Response(http.StatusOK, e.OK, productInfo)
}
