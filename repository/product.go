package repository

import (
	"context"
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/models"
	"log"

	"github.com/jackc/pgx/v5"
)

func FindOneProductById(id int) (dtos.Product, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var product dtos.Product
	err := db.QueryRow(context.Background(),
		`SELECT id, image,name_product,price,discount,categories_id 
         FROM "product" WHERE id = $1`, id,
	).Scan(&product.Id, &product.Image, &product.NameProduct, &product.Price, &product.Discount, &product.CategoriesId)

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

	sql := `INSERT INTO "product" ("image", "name_product", "price","discount","categories_id") VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	row, err := db.Query(context.Background(), sql, data.Image, data.NameProduct, data.Price, data.Discount, data.CategoriesId, id)
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

func SeeAllProduct(page int, limit int) ([]dtos.Product, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * limit

	sql := `SELECT 
				p.id, 
				p.name_product, 
				p.price, 
				p.discount, 
				p.description, 
				p.categories_id, 
				COALESCE(ARRAY_AGG(pi.image) FILTER (WHERE pi.image IS NOT NULL), '{}') AS images
			FROM product p
			JOIN product_images pi 
			ON p.id = pi.product_id
			GROUP BY p.id
			ORDER BY p.id DESC
		limit $1 offset $2`

	rows, err := db.Query(context.Background(), sql, limit, offset)

	if err != nil {
		return []dtos.Product{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.Product])

	if err != nil {
		log.Println("Error saat mapping data:", err)
		return []dtos.Product{}, err
	}

	return products, err
}

func FindOneProduct(id int) models.Product {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`SELECT 
			p.id, 
			p.name_product, 
			p.price, 
			p.discount, 
			p.description, 
			p.categories_id, 
			c.name_categories,
			COALESCE(ARRAY_AGG(pi.image) FILTER (WHERE pi.image IS NOT NULL), '{}') AS images
		FROM product p
		JOIN product_images pi 
		ON p.id = pi.product_id
		JOIN category c
		ON c.id = p.categories_id
		GROUP BY p.id, p.name_product, p.price, p.discount, p.description, p.categories_id, c.name_categories		
		`,
	)

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Product])

	if err != nil {
		fmt.Println(err)
	}

	product := models.Product{}
	for _, v := range products {
		if v.Id == id {
			product = v
		}
	}
	fmt.Println(product)

	return product
}
