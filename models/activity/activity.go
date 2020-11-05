package activity

import (
	"wx-gin-master/models"
	"wx-gin-master/pkg/logging"
)

type Activity struct {
	models.Model
	Title    string `json:"title"`
	Img      string `json:"img"`
	Content  string `json:"content"`
	Abstract string `json:"abstract"`
}

func (Activity) TableName() string {
	return "activity"
}

func Create(newactivity *Activity) bool {
	if models.DB.NewRecord(newactivity) {
		models.DB.Create(newactivity)
		if !models.DB.NewRecord(newactivity) {
			return true
		}
	}
	return false
}

func Update(data *Activity) bool {
	if data.ID == 0 {
		return false
	}
	activity := Activity{}
	models.DB.Where("id = ?", data.ID).First(&activity)
	if err := models.DB.Debug().Model(activity).Update(data).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

func Delete(id int) bool {
	models.DB.Model(&Activity{}).Where("id = ?", id).Update("del", 1)
	if err := models.DB.Debug().Delete(Activity{}, "id=?", id).Error; err != nil {
		logging.Info(err)
		return false
	}
	return true
}

func List(num int) (activity []Activity, err error) {
	if err = models.DB.Order("modified_on DESC").Limit(num).Find(&activity).Error; err != nil {
		logging.Info(err)
		return
	}
	return
}

func ExportActivity() {

}
