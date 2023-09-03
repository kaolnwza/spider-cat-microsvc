package main

import (
	"strings"

	"github.com/IBM/sarama"
	handler "github.com/kaolnwza/spider-cat-microsvc/cat-service/handlers"
	"github.com/kaolnwza/spider-cat-microsvc/cat-service/publisher"
	repository "github.com/kaolnwza/spider-cat-microsvc/cat-service/repositories"
	service "github.com/kaolnwza/spider-cat-microsvc/cat-service/services"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {

	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	e := echo.New()

	healthHdr := handler.NewHealthHandler()

	catRepo := repository.NewCatRepository()
	catSvc := service.NewCatService(catRepo)
	catHdr := handler.NewCatHandler(catSvc)

	eventProd := publisher.NewEventProducer(producer)
	spiderProd := publisher.NewSpiderProducer(eventProd)
	spiderSvc := service.NewSpiderService(spiderProd)
	spiderHdr := handler.NewSpiderHandler(spiderSvc)

	e.GET("/health", healthHdr.HealthCheck)
	e.GET("/", catHdr.GetCatHPHandler)
	e.POST("/spiders/:uuid", spiderHdr.AttackSpiderHandler)

	e.Logger.Fatal(e.Start(":1111"))
}
