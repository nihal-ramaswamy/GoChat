package db

import (
	"database/sql"
	"errors"

	"github.com/nihal-ramaswamy/GoVid/internal/dto"
)

func insertIntoConference(db *sql.DB, conference *dto.Conference) error {
	if nil == db {
		return errors.New("Sql Db nil")
	}

	query := `INSERT INTO "CONFERENCE" (CODE, ADMIN, ACTIVE) VALUES ($1, $2, $3)`
	err := db.QueryRow(query, conference.Code, conference.Admin, conference.Active).Err()

	return err
}

func insertIntoUser(db *sql.DB, user *dto.User) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var id string
	query := `INSERT INTO "USER" (NAME, EMAIL, PASSWORD) VALUES ($1, $2, $3) RETURNING ID`
	err := db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)

	return id, err
}

func selectAllFromUserWhereEmailIs(db *sql.DB, email string) (dto.User, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var user dto.User
	query := `SELECT * FROM "USER" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&user)

	if err != nil {
		return user, err
	}

	return user, err
}

func selectPasswordFromUserWhereEmailIDs(db *sql.DB, email string) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}
	var password string
	query := `SELECT PASSWORD FROM "USER" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&password)

	return password, err
}
