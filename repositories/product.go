package repositories

type product struct {
	ID      int
	Name    string
	Quality int
}

type ProductRepository interface {
	GetProducts() (products []product, err error)
}
