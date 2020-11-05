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

func CreatRadio(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	//var addedRadio banner.Radio

	title := c.PostForm("title")
	//keyword := com.StrTo(c.PostForm("keyword")).MustInt()
	jumptype := c.PostForm("jumptype")

	valid.Required(title, "title").Message("title 是必需的")
	//valid.Numeric(keyword, "keyword").Message("keyword 必须是有效数值")
	valid.Required(jumptype, "jumptype").Message("jumptype 必需的")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}
	addedRadio := banner.Radio{
		Title:    title,
		JumpType: jumptype,
	}
	addedRadio.Del = 0 //初始化是否被删除的参数为0
	add_result := banner.CreatRadio(&addedRadio)

	if !add_result {
		appG.Response(http.StatusInternalServerError, e.ErrorCreatedOrderFail, nil)
		return
	}
	appG.Response(http.StatusCreated, e.Created, addedRadio)
}

func UpdateRadio(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	title := c.PostForm("title")
	//keyword := com.StrTo(c.PostForm("keyword")).MustInt()
	jumptype := c.PostForm("jumptype")

	valid.Required(title, "title").Message("title 是必需的")
	//valid.Numeric(keyword, "keyword").Message("keyword 必须是有效数字")
	valid.Required(jumptype, "jumptype").Message("jumptype 必须的")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}
	edditedRadio := banner.Radio{
		Title:    title,
		JumpType: jumptype,
	}
	//edditedRadio.ModifiedOn = //time.Now() //.Format("2006/01/02 15:04:05")
	//edditedRadio.ModifiedOn.Format("2006/01/02 15:04:05")

	edditedRadio.ID = id
	edditedRadio.Del = 0

	if !banner.UpdateRadio(&edditedRadio) {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, nil)

}

func DeleteRadio(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}

	delete_redult := banner.DeleteRadio(id)

	//models.DB.Model(&radio).

	if !delete_redult {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, nil)
}

//数据库的索引
func IndexRadio(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{} //数据验证

	num := com.StrTo(c.PostForm("num")).MustInt()
	valid.Min(num, 1, "num").Message("num必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}
	radios, err := banner.ListRadio(num)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, radios)
}
