package gateway

import (
	"github.com/pkg/errors"
	"log"
	"short-url-go/domains"
	"short-url-go/infrastructure"
)

type ShortUrlGateway struct {
	client infrastructure.RedisClient
}

func NewShortUrlGateway(client infrastructure.RedisClient) ShortUrlGateway {
	return ShortUrlGateway{client: client}
}

const ignoreErrorMsg = "redigo: nil returned"

func (u *ShortUrlGateway) FindShortUrl(UUID domains.UUID) (domains.ShortUrl, error) {

	url, err := u.client.RedisGet(UUID.UUID)
	if err != nil {
		if err.Error() == ignoreErrorMsg {
			return domains.ShortUrl{Url: url}, nil
		}
		return domains.ShortUrl{}, errors.Wrap(err, "An Error Occurred RedisConnection.")
	}

	log.Println("Found URL. URL=" + url)

	return domains.ShortUrl{Url: url}, nil
}
