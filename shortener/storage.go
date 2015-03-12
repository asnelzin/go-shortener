package shortener

type Storage interface {
	CreateRecord(url string) (string, error)
	GetUrl(hash string) (string, error)
}
