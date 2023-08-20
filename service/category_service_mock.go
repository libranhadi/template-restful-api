package service

import (
	"template-restful-api/model/domain"
)

type CategoryServiceMock interface {
	FindById(Id string) *domain.Category
}