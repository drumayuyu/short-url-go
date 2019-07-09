package main

import (
	"log"
	"net/http"
	"short-url-go/infrastructure"
	"short-url-go/interfaces/controller"
	"short-url-go/interfaces/gateway"
	"short-url-go/usecases"
)

const port = "3838"

func main() {
	http.HandleFunc("/", initShortUrlController().GetRequest)
	log.Println("Server started. port=" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initShortUrlController() *controller.ShortUrlController {
	redisClient := infrastructure.NewRedisClient()
	repo := gateway.NewShortUrlGateway(redisClient)
	usc := usecases.NewShortUrlUsecase(&repo)
	ctr := controller.NewShortUrlController(usc)
	return &ctr
}
