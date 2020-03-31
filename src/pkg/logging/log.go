package logging

import (
	"fmt"
	"os"
	"log"
	"runtime"
	"goms/pkg/setting"
	"goms/pkg/file"
)

var (
	F *os.File
	DefaultPrefix 		= 	""
	DefaultCallerDepth 	= 	2
	logger *log.Logger
)

func Setup() {
	F, err := file.MustOpen(setting.AppSetting.SavePath, setting.AppSetting.LoggingFile)
	if err != nil {
		fmt.Println("Log setup: ", err)
		log.Fatal(err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix("DEBUG")
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix("INFO")
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix("WARN")
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix("Error")
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix("FATAL")
	logger.Println(v)
}
 
func setPrefix(level string) {
	logPrefix  := ""
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s]	[%s:%d]", level, file, line)
	} else {
		logPrefix = fmt.Sprintf("[%s]	", level)
	}

	logger.SetPrefix(logPrefix)
}