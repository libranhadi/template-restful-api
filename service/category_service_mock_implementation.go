package service


import(
	"template-restful-api/model/web"
	"template-restful-api/repository"
	"template-restful-api/helper"
)

type CategoryServiceMockImpl struct {
	CategoryRepositoryMock *repository.CategoryRepositoryMock
}

func (repoMock *CategoryServiceMockImpl) FindById(id string) (*web.WebResponse){
	category, err := repoMock.CategoryRepositoryMock.FindById(id);
	result := helper.CreateResponse(200, "OK", category)
	if err != nil {
		result = helper.CreateResponse(404, "NOT FOUND", category)
	}
	
	return result;
}