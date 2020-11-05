package admin

import (
	"wx-gin-master/models"
	"wx-gin-master/pkg/logging"
)

type Admin struct {
	models.Model
	Nickname string `json:"nick_name"` // 昵称
	Avatar   string `json:"avatar"`    // 头像
	Username string `json:"username"`  // 账号
	Password string `json:"-"`         // 密码
}

func (Admin) TableName() string {
	return "admin"
}

// GetInfo 通过id获取管理员信息
func GetInfo(id int) (user Admin) {
	if err := models.DB.Debug().Where(map[string]interface{}{"id": id}).Find(&user).Error; err != nil {
		logging.Info(err)
	}
	return
}

// Login 登录
func Login(user Admin) bool {
	var find Admin
	if models.DB.Where(user).Select("id").First(&find); find.ID > 0 {
		return true
	}
	return false
}
