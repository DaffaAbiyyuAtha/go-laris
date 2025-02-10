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
