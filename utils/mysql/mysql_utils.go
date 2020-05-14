package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/cookem1/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "sql: no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No records matching given ID")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error parsing DB response: %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062: // Dup Entry error
		return errors.NewBadRequestError(fmt.Sprintf("Duplicate entry: %s", sqlErr.Error()))
	}
	return errors.NewInternalServerError(fmt.Sprintf("error processing request: %s", sqlErr.Message))
}
