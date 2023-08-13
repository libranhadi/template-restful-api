package service
import (
	"context"
	"template-restful-api/model/web"
)


type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	Get(ctx context.Context)[] web.CategoryResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
}