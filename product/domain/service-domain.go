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
	revision.UpdatedAttributes = GetAllProductAttributeKeys(*attributes)
	revision.PreviousProduct = nil
	revision.NewProduct = product
	err = d.SendRevisionMessage(revision)
	if err != nil {
		return "", err
	}

	return product.ID, nil
}
func (d *DomainService) UpdateProductService(id string, productAttr *ProductAttributes) *Errors {
	//get old product
	product, err := d.Repo.GetProductByIDRepository(id)
	if err != nil {
		return err
	}

	//create new product from attributes
	newProduct, err := FillProductByNewAttributes(*product, productAttr)
	if err != nil {
		return err
	}
	newProduct.LastUpdatedAt = NowTime()

	//update product in db
	err = d.Repo.UpdateProductRepository(newProduct)
	if err != nil {
		return err
	}

	//fill revision fields and send
	revision := &Revisions{}
	revision.ProductID = id
	revision.PreviousProduct = product
	revision.NewProduct = newProduct
	revision.UpdatedAttributes = GetDifferentKeysBetweenTwoStructs(*product, *newProduct)
	err = d.SendRevisionMessage(revision)
	if err != nil {
		return err
	}

	return nil
}
func (d *DomainService) GetProductByIDService(id string) (*Products, *Errors) {
	return d.Repo.GetProductByIDRepository(id)
}
func (d *DomainService) GetAllProductsService(skip, limit int64) ([]Products, *Errors) {
	return d.Repo.GetAllProductsRepository(skip, limit)
}
