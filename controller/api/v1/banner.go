package v1

import (
	"net/http"
	"wx-gin-master/models/banner"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func CreateBanner(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{} //数据验证

	name := c.PostForm("name")
	imageurl := c.PostForm("imageurl")
	relateid := com.StrTo(c.PostForm("relateid")).MustInt()
	jumptype := com.StrTo(c.PostForm("jumptype")).MustInt()

	valid.Required(name, "name").Message("name is necessary!")
	valid.Required(imageurl, "imageurl").Message("imageurl is necessary")
	valid.Required(relateid, "relateid").Message("relateid 必须是有效数值")
	valid.Required(jumptype, "jumptype").Message("jumptype 必须是有效数值")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)

		return
	}
	addedBanner := banner.Banner{
		Name:     name,
		ImageUrl: imageurl,
		RelateId: relateid,
		JumpType: jumptype,
	}

	//banner_ID在Creat之后会按主键自增
	if !banner.CreateBanner(&addedBanner) {
		appG.Response(http.StatusInternalServerError, e.ErrorCreatedOrderFail, nil)
		return
	}
	appG.Response(http.StatusCreated, e.Created, addedBanner)
}

//更新轮播图
func UpdateBanner(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{} //数据验证

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	name := c.PostForm("name")
	imageurl := c.PostForm("imageurl")
	//relateid := c.PostForm("relateid")
	//jumptype := c.PostForm("jumptype")
	relateid := com.StrTo(c.PostForm("relateid")).MustInt()
	jumptype := com.StrTo(c.PostForm("jumptype")).MustInt()

	valid.Required(name, "name").Message("name 必须")
	valid.Required(imageurl, "imageurl").Message("imageurl 必须")
	valid.Required(relateid, "relateid").Message("relateid 必须有效数字")
	valid.Required(jumptype, "jumptype").Message("jumptype 必须有效数字")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}

	editedbanner := banner.Banner{
		Name:     name,
		ImageUrl: imageurl,
		RelateId: relateid,
		JumpType: jumptype,
	}
	editedbanner.ID = id
	editedbanner.Del = 0

	if !banner.UpdateBanner(&editedbanner) {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, nil)
}

func DeleteBanner(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}

	if !banner.DeleteBanner(id) {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, nil)
}

// @Summary 获取轮播图
// @Produce json
// @param num path int true "num获取轮播图数量"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/banner/indexbanner [post]
func IndexBanner(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{} //数据验证

	num := com.StrTo(c.PostForm("num")).MustInt()
	valid.Min(num, 1, "num").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}
	banners, err := banner.ListBanner(num)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, banners)
}
