package domain

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"reflect"
	"strconv"
	"time"
)

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
func SetError(key string, message string) *Errors {
	return &Errors{Key: key, Message: message}
}
func GenerateID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
func NowTime() int64 {
	return time.Now().Unix()
}
func GetSkipLimitFromQuery(c *gin.Context) (int64, int64) {
	skip, ok := c.GetQuery("skip")
	if !ok {
		skip = "1"
	}
	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "10"
	}
	_skip, err := strconv.ParseInt(skip, 10, 64)
	if err != nil {
		_skip = 0
	}
	_limit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_limit = 0
	}

	return _skip, _limit
}

// service helper
func GetAllProductAttributeKeys(attr ProductAttributes) []string {
	val := reflect.ValueOf(attr)
	var attributeKeys []string
	fmt.Println(val.Type().NumField())
	for i := 0; i < val.Type().NumField(); i++ {
		attributeKeys = append(attributeKeys, val.Type().Field(i).Tag.Get("json"))
	}
	return attributeKeys
}
func ExtractAttributesFromProduct(product *Products) (*ProductAttributes, *Errors) {
	_product, err := json.Marshal(product)
	if err != nil {
		return nil, SetError(InvalidationErr, err.Error())
	}

	attributes := &ProductAttributes{}
	err = json.Unmarshal(_product, attributes)
	if err != nil {
		return nil, SetError(InvalidationErr, err.Error())
	}

	return attributes, nil
}
func GetDifferentKeysBetweenTwoStructs(a, b interface{}) []string {
	var differentKeys []string

	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)
	if valA.Kind() != reflect.Struct || valB.Kind() != reflect.Struct {
		return differentKeys
	}

	typeA := valA.Type()
	typeB := valB.Type()
	if typeA != typeB {
		return differentKeys
	}

	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Field(i)
		fieldB := valB.Field(i)
		if fieldA.Interface() != fieldB.Interface() {
			differentKeys = append(differentKeys, typeA.Field(i).Tag.Get("json"))
		}
	}

	return differentKeys
}
func FillProductByNewAttributes(product Products, attributes *ProductAttributes) (*Products, *Errors) {
	_attributes, err := json.Marshal(attributes)
	if err != nil {
		return nil, SetError(ServiceUnknownErr, err.Error())
	}
	err = json.Unmarshal(_attributes, &product)
	if err != nil {
		return nil, SetError(ServiceUnknownErr, err.Error())
	}
	return &product, nil
}
func (d *DomainService) SendRevisionMessage(revision *Revisions) *Errors {
	revisionBytes, _err := json.Marshal(revision)
	if _err != nil {
		return SetError(InvalidationErr, _err.Error())
	}
	err := d.Repo.PublishMessageRedis("insert-revision", revisionBytes)
	if err != nil {
		return err
	}
	return nil
}
