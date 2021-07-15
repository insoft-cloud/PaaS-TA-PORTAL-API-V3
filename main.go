package main

import (
	"PAAS-TA-PORTAL-V3/admin"
	"PAAS-TA-PORTAL-V3/app_features"
	"PAAS-TA-PORTAL-V3/app_usage_events"
	"PAAS-TA-PORTAL-V3/apps"
	"PAAS-TA-PORTAL-V3/audit_events"
	"PAAS-TA-PORTAL-V3/buildpacks"
	"PAAS-TA-PORTAL-V3/builds"
	"PAAS-TA-PORTAL-V3/config"
	"PAAS-TA-PORTAL-V3/deployments"
	"PAAS-TA-PORTAL-V3/domains"
	"PAAS-TA-PORTAL-V3/droplets"
	"PAAS-TA-PORTAL-V3/environment_variable_groups"
	"PAAS-TA-PORTAL-V3/feature_flags"
	"PAAS-TA-PORTAL-V3/info"
	"PAAS-TA-PORTAL-V3/isolation_segment"
	"PAAS-TA-PORTAL-V3/jobs"
	"PAAS-TA-PORTAL-V3/organization_quotas"
	"PAAS-TA-PORTAL-V3/organizations"
	"PAAS-TA-PORTAL-V3/packages"
	"PAAS-TA-PORTAL-V3/processes"
	"PAAS-TA-PORTAL-V3/resource_matches"
	"PAAS-TA-PORTAL-V3/roles"
	"PAAS-TA-PORTAL-V3/routes"
	"PAAS-TA-PORTAL-V3/service_brokers"
	"PAAS-TA-PORTAL-V3/stacks"
	"PAAS-TA-PORTAL-V3/tasks"
	"PAAS-TA-PORTAL-V3/users"
	"fmt"
	_ "fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	eureka "github.com/xuanbo/eureka-client"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"io/ioutil"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	"os"
	"time"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name cf-Authorization
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	admin.AdminHandleRequests(myRouter)
	apps.AppHandleRequests(myRouter)
	app_features.AppFeatureHandleRequests(myRouter)
	app_usage_events.AppUsageEventHandleRequests(myRouter)
	audit_events.AuditEventHandleRequests(myRouter)
	builds.BuildPackHandleRequests(myRouter)
	buildpacks.BuildPackHandleRequests(myRouter)
	deployments.DeploymentHandleRequests(myRouter)
	service_brokers.ServiceBrokerHandleRequests(myRouter)
	organizations.OrganizationsRequests(myRouter)
	domains.DomainHandleRequests(myRouter)
	droplets.DropletHandleRequests(myRouter)
	organization_quotas.OrganizationQuotasHandleRequests(myRouter)
	packages.PackagesHandleRequests(myRouter)
	environment_variable_groups.EnvironmentVariableGroupsHandleRequests(myRouter)
	feature_flags.FeatureFlagHandleRequests(myRouter)
	info.InforHandleRequests(myRouter)
	isolation_segments.IsolationSegmentsHandleRequests(myRouter)
	jobs.JobsHandleRequests(myRouter)
	stacks.AppHandleRequests(myRouter)
	tasks.TaskHandleRequests(myRouter)
	users.UserHandleRequests(myRouter)
	resource_matches.ResourceMatchesHandleRequests(myRouter)
	roles.RoleHandleRequests(myRouter)
	processes.ProcessHandleRequests(myRouter)
	routes.RouteHandleRequests(myRouter)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type", "cf-Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodTrace, http.MethodOptions})

	log.Fatal(http.ListenAndServe(":"+config.Config["port"], handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))
}

func main() {
	scheduler()
	config.SetConfig()
	config.ClientSetting()
	config.ValidateConfig()
	logFiles()
	// Eureka()
	handleRequests()
}

func Eureka() {
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://127.0.0.1:2221/eureka/",
		App:                   "PORTAL-API-V3",
		Port:                  2222,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})
	client.Start()
}

type Logger struct {
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

var logs Logger

func logHandle(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	logs.Warn = log.New(warningHandle, "[WARNING]", log.Ldate|log.Ltime|log.Llongfile)
	logs.Info = log.New(infoHandle, "[INFO]", log.Ldate|log.Ltime|log.Llongfile)
	logs.Error = log.New(errorHandle, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
}

func logInit() {
	logHandle(ioutil.Discard, os.Stderr, os.Stdout)
	logs.Info.Println("INFO")
	logs.Warn.Println("WARM")
	logs.Error.Println("ERROR")
}

func logFiles() {
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
	logrus.SetOutput(io.MultiWriter(logFile, os.Stdout)) /*console창*/
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true) /*이벤트 발생 함수, 파일명이 찍힘*/
	logrus.WithFields(logrus.Fields{
		"[INFO]": "SUCCESS",
	}).Info("SUCCESS")

	logInit()
}

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func scheduler() {

	//gocron.Every(1).Second().Do(logFiles)
	//gocron.Every(2).Seconds().Do(logFiles)
	gocron.Every(1).Minute().Do(logFiles) //분 단위 로그찍기
	//gocron.Every(2).Minutes().Do(logFiles)
	//gocron.Every(1).Hour().Do(logFiles)
	//gocron.Every(2).Hours().Do(logFiles)
	//gocron.Every(1).Day().Do(logFiles)
	//gocron.Every(2).Days().Do(logFiles)
	//gocron.Every(1).Week().Do(logFiles)
	//gocron.Every(2).Weeks().Do(logFiles)

	// Do jobs on specific weekday
	//gocron.Every(1).Monday().Do(logFiles)
	//gocron.Every(1).Thursday().Do(logFiles)

	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	gocron.Every(1).Day().At("00:00").Do(logFiles)
	//gocron.Every(1).Thursday().At("00:00").Do(logFiles)

	// Begin job immediately upon start
	gocron.Every(1).Hour().From(gocron.NextTick()).Do(logFiles)

	// Begin job at a specific date/time
	t := time.Date(2021, time.July, 15, 10, 0, 0, 0, time.Local)
	gocron.Every(1).Hour().From(&t).Do(logFiles)

	// NextRun gets the next running time
	_, time := gocron.NextRun()
	fmt.Println(time)

	// Remove a specific job
	// gocron.Remove(task)

	// Clear all scheduled jobs
	// gocron.Clear()

	// Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(logFiles)
	<-s.Start()
}
