package auth

import "wx-gin-master/models"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Unionid  string `json:"-"` //用户在开放平台下的唯一标识
	Openid   string `json:"-"` //用户在某一应用下的唯一标识，因此在不同的小程序或公众号中是不一样的。
}

type AuthAdmin struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Auth) TableName() string {
	return "users"
}

func (AuthAdmin) TableName() string {
	return "admin"
}

// CheckAndReturnId 检查数据库中是否有对应的手机号和密码并返回id
func CheckAndReturnId(mobile, password string) int {
	var auth Auth
	models.DB.Select("id").Where(Auth{Mobile: mobile, Password: password}).First(&auth)
	return int(auth.ID)
}

// CheckAdmin 检查管理员账号密码
func CheckAdmin(username, password string) int {
	var auth AuthAdmin
	models.DB.Debug().Select("id").Where(AuthAdmin{Username: username, Password: password}).First(&auth)
	return int(auth.ID)
}

// CheckOpenid 检查是否有openid
func CheckOpenid(openid string) int {
	var auth Auth
	models.DB.Debug().Select("id").Where(Auth{Openid: openid}).First(&auth)
	return int(auth.ID)
}
