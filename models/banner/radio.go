package banner

import (
	"wx-gin-master/models"
	"wx-gin-master/pkg/logging"
)

type Radio struct {
	models.Model
	//Del      int    `json:"del"`       //是否删除
	Title    string `json:"title"`     //广播标题
	JumpType string `json:"jump_type"` //广播的类型
}

func (Radio) TableName() string {
	return "radio"
}

func CreatRadio(radio *Radio) bool {
	if models.DB.NewRecord(radio) {
		models.DB.Create(radio)
		if !models.DB.NewRecord(radio) {
			return true
		}
	}
	return false
}

func UpdateRadio(newradio *Radio) bool {
	if newradio.ID == 0 {
		return false
	}
	radio := Radio{}
	models.DB.Where("id = ?", newradio.ID).First(&radio)

	if err := models.DB.Debug().Model(radio).Update(newradio).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

//根据id删除radio
func DeleteRadio(id int) bool {

	models.DB.Model(&Radio{}).Where("id = ?", id).Update("del", 1)
	if err := models.DB.Debug().Delete(Radio{}, "id = ?", id).Error; err != nil {
		logging.Info(err)
		return false
	}

	return true
}

// 通过ID查询广播信息
func QueryRadioById(id int) (radio *Radio) {
	models.DB.Debug().Model(Radio{}).Where("id = ?", id).First(&radio)
	return
}

func ListRadio(num int) (radio []Radio, err error) {
	if err = models.DB.Debug().Order("modified_on DESC").Limit(num).Find(&radio).Error; err != nil {
		logging.Info(err)
		return
	}
	return
}
