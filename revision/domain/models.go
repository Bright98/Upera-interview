package domain

type Revisions struct {
	ID                 string             `json:"id" bson:"_id"`
	ProductID          string             `json:"Product_id" bson:"product_id"`
	RevisionNo         int                `json:"revision_no" bson:"revision_no"`
	UpdatedAttributes  []string           `json:"updated_attributes" bson:"updated_attributes"`
	PreviousAttributes *ProductAttributes `json:"previous_attributes" bson:"previous_attributes"`
	NewAttributes      *ProductAttributes `json:"new_attributes" bson:"new_attributes"`
	UpdatedAt          int64              `json:"updated_at" bson:"updated_at"`
}
type ProductAttributes struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Color       string `json:"color" bson:"color"`
	Price       int64  `json:"price" bson:"price"`
	ImageUrl    string `json:"image_url" bson:"image_url"`
}

// error type
type Errors struct {
	Key     string `json:"key" bson:"key"`
	Message string `json:"message" bson:"message"`
}
