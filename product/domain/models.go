package domain

type Products struct {
	ID            string `json:"id" bson:"_id"`
	Name          string `json:"name" bson:"name"`
	Description   string `json:"description" bson:"description"`
	Color         string `json:"color" bson:"color"`
	Price         int64  `json:"price" bson:"price"`
	ImageUrl      string `json:"image_url" bson:"image_url"`
	CreatedAt     int64  `json:"created_at" bson:"created_at"`
	LastUpdatedAt int64  `json:"last_updated_at" bson:"last_updated_at"`
	Status        string `json:"status" bson:"status"` // for handling status of product: active, removed
}
type ProductAttributes struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Color       string `json:"color" bson:"color"`
	Price       int64  `json:"price" bson:"price"`
	ImageUrl    string `json:"image_url" bson:"image_url"`
}
type Revisions struct {
	ProductID         string    `json:"Product_id" bson:"product_id"`
	UpdatedAttributes []string  `json:"updated_attributes" bson:"updated_attributes"`
	PreviousProduct   *Products `json:"previous_product" bson:"previous_product"`
	NewProduct        *Products `json:"new_product" bson:"new_product"`
}

// error type
type Errors struct {
	Key     string `json:"key" bson:"key"`
	Message string `json:"message" bson:"message"`
}
