package domain

import (
	"strconv"
)

type DomainService struct {
	Repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *DomainService {
	return &DomainService{Repo: repo}
}

func (d *DomainService) InsertRevisionService(revision *Revisions) (string, *Errors) {
	//find last revision no from redis or db
	revisionNo, err := d.Repo.GetLastRevisionNoRedis(revision.ProductID)
	if err != nil {
		//find from db
		revisionNo, err = d.Repo.GetLastRevisionNoOfProductRepository(revision.ProductID)
		if err != nil {
			return "", err
		}
	}

	//fill some revision fields
	revision.ID = GenerateID()
	revision.UpdatedAt = NowTime()
	revision.RevisionNo = revisionNo + 1

	//set last revision no in redis
	err = d.Repo.SetLastRevisionNoRedis(revision.ProductID, revision.RevisionNo)
	if err != nil {
		return "", err
	}

	// insert revision in db
	err = d.Repo.InsertRevisionRepository(revision)
	if err != nil {
		return "", err
	}

	return revision.ID, nil
}
func (d *DomainService) GetRevisionByProductIDAndNoService(productID, revisionNo string) (*Products, *Errors) {
	//convert revision no to right type
	revisionNumber, _err := strconv.Atoi(revisionNo)
	if _err != nil {
		return nil, SetError(InvalidationErr, _err.Error())
	}

	revision, err := d.Repo.GetRevisionByProductIDAndNoRepository(productID, revisionNumber)
	if err != nil {
		return nil, err
	}
	return revision.NewProduct, nil
}
func (d *DomainService) GetAllRevisionsOfOneProductService(skip, limit int64, productID string) ([]Revisions, *Errors) {
	return d.Repo.GetAllRevisionsOfOneProductRepository(skip, limit, productID)
}
