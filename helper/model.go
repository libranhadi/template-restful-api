package helper

import(
	"template-restful-api/model/domain"
	"template-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id : category.Id,
		Name : category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var responseCategories []web.CategoryResponse;
	for _, c := range categories {
		responseCategories = append(responseCategories, ToCategoryResponse(c))
	}

	return responseCategories
}

