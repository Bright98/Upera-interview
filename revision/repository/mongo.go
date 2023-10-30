package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"revision/domain"
	"time"
)

var (
	MongoTimeout  time.Duration
	MongoDatabase *mongo.Database
	MongoClient   *mongo.Client
)

func (r Repository) InsertRevisionRepository(revision *domain.Revisions) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(domain.RevisionCollection)
	_, err := collection.InsertOne(ctx, revision)
	if err != nil {
		return domain.SetError(domain.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) GetRevisionByIDRepository(id string) (*domain.Revisions, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(domain.RevisionCollection)
	filter := bson.M{"_id": id}
	revision := &domain.Revisions{}
	err := collection.FindOne(ctx, filter).Decode(revision)
	if err != nil {
		return nil, domain.SetError(domain.NotFoundErr, err.Error())
	}
	return revision, nil
}
func (r Repository) GetAllRevisionsOfOneProductRepository(skip, limit int64, productID string) ([]domain.Revisions, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(domain.RevisionCollection)
	filter := bson.M{"product_id": productID}
	if skip != 0 {
		skip = (skip - 1) * limit
	}
	option := options.FindOptions{Skip: &skip, Limit: &limit}
	res, err := collection.Find(ctx, filter, &option)
	if err != nil {
		return nil, domain.SetError(domain.ServiceUnknownErr, err.Error())
	}
	var result []domain.Revisions
	err = res.All(ctx, &result)
	if err != nil {
		return nil, domain.SetError(domain.ServiceUnknownErr, err.Error())
	}
	err = res.Close(ctx)
	if err != nil {
		return nil, domain.SetError(domain.ServiceUnknownErr, err.Error())
	}
	return result, nil
}
