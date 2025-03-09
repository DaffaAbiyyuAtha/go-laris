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
