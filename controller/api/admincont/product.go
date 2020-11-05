package admincont

import (
	"log"
	"net/http"
	"regexp"
	"wx-gin-master/models/product"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary 添加商品信息
// @Produce json
// @param title body string true "商品名"
// @param cover body string true "封面图"
// @param carousel body string true "图集"
// @param brand body string true "商品品牌"
// @param series body string true "商品系列名"
// @param price body number true "零售价"
// @param selling_price body number true "商品售价"
// @param cost body number true "商品成本价"
// @param tags body string false "商品标签"
// @param sales body int false "销量"
// @param inventory body int false "库存"
// @param status body int false "状态,0-下架 1-上架"
// @param on_sale body int false "状态,0-未折扣 1-折扣中"
// @Success 201 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/project/ [post]
func CreateProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	productNew := product.Product{}
	productNew.Title = c.PostForm("title")
	productNew.Content = c.PostForm("content")
	productNew.ProductType = com.StrTo(c.PostForm("product_type")).MustInt()
	productNew.Img = c.PostForm("img")
	productNew.ShareNum = com.StrTo(c.PostForm("share_num")).MustInt()
	productNew.Price = com.StrTo(c.PostForm("price")).MustFloat64()
	productNew.IsHot = com.StrTo(c.PostForm("is_hot")).MustInt()
	/*
		projectNew.Brand = c.PostForm("brand")
		projectNew.Series = c.PostForm("series")

		projectNew.SellingPrice = com.StrTo(c.PostForm("selling_price")).MustFloat64()
		projectNew.Cost = com.StrTo(c.PostForm("cost")).MustFloat64()
		projectNew.Tags = c.PostForm("tags")

		projectNew.Inventory = com.StrTo(c.PostForm("inventory")).MustInt()
		projectNew.Status = com.StrTo(c.PostForm("status")).MustInt()
		projectNew.OnSale = com.StrTo(c.PostForm("on_sale")).MustInt()
	*/
	log.Println(productNew)

	valid.Required(productNew.Title, "Title").Message("请输入商品名")
	valid.MaxSize(productNew.Title, 80, "Title").Message("请输入商品名")

	valid.Required(productNew.Content, "content").Message("请输入商品内容")

	valid.Required(productNew.ProductType, "product_type").Message("请输入商品类型")

	valid.Required(productNew.Img, "img").Message("请出示封面图")
	valid.Match(productNew.Img, regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`), "img").Message("请出入封面图")

	//valid.Required(projectNew.Series, "Series").Message("请输入商品系列名")
	valid.Required(productNew.ShareNum, "share_num").Message("商品的分享次数")
	valid.Required(productNew.Price, "Price").Message("请输入商品零售价")
	valid.Required(productNew.IsHot, "is_hot").Message("请输入商品是否热门")

	//valid.Required(projectNew.Carousel, "Carousel").Message("请输入图集")
	//valid.Required(projectNew.SellingPrice, "SellingPrice").Message("请输入商品售价")
	//valid.Required(projectNew.Cost, "Cost").Message("请输入商品成本价")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	if !product.Create(&productNew) {
		appG.Response(http.StatusBadRequest, e.ErrorCreateProFain, nil)
		return
	}
	appG.Response(http.StatusCreated, e.Created, nil)
}

// @Summary 修改商品信息
// @Produce json
// @Param id path int true "要修改信息的id"
// @param title body string true "商品名"
// @param cover body string true "封面图"
// @param carousel body string true "图集"
// @param brand body string true "商品品牌"
// @param series body string true "商品系列名"
// @param price body number true "零售价"
// @param selling_price body number true "商品售价"
// @param cost body number true "商品成本价"
// @param tags body string true "商品标签"
// @param sales body int false "销量"
// @param inventory body int false "库存"
// @param status body int false "状态,0-下架 1-上架"
// @param on_sale body int false "状态,0-未折扣 1-折扣中"
// @Success 201 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/product/{oid} [put]

func UpdateProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("oid")).MustInt()
	productUp := product.Product{}
	productUp.ID = id
	productUp.Title = c.PostForm("title")
	productUp.Content = c.PostForm("content")
	productUp.ProductType = com.StrTo(c.PostForm("product_type")).MustInt()
	productUp.Img = c.PostForm("img")
	productUp.ShareNum = com.StrTo(c.PostForm("share_num")).MustInt()
	productUp.Price = com.StrTo(c.PostForm("price")).MustFloat64()
	productUp.IsHot = com.StrTo(c.PostForm("is_hot")).MustInt()
	valid.Required(productUp.ID, "Id").Message("请输入Id")
	valid.Min(productUp.ID, 0, "Id").Message("请输入有效id")
	valid.Required(productUp.Title, "Title").Message("请输入商品名")
	valid.MaxSize(productUp.Title, 80, "Title").Message("请输入商品名")
	valid.Required(productUp.Img, "img").Message("请出入封面图")
	valid.Match(productUp.Img, regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`), "Img").Message("请出入封面图")

	valid.Required(productUp.ShareNum, "share_num").Message("商品的分享次数")
	valid.Required(productUp.Price, "Price").Message("请输入商品零售价")
	valid.Required(productUp.IsHot, "is_hot").Message("请输入商品是否热门")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	if !product.UpDate(&productUp) {
		appG.Response(http.StatusInternalServerError, e.ErrorUpDateProFain, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}

// @Summary 删除商品
// @Produce  json
// @Param id path int true "要移除的商品id"
// @Success 204 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/product/{oid} [delete]
func DestroyProduct(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.PostForm("oid")).MustInt()

	if !product.Destroy(id) {
		appG.Response(http.StatusInternalServerError, e.ErrorUpDateProFain, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.NoContent, nil)
}
