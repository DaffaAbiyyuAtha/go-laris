package repository

import (
	"context"
	"fmt"
	"go-laris/lib"
	"go-laris/models"

	"github.com/jackc/pgx/v5"
)

func FindAllCategories() []models.Categories {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM "category" ORDER BY "id" ASC`
	rows, _ := db.Query(context.Background(), sql)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Categories])

	if err != nil {
		fmt.Println(err)
	}

	return categories
}

func FindAllUsersWithPagination(search string, page int, limit int) ([]models.CategoriesPagination, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * limit

	sql := `SELECT "p"."id", "p"."image", "p"."name_product", "p"."price", "p"."discount", "c"."name_categories"
		FROM "product" "p"
		JOIN "category" "c"
		ON "c"."id" = "p"."categories_id"
		WHERE "c"."name_categories" ilike $1
		ORDER BY "p"."id" DESC
		limit $2 offset $3`

	rows, err := db.Query(context.Background(), sql, "%"+search+"%", limit, offset)

	if err != nil {
		return []models.CategoriesPagination{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.CategoriesPagination])

	if err != nil {
		return []models.CategoriesPagination{}, err
	}

	return users, err
}

func GetFilterProductWithCategory(search string, page int, limit int) ([]models.Product, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * limit

	sql := `
		SELECT 
			p.id, 
			p.name_product, 
			p.price, 
			p.discount, 
			p.description, 
			p.categories_id, 
			c.name_categories,
			COALESCE(ARRAY_AGG(pi.image) FILTER (WHERE pi.image IS NOT NULL), '{}') AS image
		FROM product p
		JOIN product_images pi 
		ON p.id = pi.product_id
		JOIN category c
		ON c.id = p.categories_id
		WHERE c.name_categories ILIKE $1
		GROUP BY p.id, p.name_product, p.price, p.discount, p.description, p.categories_id, c.name_categories
		ORDER BY p.id DESC
		limit $2 offset $3
	`

	rows, err := db.Query(context.Background(), sql, "%"+search+"%", limit, offset)

	if err != nil {
		return []models.Product{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Product])

	if err != nil {
		return []models.Product{}, err
	}

	return products, err
}
