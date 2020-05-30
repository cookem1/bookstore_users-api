package users

import (
	"fmt"

	"github.com/cookem1/bookstore_users-api/datasources/users_db"
	"github.com/cookem1/bookstore_users-api/domain/users/logger"
	"github.com/cookem1/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail      = "email_UNIQUE"
	errorNoRows           = "sql: no rows in result set"
	queryInsertUser       = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?; "
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?; "
	queryDeleteUser       = "DELETE FROM users WHERE id=?; "
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users where status=?; "
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database Error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when getting get user by ID", err)
		return errors.NewInternalServerError("Database Error")

		//	return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("error when preparing user save statement", err)
		return errors.NewInternalServerError("Database Error")
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when saving user", saveErr)
		return errors.NewInternalServerError("Database Error")
	}

	userID, insErr := insertResult.LastInsertId()
	if insErr != nil {
		logger.Error("error when getting last insert ID", insErr)
		return errors.NewInternalServerError("Database Error")
	}

	user.Id = userID
	return nil
}

func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("error when trying to prepare update User statement", err)
		return errors.NewInternalServerError("Database Error")
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if updateErr != nil {
		logger.Error("error when trying to execure user update statement", updateErr)
		return errors.NewInternalServerError("Database Error")
	}

	return nil

}

func (user *User) Delete() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("error when trying to prepare User delete statement", err)
		return errors.NewInternalServerError("Database Error")
	}

	defer stmt.Close()

	if _, delErr := stmt.Exec(user.Id); delErr != nil {
		logger.Error("error when trying to delete User", delErr)
		return errors.NewInternalServerError("Database Error")
	}
	return nil
}

func (user User) FindByStatus(status string) ([]User, *errors.RestErr) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)

	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.NewInternalServerError("Database Error")
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("error when trying to query user by stats from DB", err)
		return nil, errors.NewInternalServerError("Database Error")
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying go to next row in result set", err)
			return nil, errors.NewInternalServerError("Database Error")
		}

		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
