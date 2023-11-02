package test

import (
	"github.com/joho/godotenv"
	"os"
	"product/domain"
	"product/repository"
	"strconv"
	"testing"
)

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
	err = repository.RedisConnection(redisAddress, redisPassword, dbInt)
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
	insertedID := test.InsertProductTest()
	test.UpdateProductTest(insertedID)
	test.GetProductByIDTest(insertedID)
	test.GetAllProductsTest()
}

func (test *testStructure) InsertProductTest() string {
	product := &domain.Products{
		Name:        "product-test",
		Description: "test description for product",
		Color:       "yellow",
		Price:       1000,
		ImageUrl:    "https://upera-interview.com/test-image-url",
	}

	insertedID, err := test.Service.InsertProductService(product)
	if err != nil {
		test.Test.Fatalf("-> [insert product failed]: %v", err)
		return ""
	}

	test.Test.Logf("-> [insert product successfully]: %v", insertedID)
	return insertedID
}
func (test *testStructure) UpdateProductTest(productID string) {
	product := &domain.ProductAttributes{
		Name:        "product-test",
		Description: "test description for product",
		Color:       "yellow",
		Price:       1000,
		ImageUrl:    "https://upera-interview.com/test-image-url",
	}

	err := test.Service.UpdateProductService(productID, product)
	if err != nil {
		test.Test.Fatalf("-> [update product failed]: %v", err)
		return
	}

	test.Test.Logf("-> [update product successfully]")
}
func (test *testStructure) GetProductByIDTest(productID string) {
	product, err := test.Service.GetProductByIDService(productID)
	if err != nil {
		test.Test.Fatalf("-> [update product failed]: %v", err)
		return
	}

	test.Test.Logf("-> [get product successfully]: %v", product)
}
func (test *testStructure) GetAllProductsTest() {
	products, err := test.Service.GetAllProductsService(1, 10)
	if err != nil {
		test.Test.Fatalf("-> [update product failed]: %v", err)
		return
	}

	test.Test.Logf("-> [get all product successfully]: %v", products)
}
