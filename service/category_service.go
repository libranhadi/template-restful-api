package service
import (
	"context"
	"template-restful-api/model/web"
)


type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest, response web.CategoryResponse) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest, response web.CategoryResponse) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	Get(ctx context.Context, response web.CategoryResponse)[] web.CategoryResponse
	FindById(ctx context.Context, categoryId int, response web.CategoryResponse) web.CategoryResponse
}