package domain

type ServiceInterface interface {
	InsertProductService(product *Products) (string, *Errors)
	UpdateProductService(id string, productAttr *ProductAttributes) *Errors
	GetProductByIDService(id string) (*Products, *Errors)
	GetAllProductsService(skip, limit int64) ([]Products, *Errors)
}
