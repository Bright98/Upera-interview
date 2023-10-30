package domain

type RepositoryInterface interface {
	InsertRevisionRepository(revision *Revisions) *Errors
	GetRevisionByIDRepository(id string) (*Revisions, *Errors)
	GetRevisionByProductIDAndNoRepository(productID string, revisionNo int) (*Revisions, *Errors)
	GetAllRevisionsOfOneProductRepository(skip, limit int64, productID string) ([]Revisions, *Errors)
}
