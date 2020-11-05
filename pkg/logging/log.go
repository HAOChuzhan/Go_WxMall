package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"wx-gin-master/pkg/file"
	"wx-gin-master/pkg/setting"
)

// getLogFilePath 获取日志文件保存路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// getLogFileName 获取日志文件的保存名称
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup 初始化日志实例
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("打开日志文件失败，logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug 在调试级别输出日志
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println("This is a DEBUG LOG")
	//logger.Println(v)
}

// Info 以信息级别输出日志
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println("This is a Info LOG")
}

// Warn 警告级别的输出日志
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println("This is a WARNING LOG")
}

// Error 以错误级别输出日志
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println("This is a ERROR LOG")
}

// Fatal 致命级别的输出日志
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln("This is a FATAL LOG")
}

// setPrefix 设置日志输出的前缀
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
