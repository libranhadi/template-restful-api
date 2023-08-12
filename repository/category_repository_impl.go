package repository

import (
	"context"
	"database/sql"
	"template-restful-api/model/domain"
	"template-restful-api/helper"
	"errors"
)
type CategoryRepositoryImplementation struct {
	db *sql.DB
}

func (repository *CategoryRepositoryImplementation) Get(ctx context.Context, tx *sql.Tx) []domain.Category  {
	query := "SELECT * FROM category";
	rows, err := tx.QueryContext(ctx,query)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}
	return categories
}

func (repository *CategoryRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (*domain.Category, error)  {
	query := "SELECT * FROM category where id = ?";
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err);
	category := &domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err);
		return category, nil
	} else {
		return category, errors.New("Category Not Found") 
	}
}

func (repository *CategoryRepositoryImplementation) Store(ctx context.Context, tx *sql.Tx, category *domain.Category) *domain.Category  {
	query := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId();
	helper.PanicIfError(err)

	category.Id = int(id);

	return category;
}

func (repository *CategoryRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, category *domain.Category) *domain.Category  {
	query := "UPDATE category SET name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)
	return category;
}

func (repository *CategoryRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, category *domain.Category) {
	query := "DELETE from category where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)
}