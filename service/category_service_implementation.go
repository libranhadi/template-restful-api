package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"template-restful-api/repository"
	"template-restful-api/model/web"
	"template-restful-api/model/domain"
	"template-restful-api/helper"
	"template-restful-api/exception"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(CategoryRepository repository.CategoryRepository, DB * sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository : CategoryRepository,
		DB : DB,
		Validate : validate,
	}
}

func (service *CategoryServiceImpl) Get(ctx context.Context)[]web.CategoryResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx);

	categories := service.CategoryRepository.Get(ctx, tx)
	var responseCategories []web.CategoryResponse;
	for _, c := range categories {
		responseCategories = append(responseCategories, helper.ToCategoryResponse(c))
	}

	return responseCategories
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse{
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx);
	category, errFindId := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if errFindId != nil {
		panic(exception.NewErrorNotFound(errFindId.Error()))
	}

	return helper.ToCategoryResponse(category)
}


func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	errValidation := service.Validate.Struct(request)
	helper.PanicIfError(errValidation)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx);

	category := domain.Category{
		Name : request.Name,
	}
	category = service.CategoryRepository.Create(ctx, tx ,category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	errValidation := service.Validate.Struct(request)
	helper.PanicIfError(errValidation)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx);

	category, errFindId := service.CategoryRepository.FindById(ctx, tx ,request.Id)
	if errFindId != nil {
		panic(exception.NewErrorNotFound(errFindId.Error()))
	}

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx);
	category, errFindId := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if errFindId != nil {
		panic(exception.NewErrorNotFound(errFindId.Error()))
	}
	service.CategoryRepository.Delete(ctx, tx, category)
}