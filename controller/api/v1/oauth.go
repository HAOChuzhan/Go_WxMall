package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wx-gin-master/models"
	"wx-gin-master/models/user"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

const (
	APP_KEY    = "wxf5f58dbf63a03ed0"
	APP_SECRET = "1bb4430474fef81abab370d81b76d687"
	//Oauthurl   = "https://api.weixin.qq.com/sns/jscode2session"
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	code := c.PostForm("code")

	valid.Required(code, "code").Message("code是必须有的")
	//url := Oauthurl + "?appid=" + APP_KEY + "&secret=" + APP_SECRET + "&js_code=" + code + "&grant_type=authorization_code"
	url := fmt.Sprintf(code2sessionURL, APP_KEY, APP_SECRET, code)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return
	}
	var WXmap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&WXmap)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	open_id := WXmap["openid"]
	session_key := WXmap["session_key"]

	user1 := user.User{}
	data := make(map[string]interface{})
	data["nick_name"] = ""
	data["avatar"] = ""
	data["mobile"] = ""
	data["is_phone"] = 0
	data["is_oauth"] = 0
	if err := models.DB.Model(&user1).Where("open_id = ?", open_id).Where("del = ?", 0).First(&user1).Error; err != nil {
		//以下是openid找到的情况，如果没有找到的话。

		/*
			row, err := json.Marshal(user)
			phone := row["mobile"]
			user_id := row["id"]
		*/
		//phone := user1.Mobile
		user_id := user1.ID
		user2 := user.User{
			SessionKey: session_key,
		}
		//var user1 user.User

		models.DB.Model(&user2).Where("id = ?", user_id).Update(&user2)

		token, err := util.GenerateToken(user_id)
		if err != nil {
			return
		}
		data["token"] = token
		//modelss.DB.Model().Update
		if user1.Mobile != "" {
			data["is_phone"] = 1
			data["mobile"] = user1.Mobile
		}
		if user1.Nickname != "" && user1.Avatar != "" {
			data["is_oauth"] = 1
			data["nick_name"] = user1.Nickname // user.commonTextDecode
			data["avatar"] = user1.Avatar
		} else {
			user3 := user.User{
				Openid:     open_id,
				SessionKey: session_key,
			}
			user3.Del = 0
			models.DB.Select("id").Model(&user3).Create(&user3)
			api_token, err := util.GenerateToken(user3.ID)
			if err != nil {
				return
			}
			data["token"] = api_token
		}
	}

	appG.Response(http.StatusOK, e.OK, data)
}

/*
获取授权头像和昵称
*/
/*
func Oauth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	yb_user := c.Postform("user")
	Avatar := c.DefaultPostForm("avatar", "")

}
*/
