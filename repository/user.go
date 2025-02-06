package repository

import (
	"context"
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"

	"github.com/jackc/pgx/v5"
)

func CreateUser(joinRegist dtos.JoinRegist) (*dtos.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	joinRegist.Password = lib.Encrypt(joinRegist.Password)

	var userId int
	err := db.QueryRow(
		context.Background(),
		`insert into "user" ("email","password","role_id") values ($1,$2,$3) RETURNING "id"`,
		joinRegist.Email, joinRegist.Password, joinRegist.RoleId,
	).Scan(&userId)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into users table: %v", err)
	}

	profile := dtos.Profile{
		UserId:   userId,
		FullName: joinRegist.Results.FullName,
	}
	err = db.QueryRow(
		context.Background(),
		`INSERT INTO "profile" ("pictrue", "fullname", "province", "city", "postal_code", "country", "mobile", "address","user_id")VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9) RETURNING id, pictrue, fullname, province,city,postal_code,country,mobile,address,user_id`,
		joinRegist.Results.Picture, joinRegist.Results.FullName, joinRegist.Results.Province, joinRegist.Results.City, joinRegist.Results.PostalCode, joinRegist.Results.Country, joinRegist.Results.Mobile, joinRegist.Results.Address, userId,
	).Scan(
		&profile.Id, &profile.Picture, &profile.FullName, &profile.Province,
		&profile.City, &profile.PostalCode, &profile.Country, &profile.Mobile, &profile.Address, &profile.UserId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert into profile table: %v", err)
	}

	return &profile, nil

}

func FindAllUser() []dtos.Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM "profile" ORDER BY "id" ASC`
	rows, _ := db.Query(context.Background(), sql)
	profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.Profile])

	if err != nil {
		fmt.Println(err)
	}

	return profile
}

func FindOneUserByEmail(email string) dtos.User {
	db := lib.DB()
	defer db.Close(context.Background())
	sql := `select * from "user" where "email" = $1`
	rows, _ := db.Query(context.Background(), sql, email)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.User])

	if err != nil {
		fmt.Println(err)
	}

	user := dtos.User{}

	for _, val := range users {
		if val.Email == email {
			user = val
		}
	}
	return user
}
