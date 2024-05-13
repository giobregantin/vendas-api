package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hsxflowers/vendas-api/config"
	"github.com/hsxflowers/vendas-api/internal/http/router"
	"github.com/rs/cors"
)

const TIMEOUT = 30 * time.Second

// @title			vendas-api
// @contact.name	hsxflowers
// @version			1.0
// @BasePath		/v1
func main() {
	var err error
	config.Envs, err = config.LoadEnvVars()

	if err != nil {
		log.Fatalln("Failed loading env", err)
	}

	h := cors.Default().Handler(router.Handlers(config.Envs))

	err = http.ListenAndServe(":8090", h)
	if err != nil {
		log.Fatal("Error running API: ", err)
		os.Exit(1)
	}
}
