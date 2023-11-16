package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

/*
	Settings 日志配置信息
*/
type Settings struct {
	Path       string `yaml:"Path"`        // 日志路径
	Name       string `yaml:"Name"`        // 日志文件名称
	Ext        string `yaml:"Ext"`         // 扩展名
	TimeFormat string `yaml:"time-format"` // 时间戳格式
}

var (
	logFile            *os.File
	logger             *log.Logger
	defaultPrefix      = "" // 默认前缀
	defaultCallerDepth = 2
	mu                 sync.Mutex
	logPrefix          = ""
	logLevelFLags      = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type logLevel int

const (
	DEBUG logLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)
const flag = log.LstdFlags

func init() {
	logger = log.New(os.Stdout, defaultPrefix, flag)
}

// 初始化配置 logger

func Setup(settings *Settings) {
	var err error
	// 设置log名称
	fileName := fmt.Sprintf("%s-%s.%s",
		settings.Name,
		time.Now().Format(settings.TimeFormat),
		settings.Ext)

	logFile, err = mustOpen(fileName, settings.Path)
	if err != nil {
		log.Fatalf("logging setup error %s", err)
		return
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	logger = log.New(mw, defaultPrefix, flag)
}

/*
	setPrefix 设置日志信息前缀
*/
func setPrefix(level logLevel) {

	logPrefix = fmt.Sprintf("[%s] ", logLevelFLags[level])
	_, file, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d] ", logLevelFLags[level], filepath.Base(file), line)
	}
	logger.SetPrefix(logPrefix)
}

/*
	Debug 日志等级
*/
func Debug(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(DEBUG)
	logger.Println(v...)
}

/*
	Info 日志等级
*/
func Info(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(INFO)
	logger.Println(v...)
}

/*
	Error 日志等级
*/
func Error(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(ERROR)
	logger.Println(v...)
}

/*
	WARN 日志等级
*/
func Warn(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(WARN)
	logger.Println(v...)
}

/*
	Fatal 日志等级
*/
func Fatal(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(FATAL)
	logger.Println(v...)
}
