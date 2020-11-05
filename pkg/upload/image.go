package upload

import (
	"fmt"
	"github.com/YFJie96/wx-mall/pkg/file"
	"github.com/YFJie96/wx-mall/pkg/logging"
	"github.com/YFJie96/wx-mall/pkg/setting"
	"github.com/YFJie96/wx-mall/pkg/util"

	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// GetImageFullUrl 获取完整的访问路径
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName 获取图像名称
func GetImageName(name string) string {
	//path.Ext获取扩展名
	ext := path.Ext(name)
	//strings.TrimSuffix 修剪后缀
	fileName := strings.TrimSuffix(name, ext)
	//获取MD5码
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath 获取保存路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath 获取完整的保存路径
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt 检查图像文件后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		//strings.ToUpper 转成大写
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize 检查图像尺寸
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImage 检查文件是否存在
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
