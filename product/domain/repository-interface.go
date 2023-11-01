package domain

type RepositoryInterface interface {
	InsertProductRepository(product *Products) *Errors
	UpdateProductRepository(product *Products) *Errors
	GetProductByIDRepository(id string) (*Products, *Errors)
	GetAllProductsRepository(skip, limit int64) ([]Products, *Errors)
	PublishMessageRedis(channel string, message []byte) *Errors
}
