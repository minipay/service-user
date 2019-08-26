package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	_ "net/http/pprof"

	_articleHttpDeliver "cleanbase/article/handler"
	_articleRepo "cleanbase/article/repository"
	_articleUcase "cleanbase/article/usecase"
	_authorRepo "cleanbase/author/repository"
	_authorUsecase "cleanbase/author/usecase"
)

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	logrus.Info("Using Config file: ", viper.ConfigFileUsed())

	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("Comment service is Running in Debug Mode")
		return
	}
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Warn("Comment service is Running in Production Mode")
	logrus.SetFormatter(&logrus.JSONFormatter{})
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
		log.Fatal(err)
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	echoServer := echo.New()
	//articleRepo := _articleRepo.NewMysqlArticleRepository(db)
	authorRepo := _authorRepo.NewMysqlAuthorRepository(db)
	// Register the handler
	articleRepo := _articleRepo.NewMysqlArticleRepository(db)
	timeoutContext := time.Duration(viper.GetInt("contextTimeout")) * time.Second
	au := _articleUcase.NewArticleUseCase(articleRepo, authorRepo, timeoutContext)
	at := _authorUsecase.NewAuthorUseCase(authorRepo, timeoutContext)

	_articleHttpDeliver.NewArticleHandler(echoServer, au, at)

	errCh := make(chan error)

	go func(ch chan error) {
		log.Println("Starting HTTP servers")
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