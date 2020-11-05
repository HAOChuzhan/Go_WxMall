package export

import (
	"github.com/YFJie96/wx-mall/pkg/setting"
)

const EXT = ".xlsx"

// GetExcelFullUrl 获取Excel文件的完整访问路径
func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

// GetExcelPath 获取Excel文件的相对保存路径
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath 获取Excel文件的完整保存路径
func GetExcelFullPath() string {
	// runtime/export/
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
