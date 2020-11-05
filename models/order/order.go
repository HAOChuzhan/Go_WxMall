package order

import (
	"wx-gin-master/models"
	"wx-gin-master/models/wallet"
	"wx-gin-master/pkg/logging"
)

type Order struct {
	models.Model
	UserId int `json:"user_id"` // 用户id
	//ProvinceName string  `json:"province_name"` // 省份
	//CityName     string  `json:"city_name"`     // 城市
	//CountyName   string  `json:"county_name"`   // 区县
	//DetailInfo   string  `json:"detail_info"`   // 详细地址
	//PostalCode   string  `json:"postal_code"`   // 邮编
	UserName  string `json:"user_name"`  // 收货人姓名
	TelNumber string `json:"tel_number"` // 联系电话
	//ExpressTitle string  `json:"express_title"` // 物流公司
	//ExpressCode  string  `json:"express_code"`  // 物流编号
	//ExpressNo    string  `json:"express_no"`    // 物流单号
	ExpressTime int64   `json:"express_time"` // 发货时间
	Total       float64 `json:"total"`        // 汇总金额
	SumPay      float64 `json:"sum_pay"`      // 结算金额
	Status      int     `json:"status"`       // 状态,0-未结算 1-已结算(待发货) 2-已发货(待收货) 3-已完成 9-异常
	PushedAt    int64   `json:"pushed_at"`    // 推送时间
}

func (Order) TableName() string {
	return "orders"
}

// Created 创建订单
func Created(order *Order) bool {
	if err := models.DB.Debug().Create(order).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// Settlement 完成支付
func Settlement(userId, oid int) bool {
	order := QueryOrderById(oid)
	if wallet.CostBalance(order.SumPay, userId) {
		if err := models.DB.Debug().Model(order).UpdateColumn("status", 1).Error; err != nil {
			logging.Info(err)
			return false
		}
		return true
	} else {
		return false
	}
}

// Update 更新订单信息
func Update(id int, order *Order) bool {
	if err := models.DB.Debug().Model(Order{}).Where("id = ?", id).Update(*order).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// Destroy 销毁订单
func Destroy(oid int) bool {
	if err := models.DB.Unscoped().Delete(Order{}, "id = ?", oid).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// Done 完成订单
func Done(oid int) bool {
	if err := models.DB.Delete(Order{}, "id = ?", oid).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// QueryOrderByUserId 通过用户ID查询订单信息
func QueryOrderByUserId(userId int) (order []Order) {
	models.DB.Debug().Model(Order{}).Where("user_id = ?", userId).Find(&order)
	return
}

// 通过ID查询订单信息
func QueryOrderById(id int) (order Order) {
	models.DB.Debug().Model(Order{}).Where("id = ?", id).First(&order)
	return
}

// IsOwner 判断订单ID和用户ID是否是同一个用户所拥有
func IsOwner(userId, id int) bool {
	var data int
	models.DB.Debug().Model(Order{}).Select("1").Where("id = ? AND user_id = ?", id, userId).First(&data)
	if data != 0 {
		return true
	}
	return false
}

/*
 * 导出订单
 * @param $offset = 0
 * @param $limit
 * @param $order_id
 * @return array
 */
/*
func Export() {
	models.DB.Order("modified_on DESC").Offset(offset).array
}
*/
