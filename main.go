package main

import (
	appmiddleware "cleanbase/middleware"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "net/http/pprof"

	_userHttpDeliver "cleanbase/app/user/handler"
	_userRepo "cleanbase/app/user/repository"
	_userUsecase "cleanbase/app/user/usecase"
)

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	logrus.Info("Using Config fil: ", viper.ConfigFileUsed())
	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("Comment service is Running in Debug Mode")
		return
	}
}

func init() {
	initConfig()
}

func main() {
	dbHost := viper.GetString("mysql.host")
	dbPort := viper.GetString("mysql.port")
	dbUser := viper.GetString("mysql.user")
	dbPass := viper.GetString("mysql.pass")
	dbName := viper.GetString("mysql.name")
	dsn := dbUser + `:` + dbPass + `@tcp(` + dbHost + `:` + dbPort + `)/` + dbName + `?parseTime=1&loc=Asia%2FJakarta`

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Minute * 10)

	echoServer := echo.New()
	echoServer.Debug = viper.GetBool("debug")
	middL := appmiddleware.InitMiddleware()
	echoServer.Use(middL.CORS)
	echoServer.HTTPErrorHandler = appmiddleware.ErrorHandler

	echoServer.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
			//Output:os.Stdout,
	}))
	echoServer.Use(middleware.Recover())

	timeoutContext := time.Duration(viper.GetInt("contextTimeout")) * time.Second

	userRepo := _userRepo.NewMysqlUserRepository(db)
	ut := _userUsecase.NewUserUseCase(userRepo, timeoutContext)

	_userHttpDeliver.NewUserHandler(echoServer, ut)

	errCh := make(chan error)

	go func(ch chan error) {
		log.Println("Starting HTTP serverss")
		errCh <- echoServer.Start(":9090")
	}(errCh)

	go func(ch chan error) {
		log.Println("Starting Profiling HTTP server")
		errCh <- http.ListenAndServe(":8080", nil)
	}(errCh)

	for {
		log.Fatal(<-errCh)
	}
}