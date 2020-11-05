package user

import (
	"time"
	"wx-gin-master/models"
	"wx-gin-master/models/product"
	"wx-gin-master/pkg/logging"
)

type Favorite struct {
	PId       int
	UserID    int
	CreatedOn int64
}

func (Favorite) TableName() string {
	return "favorites"
}

// AddFavorite 添加商品到收藏夹
func AddFavorite(userId, pId int) bool {
	favs := ListFavorite(userId)
	for _, item := range favs {
		if item.PId == pId {
			return false
		}
	}
	fav := Favorite{UserID: userId, PId: pId, CreatedOn: time.Now().Unix()}
	if err := models.DB.Debug().Create(&fav).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// DestroyFavorite 把商品从收藏夹中移除
func DestroyFavorite(userId, pId int) bool {
	fav := Favorite{UserID: userId, PId: pId}
	if err := models.DB.Debug().Unscoped().Delete(&fav, "user_id = ? AND p_id = ?", fav.UserID, fav.PId).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// ListFavorite 根据Uid列出收藏的商品
func ListFavorite(userId int) (favorites []Favorite) {
	models.DB.Model(Favorite{}).Where("user_id = ?", userId).Find(&favorites)
	return
}

// ShowFavorites 显示收藏的商品
func ShowFavorites(pageNum int, pageSize int, userId int) (Products []product.Product, count int) {
	db := models.DB.Table("favorites").Select("p_id").Where("user_id = ?", userId)
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
	var product product.Product
	models.DB.Where("id in (?)", pids).Offset(pageNum).Limit(pageSize).Find(&Products)
	models.DB.Model(&product).Where("id in (?)", pids).Count(&count)
	return
}
