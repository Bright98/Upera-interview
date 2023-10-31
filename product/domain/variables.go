package domain

// collections
const (
	ProductCollection = "products"
)

// status keys
const (
	ProductActiveStatus  = "active"
	ProductRemovedStatus = "removed"
)

// error keys
const (
	CantInsertErr     = "CANT_INSERT"
	CantUpdateErr     = "CANT_UPDATE"
	CantRemoveErr     = "CANT_REMOVE"
	NotFoundErr       = "NOT_FOUND"
	CantCountErr      = "CANT_COUNT"
	ServiceUnknownErr = "SERVICE_UNKNOWN"
	InvalidationErr   = "INVALIDATION"
	CantPublishErr    = "CANT_PUBLISH"
)
