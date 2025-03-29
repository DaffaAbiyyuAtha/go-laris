package repository

import (
	"context"
	"fmt"
	"go-laris/dtos"
	"go-laris/lib"
	"go-laris/models"

	"github.com/jackc/pgx/v5"
)

func CreateUser(joinRegist dtos.JoinRegist) (*dtos.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	joinRegist.Password, _ = lib.Encrypt(joinRegist.Password)

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
		`INSERT INTO "profile" ("picture", "fullname", "province", "city", "postal_code","gender", "country", "mobile", "address","user_id")VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9, $10) RETURNING id, picture, fullname, province,city,postal_code,gender,country,mobile,address,user_id`,
		joinRegist.Results.Picture, joinRegist.Results.FullName, joinRegist.Results.Province, joinRegist.Results.City, joinRegist.Results.PostalCode, joinRegist.Results.Gender, joinRegist.Results.Country, joinRegist.Results.Mobile, joinRegist.Results.Address, userId,
	).Scan(
		&profile.Id, &profile.Picture, &profile.FullName, &profile.Province,
		&profile.City, &profile.PostalCode, &profile.Gender, &profile.Country, &profile.Mobile, &profile.Address, &profile.UserId,
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

func FindOneProfile(id int) (dtos.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(),
		`select * from "profile" where "user_id" = $1`, id,
	)
	if err != nil {
		return dtos.Profile{}, err
	}
	profile, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.Profile])
	if err != nil {
		return dtos.Profile{}, err
	}
	return profile, nil
}

func FindOneUser(id int) dtos.User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`SELECT * FROM "user" ORDER BY "id" DESC`,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.User])

	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	}

	user := dtos.User{}
	for _, v := range users {
		if v.Id == id {
			user = v
		}
	}
	return user
}

func UpdateProfile(data dtos.Profile, id int) (dtos.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE "profile" 
	SET ("picture", "fullname", "province", "city", "postal_code","gender", "country", "mobile", "address")  =
	($1,$2, $3, $4, $5, $6,$7, $8, $9)
	WHERE "user_id" = $10
	RETURNING *`
	fmt.Printf("Executing SQL with values: %+v\n", data)
	row := db.QueryRow(context.Background(), sql,
		data.Picture, data.FullName, data.Province, data.City,
		data.PostalCode, data.Gender, data.Country, data.Mobile, data.Address, id,
	)

	var profile dtos.Profile
	err := row.Scan(
		&profile.Id, &profile.Picture, &profile.FullName, &profile.Province,
		&profile.City, &profile.PostalCode, &profile.Gender, &profile.Country, &profile.Mobile, &profile.Address, &profile.UserId,
	)
	if err != nil {
		return dtos.Profile{}, fmt.Errorf("failed to update profile: %v", err)
	}

	return profile, nil
}

func FindUser(id int) (models.User, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM "user" WHERE id = $1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.User{}, err
	}

	users, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[models.User])

	if err != nil {
		return models.User{}, err
	}

	return users, err
}

func FindOneUserByEmailForRegist(email string) (dtos.UserRegist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT u.id, u.email, u.password, u.role_id, p.fullname
		FROM "user" u
		LEFT JOIN "profile" p ON u.id = p.user_id
		WHERE u.email = $1
		LIMIT 1
	`

	row := db.QueryRow(context.Background(), sql, email)

	var user dtos.UserRegist
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.RoleId,
		&user.FullName,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return dtos.UserRegist{}, nil
		}
		return dtos.UserRegist{}, err
	}

	return user, nil
}

func FindUsersByRoleforOwner() ([]models.ManagerOwner, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT 
			u.id AS user_id, 
			u.email, 
			p.fullname, 
			r.name AS role_name
		FROM "user" u
		LEFT JOIN "profile" p ON u.id = p.user_id
		LEFT JOIN "role" r ON u.role_id = r.id
		WHERE r.name IN ('admin', 'user')
	`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.ManagerOwner

	for rows.Next() {
		var user models.ManagerOwner
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FullName,
			&user.RoleName,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FindManageUsersByFullName(fullName string) ([]models.ManagerOwner, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT 
			u.id AS user_id, 
			u.email, 
			p.fullname, 
			r.name AS role_name
		FROM "user" u
		LEFT JOIN "profile" p ON u.id = p.user_id
		LEFT JOIN "role" r ON u.role_id = r.id
		WHERE r.name IN ('admin', 'user') 
		AND p.fullname ILIKE '%' || $1 || '%'
	`

	rows, err := db.Query(context.Background(), sql, fullName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.ManagerOwner

	for rows.Next() {
		var user models.ManagerOwner
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FullName,
			&user.RoleName,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func DeleteUserforOwner(userID int) ([]dtos.Profile, error) {
	db := lib.DB()

	tx, err := db.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	sqlGetProfiles := `
		SELECT id, picture, fullname, province, city, postal_code, gender, country, mobile, address, user_id
		FROM "profile"
		WHERE user_id = $1
	`
	rows, err := tx.Query(context.Background(), sqlGetProfiles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []dtos.Profile
	for rows.Next() {
		var profile dtos.Profile
		err := rows.Scan(
			&profile.Id, &profile.Picture, &profile.FullName, &profile.Province,
			&profile.City, &profile.PostalCode, &profile.Gender, &profile.Country,
			&profile.Mobile, &profile.Address, &profile.UserId,
		)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	sqlDeleteWishlist := `DELETE FROM "wishlist" WHERE profile_id IN (SELECT id FROM "profile" WHERE user_id = $1) RETURNING id`
	var deletedWishlistID int
	_ = tx.QueryRow(context.Background(), sqlDeleteWishlist, userID).Scan(&deletedWishlistID)

	sqlDeleteUser := `DELETE FROM "user" WHERE id = $1 RETURNING id`
	var deletedUserID int
	err = tx.QueryRow(context.Background(), sqlDeleteUser, userID).Scan(&deletedUserID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func FindUsersByRoleforAdmin() ([]models.ManagerOwner, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT 
			u.id AS user_id, 
			u.email, 
			p.fullname, 
			r.name AS role_name
		FROM "user" u
		LEFT JOIN "profile" p ON u.id = p.user_id
		LEFT JOIN "role" r ON u.role_id = r.id
		WHERE r.name = 'user'
	`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.ManagerOwner

	for rows.Next() {
		var user models.ManagerOwner
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FullName,
			&user.RoleName,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FindManageUsersByFullNamefoAdmin(fullName string) ([]models.ManagerOwner, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT 
			u.id AS user_id, 
			u.email, 
			p.fullname, 
			r.name AS role_name
		FROM "user" u
		LEFT JOIN "profile" p ON u.id = p.user_id
		LEFT JOIN "role" r ON u.role_id = r.id
		WHERE r.name IN ('user') 
		AND p.fullname ILIKE '%' || $1 || '%'
	`

	rows, err := db.Query(context.Background(), sql, fullName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.ManagerOwner

	for rows.Next() {
		var user models.ManagerOwner
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FullName,
			&user.RoleName,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
