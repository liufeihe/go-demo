package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"` // 图书ISBN
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

type Store interface {
	Create(*Book) error
	Update(*Book) error
	Get(string) (Book, error)
	GetAll() ([]Book, error)
	Delete(string) error
}
