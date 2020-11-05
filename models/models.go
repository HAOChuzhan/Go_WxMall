package models

import (
	"fmt"
	"log"
	"time"
	"wx-gin-master/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//第二项我们仅导入而不使用。这个导入操作，gorm执行了下述操作
)

//数据库初始化
var DB *gorm.DB

type Model struct {
	ID         int        `gorm:"primary_key" json:"id"`
	Del        int        `json:"del"` //是否删除
	CreatedOn  time.Time  `json:"-"`
	ModifiedOn time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

// Setup 初始化数据库实例
func Setup() {
	var err error
	DB, err = gorm.Open(setting.DatabaseSetting.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("连接数据库失败，models.Setup err: %v", err)
	}

	//把回调函数注册进 GORM 的钩子里
	//创建回调
	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//更新回调
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	// 设置数据表前缀
	/*gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}*/

	//开启单表模式
	DB.SingularTable(true)
	//开启记录模式
	DB.LogMode(true)
	//设置最大空闲连接数
	DB.DB().SetMaxIdleConns(10)
	//设置最大开放连接数
	DB.DB().SetMaxOpenConns(100)
}

//updateTimeStampForCreateCallback 创建时将设置`CreatedOn`，`ModifiedOn`
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now() //.Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback 更新时将设置`ModifiedOn`
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now()) //.Unix()
	}
}
