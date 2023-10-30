package api

import "revision/domain"

type RestHandler struct {
	ServiceInterface domain.ServiceInterface
}

func NewRestApi(serviceInterface domain.ServiceInterface) *RestHandler {
	return &RestHandler{ServiceInterface: serviceInterface}
}
