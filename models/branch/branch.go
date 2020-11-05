package branch

import (
	"time"
	"wx-gin-master/models"
)

type Branch struct {
	models.Model

	CorpId      int       `json:"corp_id"`
	AgentId     int       `json:"agent_id"`
	Secret      string    `json:"secret"` //企业密钥
	AccessToken string    `json:"access_token"`
	ExpireTime  time.Time `json:"expire_time"` //到期时间

}

func (Branch) TableName() string {
	return "branch"
}
