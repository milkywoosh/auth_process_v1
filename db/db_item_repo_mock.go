package dbora

import (
	"database/sql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

type ItemRepositoryMock struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

func NewItemRepositoryMock() (*ItemRepositoryMock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	return &ItemRepositoryMock{DB: db, Mock: mock}, nil
}

/*
func NewItemRepository(db *sql.DB) *DBItemRepository {
	return &DBItemRepository{Conn: db}
}

*/
