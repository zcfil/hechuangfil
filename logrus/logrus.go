package logrus

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

// ContextHook 用来显示日志行号和文件名字
type ContextHook struct {
}
func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(8); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["file"] = path.Base(file)
		entry.Data["func"] = path.Base(funcName)
		entry.Data["line"] = line
	}
	return nil
}
// 一定要在这里初始化

var Log *logrus.Logger
func NewLogger(fileName string) *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  fileName,
		logrus.WarnLevel:  fileName,
		logrus.ErrorLevel: fileName,
		logrus.PanicLevel: fileName,
	}

	Log = logrus.New()
	Log.Hooks.Add(ContextHook{})
	Log.Hooks.Add(lfshook.NewHook(pathMap, &logrus.JSONFormatter{}))
	//Log.SetReportCaller(true)  // 显示行号
	//logrus.LogFunction()

	return Log
}


func Info(args... interface{}){
	Log.Info(args...)
}

func Warning(args... interface{}) {
	Log.Warning(args...)
}

func Error(args... interface{}) {
	Log.Error(args...)
}

func Fatal(args... interface{}) {
	Log.Fatal(args...)
}