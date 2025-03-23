package repository

import (
	"context"
	"fmt"
	"go-laris/dtos"
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

// func CreateWishlist(productId int, id int) error {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	var exists bool
// 	err := db.QueryRow(
// 		context.Background(),
// 		` SELECT EXISTS (SELECT 1 FROM "wishlist" WHERE user_id = $1 AND product_id = $2)`,
// 		id, productId,
// 	).Scan(&exists)
// 	if err != nil {
// 		return fmt.Errorf("failed to check existing wishlist entry:%w", err)
// 	}

// 	rows, err := db.Query(
// 		context.Background(),
// 		`INSERT INTO "wishlist" (user_id, product_id) VALUES ($1, $2) RETURNING user_id`,
// 		id, productId,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("failed to insert wishlist entry: %w", err)
// 	}

// 	defer rows.Close()
// 	return nil
// }

// func DeleteWishlist(userId int, productId int) error {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	results, err := db.Exec(
// 		context.Background(),
// 		`DELETE FROM "wishlist" WHERE "user_id" = $1 AND "product_id" = $2`,
// 		userId, productId,
// 	)

// 	if err != nil {
// 		return fmt.Errorf("failed to delete wishlist item:%w", err)
// 	}

// 	if results.RowsAffected() == 0 {
// 		return fmt.Errorf("wishlist item not found")
// 	}
// 	return nil
// }

func FindWishlistByProfileId(id int) ([]dtos.Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(),
		`SELECT 
			w.id,
			w.profile_id,
			w.product_id,
			pt.name_product,
			c.name_categories,
			ARRAY_AGG(pi.image) AS images
		FROM 
			wishlist w
		JOIN 
			profile p ON w.profile_id = p.id
		JOIN 
			product pt ON w.product_id = pt.id
		JOIN 
			category c ON c.id = pt.categories_id
		JOIN 
			product_images pi ON pi.product_id = pt.id
		WHERE 
			p.id = $1
		GROUP BY 
			w.id, w.profile_id, w.product_id, pt.name_product, c.name_categories`,
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to find wishlist: %w", err)
	}
	defer rows.Close()

	var wishlists []dtos.Wishlist
	for rows.Next() {
		var wishlist dtos.Wishlist
		err := rows.Scan(
			&wishlist.Id,
			&wishlist.ProfileId,
			&wishlist.ProductId,
			&wishlist.NameProduct,
			&wishlist.NameCategory,
			&wishlist.Image,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan wishlist: %w", err)
		}
		wishlists = append(wishlists, wishlist)
	}

	if len(wishlists) == 0 {
		return nil, fmt.Errorf("wishlist not found")
	}

	return wishlists, nil
}

func CreateWishlist(profileId, productId int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	var exists bool
	err := db.QueryRow(context.Background(),
		`SELECT EXISTS (
			SELECT 1 FROM wishlist 
			WHERE profile_id = $1 AND product_id = $2
		)`, profileId, productId,
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check existing wishlist entry: %w", err)
	}

	if exists {
		return fmt.Errorf("wishlist already exists")
	}

	var id int
	err = db.QueryRow(context.Background(),
		`INSERT INTO wishlist (profile_id, product_id) 
		 VALUES ($1, $2) 
		 RETURNING id`,
		profileId, productId,
	).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to create wishlist: %w", err)
	}

	return nil
}

func DeleteWishlist(profileId, productId int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	_, err := db.Exec(context.Background(),
		"DELETE FROM wishlist WHERE profile_id = $1 AND product_id = $2",
		profileId, productId,
	)

	if err != nil {
		return fmt.Errorf("failed to delete wishlist: %w", err)
	}
	return nil
}

func GetWishlistByProfileAndProductName(profileId int, productName string) ([]dtos.Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	query := `
		SELECT 
			w.id,
			w.profile_id,
			w.product_id,
			pt.name_product,
			c.name_categories,
			ARRAY_AGG(pi.image) AS images
		FROM 
			wishlist w
		JOIN 
			profile p ON w.profile_id = p.id
		JOIN 
			product pt ON w.product_id = pt.id
		JOIN 
			category c ON c.id = pt.categories_id
		JOIN 
			product_images pi ON pi.product_id = pt.id
		WHERE 
			w.profile_id = $1 AND pt.name_product ILIKE $2
		GROUP BY 
			w.id, w.profile_id, w.product_id, pt.name_product, c.name_categories;
	`

	rows, err := db.Query(context.Background(), query, profileId, "%"+productName+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to get wishlist: %w", err)
	}
	defer rows.Close()

	var wishlist []dtos.Wishlist
	for rows.Next() {
		var item dtos.Wishlist
		if err := rows.Scan(&item.Id, &item.ProfileId, &item.ProductId, &item.NameProduct, &item.NameCategory, &item.Image); err != nil {
			return nil, fmt.Errorf("failed to scan wishlist: %w", err)
		}
		wishlist = append(wishlist, item)
	}

	return wishlist, nil
}
