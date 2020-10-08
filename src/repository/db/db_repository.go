package db

import (
	"github.com/garyjdn/bookstore_oauth-api/src/domain/access_token"
	"github.com/garyjdn/bookstore_oauth-api/src/utils/errors"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.InternalServerError("database connection not implemented yet")
}
