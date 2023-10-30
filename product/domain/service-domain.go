package domain

type DomainService struct {
	Repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *DomainService {
	return &DomainService{Repo: repo}
}

func (d *DomainService) InsertProductService(product *Products) (string, *Errors) {
	product.ID = GenerateID()
	product.CreatedAt = NowTime()
	product.LastUpdatedAt = NowTime()
	product.Status = ProductActiveStatus

	err := d.Repo.InsertProductRepository(product)
	if err != nil {
		return "", err
	}

	attributes, err := ExtractAttributesFromProduct(product)
	if err != nil {
		return "", err
	}
	revision := &Revisions{}
	revision.ProductID = product.ID
	revision.UpdatedAttributes = GetAllProductAttributeKeys(attributes)
	revision.NewAttributes = attributes
	// TODO: insert first record in revisions

	return product.ID, nil
}
func (d *DomainService) UpdateProductService(id string, productAttr *ProductAttributes) *Errors {
	product, err := d.Repo.GetProductByIDRepository(id)
	if err != nil {
		return err
	}

	oldAttributes, err := ExtractAttributesFromProduct(product)
	if err != nil {
		return err
	}

	product.LastUpdatedAt = NowTime()
	product, err = FillProductByNewAttributes(product, productAttr)
	if err != nil {
		return err
	}

	err = d.Repo.UpdateProductRepository(product)
	if err != nil {
		return err
	}

	// TODO: insert revision
	revision := &Revisions{}
	revision.ProductID = id
	revision.PreviousAttributes = oldAttributes
	revision.NewAttributes = productAttr
	revision.UpdatedAttributes = GetDifferentKeysBetweenTwoStructs(oldAttributes, productAttr)

	return nil
}
func (d *DomainService) GetProductByIDService(id string) (*Products, *Errors) {
	return d.Repo.GetProductByIDRepository(id)
}
func (d *DomainService) GetAllProductsService(skip, limit int64) ([]Products, *Errors) {
	return d.Repo.GetAllProductsRepository(skip, limit)
}
