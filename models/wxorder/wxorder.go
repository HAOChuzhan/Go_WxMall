package wxorder

import (
	"wx-gin-master/models"
	"wx-gin-master/pkg/logging"
)

type WxOrder struct {
	models.Model
	UserId        int     `json:"user_id"`        // 用户id
	AppID         int     `json:"app_id"`         // APPID
	Body          string  `json:"body"`           // 商品描述
	OutTradeNo    string  `json:"out_trade_no"`   // 商户订单号
	NonceStr      string  `json:"nonce_str"`      // 随机字符串
	Sign          string  `json:"sign"`           // 签名
	SumPay        float64 `json:"sum_pay"`        // 结算金额
	TotalFee      int     `json:"total_fee"`      // 标价金额,单位为积分
	Detail        string  `json:"detail"`         // 商品详情
	Attach        string  `json:"attach"`         // 附加数据
	TransactionId string  `json:"transaction_id"` // 微信支付订单号
}

func (WxOrder) TableName() string {
	return "wxorders"
}

// Create 新建订单
func Create(order *WxOrder) bool {
	if models.DB.NewRecord(*order) {
		models.DB.Debug().Create(*order)
		return !models.DB.NewRecord(*order)
	}

	return false
}

// Update 更新订单
func Update(order *WxOrder) bool {
	if err := models.DB.Debug().Model(WxOrder{}).Update(*order).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// QueryByTradeNo 通过订单号查询订单信息
func QueryByTradeNo(no string) (order WxOrder) {
	models.DB.Debug().Model(&order).Where("tradeNo = ?", no).First(&order)
	return
}

// Destroy 销毁订单
func Destroy(order *WxOrder) bool {
	if err := models.DB.Debug().Model(WxOrder{}).Delete(*order).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}
