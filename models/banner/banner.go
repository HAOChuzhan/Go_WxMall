package banner

/*
2020年10月29日完善
*/
import (
	"wx-gin-master/models"
	"wx-gin-master/pkg/logging"
)

type Banner struct {
	models.Model

	Name     string `json:"name"`      // Banner名称
	ImageUrl string `json:"image_url"` // 轮播图地址
	RelateId int    `json:"relate_id"` // 执行关键字，根据不同的type含义不同
	JumpType int    `json:"jump_type"` // 跳转类型 0，无导向；1：导向商品;2:导向专题
}

func (Banner) TableName() string {
	return "banner"
}

// Create 创建新轮播
func CreateBanner(banner *Banner) bool {
	if models.DB.NewRecord(banner) {
		models.DB.Create(banner)
		if !models.DB.NewRecord(banner) {
			return true
		}
	}
	return false
}

// UpDate 更新轮播图
func UpdateBanner(date *Banner) bool {
	if date.ID == 0 {
		return false
	}
	banner := Banner{}
	models.DB.Where("id = ?", date.ID).First(&banner)

	if err := models.DB.Debug().Model(banner).Update(date).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

/*
func UpdateRadio(newradio *Radio) bool {
	if newradio.ID == 0 {
		return false
	}
	radio := Radio{}
	models.DB.Where("id = ?", newradio.ID).First(&radio)

	if err := models.DB.Debug().Model(radio).Update(newradio); err != nil {
		logging.Info(err)
		return true
	}
	return false
}
*/

// Destroy 根据id 删除轮播图
func DeleteBanner(id int) bool {
	models.DB.Model(&Banner{}).Where("id = ?", id).Update("del", 1)
	if err := models.DB.Debug().Delete(Banner{}, "id = ?", id).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

// List 返回num个轮播图
func ListBanner(num int) (banner []Banner, err error) {
	if err = models.DB.Debug().Order("modified_on DESC").Limit(num).Find(&banner).Error; err != nil {
		logging.Info(err)
		return
	}
	return
}
