package logging

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

//step4.1
var Logger *logrus.Logger

func NewLogger() *logrus.Logger {
	if Logger != nil {
		return Logger
	}
	dir, _ := os.Getwd()
	baseLogPath := path.Join(dir, "log", "highlight_gin.log")
	writer, _ := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	writerMap := lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}
	Logger = logrus.New()

	Logger.Hooks.Add(lfshook.NewHook(
		writerMap,
		&logrus.JSONFormatter{},
	))
	return Logger
}

func init() {

	Logger = NewLogger()
	// Log as JSON instead of the default ASCII formatter.
	//Log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//Log.SetLevel(log.WarnLevel)
	Logger.SetReportCaller(true)
}
