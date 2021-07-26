package config

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type logConfig struct {
	Warn  *log.Logger `yaml:"warn"`
	Info  *log.Logger `yaml:"info"`
	Error *log.Logger `yaml:"error"`
}

var logs *logConfig

var Infolog *logrus.Logger = logrus.New()
var Errorlog *logrus.Logger = logrus.New()

func LogHandle(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	//logs.Warn = log.New(warningHandle, "[WARNING]", log.Ldate|log.Ltime|log.Llongfile)
	//Infolog.Info(infoHandle, "[INFO]", log.Ldate|log.Ltime|log.Llongfile)
	//logs.Error = log.New(errorHandle, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
}

func ErrorLogHandle(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	//logs.Warn = log.New(warningHandle, "[WARNING]", log.Ldate|log.Ltime|log.Llongfile)
	//Infolog.Info(infoHandle, "[INFO]", log.Ldate|log.Ltime|log.Llongfile)
	//logs.Error = log.New(errorHandle, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
}

func LogInit() {
	LogHandle(ioutil.Discard, os.Stderr, os.Stdout)
}

func ErrorInit() {
	ErrorLogHandle(ioutil.Discard, os.Stderr, os.Stdout)
}

func LogFiles() {
	currentTime := time.Now()
	date := currentTime.String()
	Infolog.SetOutput(os.Stdout)

	InfoLogFile := &lumberjack.Logger{
		Filename:   Config["log_path"] + "/paas-ta-portal-api-v3-" + date[:10] + ".log",
		MaxSize:    500, /*log파일의 최대 사이즈*/
		MaxAge:     3,   /* 보존 할 최대 이전 로그 파일 수 */
		MaxBackups: 28,  /*타임 스탬프를 기준으로 오래된 로그 파일을 보관할 수 있는 최대 일수*/
		LocalTime:  false,
		Compress:   false, /*압축 여부*/
	}
	Infolog.SetOutput(InfoLogFile)

	Infolog.SetFormatter(&logrus.TextFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "[TIME]",
			logrus.FieldKeyLevel: "[LEVEL]",
			logrus.FieldKeyMsg:   "/",
		},
	})

	// 로그파일에 내용출력 만든다.
	Infolog.SetOutput(io.MultiWriter(InfoLogFile, os.Stdout)) /*console창에 로그내용을 출력한다*/
	Infolog.WithFields(logrus.Fields{
		"[INFO]": "SUCCESS",
	}).Info("[IN_SOFT] START SUCCESS FILE")

	LogInit()
}

func ErrorFiles() {
	currentTime := time.Now()
	date := currentTime.String()
	Errorlog.SetOutput(os.Stdout)

	ErrorLogFile := &lumberjack.Logger{
		Filename:   Config["log_path"] + "/paas-ta-portal-api-v3-" + date[:10] + "-error.log",
		MaxSize:    500, /*log파일의 최대 사이즈*/
		MaxAge:     3,   /* 보존 할 최대 이전 로그 파일 수 */
		MaxBackups: 28,  /*타임 스탬프를 기준으로 오래된 로그 파일을 보관할 수 있는 최대 일수*/
		LocalTime:  false,
		Compress:   false, /*압축 여부*/
	}
	Errorlog.SetOutput(ErrorLogFile)

	Errorlog.SetFormatter(&logrus.TextFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "[TIME]",
			logrus.FieldKeyLevel: "[LEVEL]",
			logrus.FieldKeyMsg:   "/",
		},
	})
	// Errorlog.SetOutput(logFile)                            // 로그파일을 만든다.
	Errorlog.SetOutput(io.MultiWriter(ErrorLogFile, os.Stdout)) /*console창에 로그내용을 출력한다*/
	Errorlog.WithFields(logrus.Fields{
		"[ERROR]": "FAIL",
	}).Error("[IN_SOFT] START ERROR FILE")

	ErrorInit()
}

func Schedular() {
	gocron.Every(30).Minutes().Do(LogFiles)
	gocron.Every(30).Minutes().Do(ErrorFiles)
	<-gocron.Start()
}
