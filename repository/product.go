package repository

import (
	"context"
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"

	"github.com/jackc/pgx/v5"
)

func FindOneProductById(id int) (dtos.Product, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var product dtos.Product
	err := db.QueryRow(context.Background(),
		`SELECT id, image,name_product,price,discount,total,categories_id 
         FROM "product" WHERE id = $1`, id,
	).Scan(&product.Id, &product.Image, &product.NameProduct, &product.Price, &product.Discount, &product.Total, &product.CategoriesId)

	if err != nil {
		return dtos.Product{}, fmt.Errorf("failed to find product: %w", err)
	}

	return product, nil
}

func FindAllProduct(search string, limit int, page int) []dtos.Product {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := (page - 1) * limit

	sql := `SELECT * FROM "product" WHERE "name_product" ILIKE '%' || $1 || '%' ORDER BY "name_product" DESC OFFSET $2 LIMIT $3`
	rows, _ := db.Query(context.Background(), sql, search, offset, limit)
	product, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.Product])
	if err != nil {
		fmt.Println(err)
	}

	return product

}

func CreateProduct(data dtos.Product, id int) (dtos.Product, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "product" ("image", "name_product", "price","discount","total","categories_id") VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	row, err := db.Query(context.Background(), sql, data.Image, data.NameProduct, data.Price, data.Discount, data.Total, data.CategoriesId, id)
	if err != nil {
		return dtos.Product{}, err
	}

	product, err := pgx.CollectOneRow(row, pgx.RowToStructByName[dtos.Product])
	if err != nil {
		return dtos.Product{}, err
	}

	return product, err

}

func DeleteProduct(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	results, err := db.Exec(context.Background(), `DELETE FROM "product" WHERE "id" = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if results.RowsAffected() == 0 {
		return fmt.Errorf("no user found")
	}
	return nil
}
