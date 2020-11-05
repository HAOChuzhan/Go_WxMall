package user

import (
	"time"
	"wx-gin-master/models"
	"wx-gin-master/models/product"
	"wx-gin-master/pkg/logging"
)

type Cart struct {
	ID     int `gorm:"primary_key" json:"id"`
	UserId int `json:"user_id"` // 用户id
	PId    int `json:"p_id"`    // 商品主键
	OId    int `json:"o_id"`    // 订单主键
	Number int `json:"number"`  // 数量
	Status int `json:"status"`  // 状态,0-未生成订单 1-已生成订单

	CreatedOn int64      `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Cart) TableName() string {
	return "carts"
}

// PutInCart 放进购物车
func PutInCart(userId, pId int) bool {
	if IsInCart(userId, pId) {
		cart := QueryCart(userId, pId)
		NumberIncrease(cart.ID)
		return true
	} else {
		if ok := CreateCartRow(userId, pId); ok {
			return true
		}
		return false
	}
}

// CreateCartRow 新增购物车商品
func CreateCartRow(userId, pId int) bool {
	cart := Cart{UserId: userId, PId: pId, Number: 1, CreatedOn: time.Now().Unix()}
	if err := models.DB.Debug().Create(&cart).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// CartsProducts 获取购物车中的商品
func CartsProducts(userId int) (products []product.Product) {
	db := models.DB.Table("carts").Select("p_id").Where("user_id = ?", userId)
	rows, err := db.Rows()
	if err != nil {
		logging.Info(err)
	}
	var pids []int
	for rows.Next() {
		var pid int
		_ = rows.Scan(&pid)
		pids = append(pids, pid)
	}
	models.DB.Where("id in (?)", pids).Find(&products)
	return
}

// NumberIncrease 添加购物车商品数量
func NumberIncrease(cId int) bool {
	cart := QueryCartById(cId)
	cart.Number += 1
	if err := models.DB.Debug().Model(cart).UpdateColumn("number", cart.Number).Error; err != nil {
		return false
	}
	return true
}

// NumberDecrease 减少购物车商品数量
func NumberDecrease(cId int) bool {
	cart := QueryCartById(cId)
	cart.Number -= 1
	if cart.Number <= 0 {
		DropFromCart(cId)
	}
	if err := models.DB.Debug().Model(cart).UpdateColumn("number", cart.Number).Error; err != nil {
		return false
	}

	return true
}

// DropFromCart 从购物车中移除
func DropFromCart(cId int) bool {
	cart := Cart{ID: cId}
	if err := models.DB.Debug().Delete(&cart, "id = ?", cart.ID).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// QueryCart 查询购物车表
func QueryCart(userId, pId int) (cart Cart) {
	models.DB.Debug().Model(Cart{}).Where("user_id = ? AND p_id = ?", userId, pId).First(&cart)
	return
}

// QueryCartById 通过Id获取购物车信息
func QueryCartById(cId int) (cart Cart) {
	models.DB.Debug().Model(Cart{}).Where("id = ?", cId).First(&cart)
	return
}

// IsInCart 判断是否在购物车中
func IsInCart(userId, pId int) bool {
	cars := ListCart(userId)
	for _, item := range cars {
		if item.PId == pId {
			return true
		}
	}
	return false
}

// ListCart 获取购物车列表
func ListCart(userId int) (carts []Cart) {
	models.DB.Debug().Model(Cart{}).Where("user_id = ?", userId).Find(&carts)
	return
}

// OrderCreated 创建订单
func OrderCreated(ids interface{}, oid int) bool {
	if err := models.DB.Debug().Model(Cart{}).Where("id in (?)", ids).UpdateColumn("o_id", oid).Error; err != nil {
		return false
	}

	return true
}

// 根据ids查询多个购物车信息
func QueryCarts(ids interface{}) (carts []Cart) {
	models.DB.Debug().Model(Cart{}).Where("id in (?)", ids).Find(&carts)
	return
}

// 通过订单id查询购物车
func QueryCartByOid(oid int) (carts []Cart) {
	models.DB.Debug().Model(Cart{}).Where("o_id = ?", oid).Find(&carts)
	return
}
