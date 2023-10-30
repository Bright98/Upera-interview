package domain

type ServiceInterface interface {
	InsertRevisionService(revision *Revisions) (string, *Errors)
	GetRevisionByProductIDAndNoService(productID, revisionNo string) (*Products, *Errors)
	GetAllRevisionsOfOneProductService(skip, limit int64, productID string) ([]Revisions, *Errors)
}
