package controller

import (
	"log"
	"net/http"
	"short-url-go/domains"
	"short-url-go/usecases"
	"strings"
)

type ShortUrlController struct {
	usecases usecases.ShortUrlUsecase
}

func NewShortUrlController(u usecases.ShortUrlUsecase) ShortUrlController {
	return ShortUrlController{usecases: u}
}

func (c *ShortUrlController) GetRequest(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")
	uuid := domains.UUID{UUID: id}

	log.Println("Received ID. ID=" + id)

	shortUrl, err := c.usecases.GetShortURL(uuid)
	if err != nil {
		log.Println(err)
	}

	log.Println("Redirect to " + shortUrl.Url)
	http.Redirect(w, r, shortUrl.Url, 301)
}
