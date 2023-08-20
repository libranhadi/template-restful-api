package repository

import(
	"errors"
	"github.com/stretchr/testify/mock"


	"template-restful-api/model/domain"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (repoMock *CategoryRepositoryMock) FindById(id string) (*domain.Category, error){
	args := repoMock.Mock.Called(id);
	if args.Get(0) == nil {
		return nil, errors.New("Category Not Found !");
	}
	category := args.Get(0).(domain.Category);
	return &category, nil
}