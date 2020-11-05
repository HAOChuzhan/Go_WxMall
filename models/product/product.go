package product

import (
	"wx-gin-master/models"
	"wx-gin-master/pkg/logging"

	"github.com/jinzhu/gorm"
)

type Product struct {
	models.Model

	Title       string  `json:"title"`        // 商品名
	Content     string  `json:"content"`      // 项目详情
	ProductType int     `json:"product_type"` // 项目类型
	Img         string  `json:"Img"`          // 商品封面图
	ShareNum    int     `json:"share_num"`    // 分享人数
	Price       float64 `json:"price"`        // 零售价
	IsHot       int     `json:"is_hot"`       // 是否为热门项目

	//Carousel     string  `json:"carousel"`      // 商品图集
	//Brand        string  `json:"brand"`         // 品牌名
	//Series       string  `json:"series"`        // 系列名
	//SellingPrice float64 `json:"selling_price"` // 售价
	//Cost         float64 `json:"cost"`          // 成本价
	//Tags         string  `json:"tags"`          // 标签,多个逗号分隔
	//Sales        int     `json:"sales"`         // 销量
	//Inventory    int     `json:"inventory"`     // 库存
	//Status       int     `json:"status"`        // 状态,0-下架 1-上架
	//OnSale       int     `json:"on_sale"`       // 状态,0-未折扣 1-折扣中
}

func (Product) TableName() string {
	return "product"
}

// Create 创建新产品
func Create(product *Product) bool {
	if models.DB.NewRecord(product) {
		models.DB.Create(product)
		if !models.DB.NewRecord(product) {
			return true
		}
	}
	return false
}

// UpDate 更新产品
func UpDate(date *Product) bool {
	if date.ID == 0 {
		return false
	}
	product := Product{}
	models.DB.Where("id = ?", date.ID).First(&product)
	if err := models.DB.Debug().Model(product).Update(date).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// Destroy 根据id 删除产品
func Destroy(id int) bool {
	if err := models.DB.Debug().Delete(Product{}, "id = ?", id).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// Show 根据 id 获取产品
func Show(id int) (product Product) {
	if err := models.DB.Debug().Where(map[string]interface{}{"id": id}).First(&product).Error; err != nil {
		logging.Info(err)
		return product
	}
	return
}

// List 返回符合条件的 product 列表
func List(pageNum int, pageSize int, maps interface{}) (products []Product, count int) {
	models.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&products)
	models.DB.Model(Product{}).Where(maps).Count(&count)
	return
}

// SearchInTitle 通过商品名搜索
func SearchInTitle(pageNum int, pageSize int, data string) (products []Product, count int) {
	models.DB.Debug().Where("title LIKE ?", data).Offset(pageNum).Limit(pageSize).Find(&products)
	models.DB.Debug().Model(Product{}).Where("`title` LIKE ?", data).Count(&count)
	return
}

/*
// SearchInTag 通过标签搜索
func SearchInTag(pageNum int, pageSize int, data string) (products []Product, count int) {
	data = fmt.Sprintf("%%%s%%", data)
	models.DB.Debug().Where("tags LIKE ?", data).Offset(pageNum).Limit(pageSize).Find(&products)
	models.DB.Debug().Model(Product{}).Where("`tags` LIKE ?", data).Count(&count)
	return
}
*/
// UpdateInventory 增加或减少库存
func UpdateInventory(id int, num int) bool {
	if err := models.DB.Debug().Table("product").Where("id = ?", id).
		Update("inventory", gorm.Expr("inventory + ?", num)).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}
