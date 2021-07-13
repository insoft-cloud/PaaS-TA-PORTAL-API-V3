package main

import (
	_ "fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func main() {
	currentTime := time.Now()
	date := currentTime.String()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logFile := &lumberjack.Logger{
		Filename:   "./paas-ta-portal-api-v3-" + date[:10] + ".log",
		MaxSize:    500, /*log파일의 최대 사이즈*/
		MaxAge:     3,   /* 보존 할 최대 이전 로그 파일 수 */
		MaxBackups: 28,  /*타임 스탬프를 기준으로 오래된 로그 파일을 보관할 수 있는 최대 일수*/
		LocalTime:  false,
		Compress:   false, /*압축 여부*/
	}
	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true) /*이벤트 발생 함수, 파일명이 찍힘*/
	logrus.WithFields(logrus.Fields{
		"INFO": "SUCCESS",
	}).Info("아이엔소프트 LOG TEST")
}
