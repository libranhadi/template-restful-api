package service

import (
	"context"
	"github.com/go-playground/validator"
	"template-restful-api/repository"
	"template-restful-api/model/web"
	"template-restful-api/model/domain"
	"template-restful-api/helper"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
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
	helper.PanicIfError(errFindId)

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

	category = service.CategoryRepository.Store(ctx, tx ,category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	errValidation := service.Validate.Struct(request)
	helper.PanicIfError(errValidation)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx);

	category, errFindId := service.CategoryRepository.FindById(ctx, tx ,request.Id)
	helper.PanicIfError(errFindId)

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx);
	category, errFindId := service.CategoryRepository.FindById(ctx, tx, categoryId)

	helper.PanicIfError(errFindId)
	service.CategoryRepository.Delete(ctx, tx, category)
}