package router

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/hsxflowers/vendas-api/config"
	produtosHandler "github.com/hsxflowers/vendas-api/internal/http/produtos" // produtos Handler
	"github.com/hsxflowers/vendas-api/pkg/broker"
	"github.com/hsxflowers/vendas-api/pkg/broker/kafkapkg"
	produtos "github.com/hsxflowers/vendas-api/produtos"
	produtosDatabase "github.com/hsxflowers/vendas-api/produtos/db"
	"github.com/hsxflowers/vendas-api/produtos/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Handlers(envs *config.Environments) *echo.Echo {
	e := echo.New()
	ctx := context.Background()

	var produtosDb domain.ProdutosDatabase
	var err error

	db, err := sql.Open("postgres", "postgres://dbecommerce_n9hn_user:0Scqh1LQR7sKG1EYNBaBmx4VPWirvJtF@dpg-co7upesf7o1s738n3u5g-a.oregon-postgres.render.com/dbecommerce_n9hn?sslmode=require")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	produtosDb = produtosDatabase.NewSQLStore(db)

	log.Debug("")

	consumer, err := kafkapkg.NewKafkaConsumer(&kafkapkg.KafkaConsumerConfig{
		Broker:     envs.KafkaBrokers,
		GroupId:    "envs.KafkaConsumerGroupId",
		AutoOffset: "earliest",
	})
	if err != nil {
		log.Fatal("Failed to create kafka consumer", err)
	}

	produtosProducer, err := kafkapkg.NewKafkaProducer(envs.KafkaBrokers)
	if err != nil {
		log.Fatal("Failed to create webhook kafka producer", err)
	}

	produtosBroker := broker.NewBroker(consumer, produtosProducer)

	produtosConsumer := produtos.NewConsumer(produtosBroker)
	produtosRepository := produtos.NewProdutosRepository(produtosDb)
	produtosService := produtos.NewProdutosService(produtosRepository, produtosConsumer)
	produtosHandler := produtosHandler.NewProdutosHandler(ctx, produtosService)

	produtos := e.Group("produtos")

	produtos.POST("", func(c echo.Context) error {
		result, err := produtosHandler.Create(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	})

	return e
}
