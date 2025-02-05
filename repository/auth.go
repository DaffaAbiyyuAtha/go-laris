package repository

import (
	"context"
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"

	"github.com/jackc/pgx/v5"
)

func CreateUser(newUser dtos.User) dtos.User {
	db := lib.DB()
	defer db.Close(context.Background())

	newUser.Password = lib.Encrypt(newUser.Password)

	sql := `insert into "user" ("email","password","role_id","profile_id") values ($1,$2,$3,$4) RETURNING "id","email","password","role_id","profile_id"`
	row := db.QueryRow(context.Background(), sql, newUser.Email, newUser.Password, newUser.RoleId, newUser.ProfileId)
	var results dtos.User
	row.Scan(&results.Id, &results.Email, &results.Password, &results.RoleId, &results.ProfileId)
	return results
}

func FindOneUserByEmail(email string) dtos.User {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(
		context.Background(),
		`select * from "user" where "email" = $1`,
		email,
	)

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
