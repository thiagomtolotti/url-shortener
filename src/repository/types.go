package repository

type Repository interface {
	GetURL(id string) (string, error)
	CreateURL(url string, id string) error
	Exists(id string) (bool, error)
}
