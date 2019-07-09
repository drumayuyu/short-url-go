package usecases

import (
	"github.com/pkg/errors"
	"short-url-go/domains"
)

type ShortUrlUsecase struct {
	shortUrlRepository domains.ShortUrlRepository
}

func NewShortUrlUsecase(r domains.ShortUrlRepository) ShortUrlUsecase {
	return ShortUrlUsecase{shortUrlRepository: r}
}

func (u *ShortUrlUsecase) GetShortURL(ID domains.UUID) (domains.ShortUrl, error) {
	shortUrl, err := u.shortUrlRepository.FindShortUrl(ID)
	if err != nil {
		return domains.ShortUrl{}, errors.Wrap(err, "An Error Occurred Searching ShortUrl. UUID="+ID.UUID)
	}
	return shortUrl, nil
}
