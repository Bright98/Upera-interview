package domain

type DomainService struct {
	Repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *DomainService {
	return &DomainService{Repo: repo}
}

func (d *DomainService) InsertRevisionService(revision *Revisions) (string, *Errors) {
	revision.ID = GenerateID()
	revision.UpdatedAt = NowTime()
	// TODO: fill revision_no

	err := d.Repo.InsertRevisionRepository(revision)
	if err != nil {
		return "", err
	}

	return revision.ID, nil
}
func (d *DomainService) GetAllRevisionsOfOneProductService(skip, limit int64, productID string) ([]Revisions, *Errors) {
	return d.Repo.GetAllRevisionsOfOneProductRepository(skip, limit, productID)
}
