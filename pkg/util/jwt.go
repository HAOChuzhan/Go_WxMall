package util

import (
	"time"
	"wx-gin-master/models/admin"
	"wx-gin-master/models/user"
	"wx-gin-master/pkg/setting"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

// Claims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type Claims struct {
	User user.User
	jwt.StandardClaims
}

type AdminClaims struct {
	User admin.Admin
	jwt.StandardClaims
}

// GenerateToken 生成用于身份验证的令牌
func GenerateToken(id int) (string, error) {
	nowTime := time.Now()
	//设置到期时间为3个小时后
	expireTime := nowTime.Add(3 * time.Hour)
	user := user.GetInfo(id)
	claims := Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-init",
		},
	}
	// 使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// GenerateAdmin 生成管理员Token
func GenerateAdmin(id int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	user := admin.GetInfo(id)
	claims := AdminClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			Issuer:    "gin-wx-mall",     //签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// ParseAdmin 解析管理员token
func ParseAdmin(token string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
