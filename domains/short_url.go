package domains

type ShortUrl struct {
	Url string
}

type UUID struct {
	UUID string
}

type ShortUrlRepository interface {
	FindShortUrl(UUID UUID) (ShortUrl, error)
}
