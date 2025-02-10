package repository

import (
	"context"
	"fmt"
	"go-laris/lib"
	"go-laris/models"

	"github.com/jackc/pgx/v5"
)

func FindAllWishlist() []models.Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "wishlist" order by "id" asc`)

	wishlists, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Wishlist])
	if err != nil {
		fmt.Println(err)
	}
	return wishlists
}

func FindOneWishlist(id int) ([]models.Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(), `select * from "wishlist" where "user_id" = $1 order by "id" asc`, id)

	if err != nil {
		return nil, fmt.Errorf("failed to query wishlist:%w", err)
	}

	defer rows.Close()
	wishlist, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Wishlist])

	if err != nil {
		return nil, fmt.Errorf("failed to collect wishlist rows:%w", err)
	}

	return wishlist, nil
}

func CreateWishlist(productId int, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	var exists bool
	err := db.QueryRow(
		context.Background(),
		` SELECT EXISTS (SELECT 1 FROM "wishlist" WHERE user_id = $1 AND product_id = $2)`,
		id, productId,
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check existing wishlist entry:%w", err)
	}

	rows, err := db.Query(
		context.Background(),
		`INSERT INTO "wishlist" (user_id, product_id) VALUES ($1, $2) RETURNING user_id`,
		id, productId,
	)
	if err != nil {
		return fmt.Errorf("failed to insert wishlist entry: %w", err)
	}

	defer rows.Close()
	return nil
}

func DeleteWishlist(userId int, productId int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	results, err := db.Exec(
		context.Background(),
		`DELETE FROM "wishlist" WHERE "user_id" = $1 AND "product_id" = $2`,
		userId, productId,
	)

	if err != nil {
		return fmt.Errorf("failed to delete wishlist item:%w", err)
	}

	if results.RowsAffected() == 0 {
		return fmt.Errorf("wishlist item not found")
	}
	return nil
}
