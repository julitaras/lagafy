package main

import (
	_ "api-dashboard/docs"
	"api-dashboard/helpers"
	"api-dashboard/models"
	"api-dashboard/pkg/setting"
	"api-dashboard/pkg/util"
	"api-dashboard/router"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Lagafy API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

var (
	//InfoLogger log
	InfoLogger *log.Logger
)

func init() {
	setting.Setup()
	util.Setup()
}

func main() {
	client, err := gorm.Open("mysql",
		setting.AppSetting.DbUser+":"+setting.AppSetting.DbPassword+"@("+setting.AppSetting.DbServer+":"+setting.AppSetting.DbPort+")/"+setting.AppSetting.DbName+"?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		fmt.Println(setting.AppSetting.DbUser + ":" + setting.AppSetting.DbPassword + "@(" + setting.AppSetting.DbServer + ":" + setting.AppSetting.DbPort + ")/" + setting.AppSetting.DbName + "?charset=utf8&parseTime=True&loc=UTC")
		log.Fatal(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//Migration
	client.AutoMigrate(&models.Travel{}, &models.Passenger{}, &models.Reservation{})

	//Foreign Keys
	client.Model(&models.Reservation{}).AddForeignKey("travel_id", "travels(id)", "RESTRICT", "RESTRICT")
	client.Model(&models.Reservation{}).AddForeignKey("passenger_id", "passengers(id)", "RESTRICT", "RESTRICT")
	client.Model(&models.Travel{}).Related(&models.Reservation{})

	//Create server
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))
	gin.SetMode(setting.AppSetting.GinMode)

	gin.DisableConsoleColor()

	gin.DefaultWriter = io.MultiWriter(getLogFile(setting.AppSetting.LogFile, true), os.Stdout)

	InfoLogger = log.New(getLogFile(setting.AppSetting.InfoLogFile, true), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//Warning = log.New(getLogFile(_warnLogFileName), "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	//Error = log.New(getLogFile(_errorLogFileName), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	//handlers.InfoLogger = InfoLogger

	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	router.SetRoutes(r, client)

	endPoint := fmt.Sprintf(":%d", setting.AppSetting.HTTPPort)

	srv := &http.Server{
		Addr:    endPoint,
		Handler: r,
	}

	//Run server in goroutine
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Wait for interrupt signal to gracefully shutdown the server with 5 seconds timeout
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// Quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}

func getLogFile(fileName string, rotate bool) *os.File {

	if rotate {
		fileName = fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02"), fileName)
	}

	var flag int

	if helpers.FileExists(context.Background(), fileName) {
		flag = os.O_APPEND
	} else {
		flag = os.O_CREATE
	}

	f, _ := os.OpenFile(fileName, flag, 0776)
	return f
}
