package user

import (
	"wx-gin-master/models"
	"wx-gin-master/pkg/logging"
)

type User struct {
	models.Model

	Password     string  `json:"-"`            // 密码
	Unionid      string  `json:"-"`            // unionId
	Openid       string  `json:"-"`            // openId
	SessionKey   string  `json:"-"`            // sessionKey
	Nickname     string  `json:"nickname"`     // 用户名
	Avatar       string  `json:"avatar"`       // 头像
	Sex          int     `json:"sex"`          // 性别,0-未知 1-男 2-女
	Mobile       string  `json:"mobile"`       // 手机
	Introduction string  `json:"introduction"` // 简介
	Balance      float64 `json:"balance"`      // 余额
	Coin         float64 `json:"coin"`         // 平台积分
}

func (User) TableName() string {
	return "users"
}

// GetInfo 根据id获取用户信息
func GetInfo(id int) (user User) {
	if err := models.DB.Debug().Where(map[string]interface{}{"id": id}).Find(&user).Error; err != nil {
		logging.Info(err)
	}
	return
}

// Login 登陆
func Login(user User) bool {
	var find User
	if models.DB.Where(user).Select("id").First(&find); find.ID > 0 {
		return true
	}
	return false
}

// CreateByPassword 根据密码新建用户
func CreateByPassword(user *User) bool {
	if models.DB.NewRecord(user) {
		models.DB.Debug().Create(user)
		return !models.DB.NewRecord(user)
	}

	return false
}

// Update 更新信息
func Update(user *User, data interface{}) bool {
	if err := models.DB.Debug().Model(user).Where("id = ?", user.ID).Update(data).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// UpdateColumn 更新指定列
func UpdateColumn(user *User, column string, data interface{}) bool {
	if err := models.DB.Debug().Model(User{}).Where("id = ?", user.ID).UpdateColumn(column, data).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// Create 新建用户
func Create(user *User) bool {
	if models.DB.NewRecord(user) {
		models.DB.Debug().Create(user)
		return !models.DB.NewRecord(user)
	}

	return false
}

// AddUnionid 添加 Unionid
func AddUnionid(userId int, unionid string) bool {
	if err := models.DB.Debug().Model(User{}).Where("id = ?", userId).UpdateColumn("unionid", unionid).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// QueryUserByUnionid 通过unionid获取用户信息
func QueryUserByUnionid(unionid string) (user User) {
	models.DB.Debug().Model(User{}).Where("unionid = ?", unionid).First(&user)
	return
}

// QueryUserByOpenid 通过Openid获取用户信息
func QueryUserByOpenid(openid string) (user User) {
	models.DB.Debug().Model(User{}).Where("openid = ?", openid).First(&user)
	return
}

// QueryUserById 通过id获取用户信息
func QueryUserById(id int) (user User) {
	models.DB.Debug().Model(User{}).Where("id = ?", id).First(&user)
	return
}

// PutSsk 添加SessionKey
func PutSsk(openid, ssk string) bool {
	user := QueryUserByOpenid(openid)
	user.SessionKey = ssk
	if err := models.DB.Debug().Model(user).Update(user).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}
