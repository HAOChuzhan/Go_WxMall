package wallet

import (
	"wx-gin-master/models"
	"wx-gin-master/models/user"
	"wx-gin-master/pkg/logging"
)

type Wallet struct {
	models.Model
	Balance float64 `json:"balance"` // 余额
	Coin    float64 `json:"coin"`    // 平台币
}

func (Wallet) TableName() string {
	return "users"
}

// TopUpBalance 充值余额
func TopUpBalance(amount float64, userId int) bool {
	user := user.GetInfo(userId)
	wallet := Wallet{}
	wallet.Balance = user.Balance + amount
	if err := models.DB.Debug().Model(wallet).UpdateColumn("balance", wallet.Balance).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// CostBalance 花费余额
func CostBalance(amount float64, userId int) bool {
	user := user.GetInfo(userId)
	wallet := Wallet{}
	wallet.Balance = user.Balance - amount
	if wallet.Balance < 0 {
		return false
	}
	if err := models.DB.Debug().Model(wallet).UpdateColumn("balance", wallet.Balance).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}
