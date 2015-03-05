package shortener

type Storage interface {
	CreateRecord(url string) string
	GetUrl(hash string) string
}
