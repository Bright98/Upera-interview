package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"product/domain"
	"time"
)

var (
	MongoTimeout  time.Duration
	MongoDatabase *mongo.Database
	MongoClient   *mongo.Client
)

func (r Repository) InsertProductRepository(product *domain.Products) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(domain.ProductCollection)
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return domain.SetError(domain.CantInsertErr, err.Error())
	}
	return nil
}
func (r Repository) UpdateProductRepository(product *domain.Products) *domain.Errors {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(domain.ProductCollection)
	filter := bson.M{"_id": product.ID, "status": domain.ProductActiveStatus}
	res, err := collection.UpdateOne(ctx, filter, bson.D{{"$set", product}})
	if err != nil {
		return domain.SetError(domain.CantUpdateErr, err.Error())
	}
	if res.MatchedCount == 0 {
		return domain.SetError(domain.NotFoundErr, "")
	}
	return nil
}
func (r Repository) GetProductByIDRepository(id string) (*domain.Products, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(domain.ProductCollection)
	filter := bson.M{"_id": id, "status": domain.ProductActiveStatus}
	product := &domain.Products{}
	err := collection.FindOne(ctx, filter).Decode(product)
	if err != nil {
		return nil, domain.SetError(domain.NotFoundErr, err.Error())
	}
	return product, nil
}
func (r Repository) GetAllProductsRepository(skip, limit int64) ([]domain.Products, *domain.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout)
	defer cancel()
	collection := MongoDatabase.Collection(domain.ProductCollection)
	filter := bson.M{"status": domain.ProductActiveStatus}
	if skip != 0 {
		skip = (skip - 1) * limit
	}
	option := options.FindOptions{Skip: &skip, Limit: &limit}
	res, err := collection.Find(ctx, filter, &option)
	if err != nil {
		return nil, domain.SetError(domain.ServiceUnknownErr, err.Error())
	}
	var result []domain.Products
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
