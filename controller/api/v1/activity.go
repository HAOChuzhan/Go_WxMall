package v1

import (
	"net/http"
	"wx-gin-master/models/activity"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"

	"github.com/unknwon/com"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func CreateActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	title := c.PostForm("title")
	img := c.PostForm("img")
	content := c.PostForm("content")
	abstract := c.PostForm("abstract")

	valid.Required(title, "title").Message("title 是必需的")
	valid.Required(img, "img").Message("img 是必需的")
	valid.Required(content, "content").Message("content 是必需的")
	valid.Required(abstract, "abstract").Message("abstract 是必需的")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	newactivity := activity.Activity{
		Title:    title,
		Img:      img,
		Content:  content,
		Abstract: abstract,
	}

	if !activity.Create(&newactivity) {
		appG.Response(http.StatusInternalServerError, e.ErrorCreatedOrderFail, nil)
		return
	}
	appG.Response(http.StatusCreated, e.Created, newactivity)

}

func UpdateActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 1, "id").Message("id 必须大于1 ")

	title := c.PostForm("title")
	img := c.PostForm("img")
	content := c.PostForm("content")
	abstract := c.PostForm("abstract")

	valid.Required(title, "title").Message("title 是必需的")
	valid.Required(img, "img").Message("img 是必需的")
	valid.Required(content, "content").Message("content 是必需的")
	valid.Required(abstract, "abstract").Message("abstract 是必需的")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Error)
		return
	}

	editedactivity := activity.Activity{
		Title:    title,
		Img:      img,
		Content:  content,
		Abstract: abstract,
	}
	editedactivity.ID = id
	editedactivity.Del = 0

	if !activity.Update(&editedactivity) {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusCreated, e.Created, editedactivity)
}

func DeleteActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 1, "id").Message("id 必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.BadRequest, nil)
		return
	}
	if !activity.Delete(id) {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusCreated, e.Created, nil)

}
func IndexActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{} //数据验证

	num := com.StrTo(c.PostForm("num")).MustInt()
	valid.Min(num, 1, "num").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, nil)
		return
	}
	activities, err := activity.List(num)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, activities)
}
