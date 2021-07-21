package config

import (
	"github.com/jasonlvhit/gocron"
	_ "github.com/jasonlvhit/gocron"
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

func logHandle(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	//logs.Warn = log.New(warningHandle, "[WARNING]", log.Ldate|log.Ltime|log.Llongfile)
	//logs.Info = log.New(infoHandle, "[INFO]", log.Ldate|log.Ltime|log.Llongfile)
	//logs.Error = log.New(errorHandle, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
}

func logInit() {
	logHandle(ioutil.Discard, os.Stderr, os.Stdout)
	scheduler()
}

func LogFiles() {
	currentTime := time.Now()
	date := currentTime.String()
	Infolog.SetFormatter(&logrus.JSONFormatter{})
	Infolog.SetOutput(os.Stdout)
	logFile := &lumberjack.Logger{
		Filename:   Config["log_path"] + "/paas-ta-portal-api-v3-" + date[:10] + ".log",
		MaxSize:    500, /*log파일의 최대 사이즈*/
		MaxAge:     3,   /* 보존 할 최대 이전 로그 파일 수 */
		MaxBackups: 28,  /*타임 스탬프를 기준으로 오래된 로그 파일을 보관할 수 있는 최대 일수*/
		LocalTime:  false,
		Compress:   false, /*압축 여부*/
	}
	Infolog.SetOutput(logFile)                            // 로그파일을 만든다.
	Infolog.SetOutput(io.MultiWriter(logFile, os.Stdout)) /*console창에 로그내용을 출력한다*/
	Infolog.WithFields(logrus.Fields{
		"[INFO]": "SUCCESS",
	}).Info("[IN_SOFT] START SUCCEED FILE")
	logInit()
}

func ErrorFiles() {
	currentTime := time.Now()
	date := currentTime.String()
	Errorlog.SetFormatter(&logrus.JSONFormatter{})
	Errorlog.SetOutput(os.Stdout)
	logFile := &lumberjack.Logger{
		Filename:   Config["log_path"] + "/paas-ta-portal-api-v3-" + date[:10] + "-error.log",
		MaxSize:    500, /*log파일의 최대 사이즈*/
		MaxAge:     3,   /* 보존 할 최대 이전 로그 파일 수 */
		MaxBackups: 28,  /*타임 스탬프를 기준으로 오래된 로그 파일을 보관할 수 있는 최대 일수*/
		LocalTime:  false,
		Compress:   false, /*압축 여부*/
	}
	Errorlog.SetOutput(logFile)                            // 로그파일을 만든다.
	Errorlog.SetOutput(io.MultiWriter(logFile, os.Stdout)) /*console창에 로그내용을 출력한다*/
	Errorlog.WithFields(logrus.Fields{
		"[ERROR]": "FAIL",
	}).Error("[IN_SOFT] START ERROR FILE")
	logInit()
}

func scheduler() {
	//gocron.Every(1).Minute().Do(LogFiles)
	//gocron.Every(1).Minute().Do(ErrorFiles)
	gocron.Every(1).Hour().Do(LogFiles)
	gocron.Every(1).Hour().Do(ErrorFiles)
	gocron.Every(1).Day().At("00:00").Do(LogFiles)
	gocron.Every(1).Day().At("00:00").Do(ErrorFiles)
	// Begin job at a specific date/time
	infoT := time.Date(2021, time.July, 21, 9, 40, 0, 0, time.Local)
	errorT := time.Date(2021, time.July, 21, 9, 40, 0, 0, time.Local)
	gocron.Every(1).Hour().From(&infoT).Do(LogFiles)
	gocron.Every(1).Hour().From(&errorT).Do(LogFiles)
	// NextRun gets the next running time
	// Remove a specific job
	// gocron.Remove(task)
	// Clear all scheduled jobs
	// gocron.Clear()
	// Start all the pending jobs
	<-gocron.Start()
	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(LogFiles)
	<-s.Start()
}
