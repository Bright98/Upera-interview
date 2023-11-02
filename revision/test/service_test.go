package test

import (
	"github.com/joho/godotenv"
	"os"
	"revision/domain"
	"revision/repository"
	"strconv"
	"testing"
)

const ProductID = "123"

type testStructure struct {
	Repo    domain.RepositoryInterface
	Service domain.ServiceInterface
	Test    *testing.T
}

func (test *testStructure) Setup() {
	//load env file
	err := godotenv.Load("../.env")
	if err != nil {
		test.Test.Fatalf("-> [load environments failed]: %v", err)
		return
	}

	//get mongo requirements from env file
	timeout := os.Getenv("MONGO_TIMEOUT")
	mongoUrl := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DATABASE")
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		test.Test.Fatalf("-> [Mongo timeout convert failed]: %v", err)
		return
	}

	//mongo connection
	err = repository.MongoConnection(mongoUrl, database, timeoutInt)
	if err != nil {
		test.Test.Fatalf("-> [Mongo connection failed]: %v", err)
		return
	}

	//get redis requirements from env file
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	db := os.Getenv("REDIS_DB")
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		test.Test.Fatalf("-> [Redis timeout convert failed]: %v", err)
		return
	}

	//redis connection
	_, _, err = repository.RedisConnection(redisAddress, redisPassword, dbInt)
	if err != nil {
		test.Test.Fatalf("-> [Redis connection failed]: %v", err)
		return
	}

	test.Repo = repository.NewRepository()
	test.Service = domain.NewService(test.Repo)
}

func TestServiceOperations(t *testing.T) {
	test := testStructure{Test: t}
	test.Setup()

	// functions
	test.InsertRevisionTest()
	test.GetRevisionByProductIDAndNoTest()
	test.GetAllRevisionsOfOneProductTest()
}

func (test *testStructure) InsertRevisionTest() {
	prevProduct := &domain.Products{
		ID:            "1",
		Name:          "product-test",
		Description:   "test description for product",
		Color:         "yellow",
		Price:         1000,
		ImageUrl:      "https://upera-interview.com/test-image-url",
		LastUpdatedAt: 1698936984,
		CreatedAt:     1698936984,
		Status:        "active",
	}
	newProduct := &domain.Products{
		ID:            "1",
		Name:          "product-test",
		Description:   "test description for product",
		Color:         "red",
		Price:         1000,
		ImageUrl:      "https://upera-interview.com/test-image-url",
		LastUpdatedAt: 1698936984,
		CreatedAt:     1698936984,
		Status:        "active",
	}
	revision := &domain.Revisions{
		ProductID:         ProductID,
		UpdatedAttributes: []string{"color"},
		PreviousProduct:   prevProduct,
		NewProduct:        newProduct,
	}

	insertedID, err := test.Service.InsertRevisionService(revision)
	if err != nil {
		test.Test.Fatalf("-> [insert product failed]: %v", err)
	}

	test.Test.Logf("-> [insert product successfully]: %v", insertedID)
}
func (test *testStructure) GetRevisionByProductIDAndNoTest() {
	product, err := test.Service.GetRevisionByProductIDAndNoService(ProductID, "0")
	if err != nil {
		test.Test.Fatalf("-> [update product failed]: %v", err)
		return
	}

	test.Test.Logf("-> [get product successfully]: %v", product)
}
func (test *testStructure) GetAllRevisionsOfOneProductTest() {
	revisions, err := test.Service.GetAllRevisionsOfOneProductService(1, 10, ProductID)
	if err != nil {
		test.Test.Fatalf("-> [update product failed]: %v", err)
		return
	}

	test.Test.Logf("-> [get all product successfully]: %v", revisions)
}
