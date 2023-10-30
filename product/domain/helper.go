package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"reflect"
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

// service helper
func GetAllProductAttributeKeys(attr *ProductAttributes) []string {
	val := reflect.ValueOf(attr)
	var attributeKeys []string
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
func FillProductByNewAttributes(product *Products, attributes *ProductAttributes) (*Products, *Errors) {
	_attributes, err := json.Marshal(attributes)
	if err != nil {
		return nil, SetError(ServiceUnknownErr, err.Error())
	}
	err = json.Unmarshal(_attributes, product)
	if err != nil {
		return nil, SetError(ServiceUnknownErr, err.Error())
	}
	return product, nil
}
