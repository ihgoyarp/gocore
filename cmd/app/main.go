package main

import (
	"log"
	"net/http"

	"goc/internal/config"
	"goc/internal/middleware"
	"goc/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.New()

	handlerChain := middleware.Recover(
		middleware.Logger(r),
	)

	log.Printf("%s running on :%s\n", cfg.AppName, cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handlerChain))
}
