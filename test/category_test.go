package test

import (
	"testing"
	
	"template-restful-api/repository"
	"template-restful-api/service"

	"github.com/stretchr/testify/assert"
)

func TestCategoryFindByIdNotFound(t *testing.T) {
	categoryRepoMock := new(repository.CategoryRepositoryMock)
	categoryServiceMock := service.CategoryServiceMockImpl{CategoryRepositoryMock: categoryRepoMock}
	
	categoryRepoMock.Mock.On("FindById", "1").Return(nil, nil)
	category := categoryServiceMock.FindById("1")
	assert.Equal(t, 404, category.Code);
	assert.Equal(t, "NOT FOUND", category.Status);
	assert.Nil(t, category.Data);
}