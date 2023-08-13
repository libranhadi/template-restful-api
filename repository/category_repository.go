package repository

import(
	"context"
	"database/sql"
	"template-restful-api/model/domain"
)

type CategoryRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, Id int) (domain.Category, error)
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
}