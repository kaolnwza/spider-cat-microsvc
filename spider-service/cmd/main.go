package main

import (
	"context"
	"strings"

	"events"

	"github.com/IBM/sarama"
	handler "github.com/kaolnwza/spider-cat-microsvc/spider-service/handlers"
	repository "github.com/kaolnwza/spider-cat-microsvc/spider-service/repositories"
	service "github.com/kaolnwza/spider-cat-microsvc/spider-service/services"
	"github.com/kaolnwza/spider-cat-microsvc/spider-service/subscriber"
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

	e := echo.New()

	spiderRepo := repository.NewSpiderRepository()
	spiderSvc := service.NewSpiderService(spiderRepo)
	spiderHdr := handler.NewSpiderHandler(spiderSvc)

	spiderEvent := subscriber.NewSpiderEventHandler(spiderSvc)
	spiderConsumer := subscriber.NewConsumerHandler(spiderEvent)

	e.GET("/", spiderHdr.GetSpidersHandler)

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	go func() {
		for {
			consumer.Consume(context.Background(), events.Topics, spiderConsumer)
		}
	}()

	go e.Logger.Fatal(e.Start(":3333"))

}
