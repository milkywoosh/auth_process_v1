package dbora

import (
	"context"
	"database/sql"
	"time"

	"oraluke.com/conn-ora1/domain"
)

type DBUserRepository struct {
	Conn DBTX
}

func NewUserRepository(db DBTX) *DBUserRepository {
	return &DBUserRepository{
		Conn: db,
	}
}

// CreateUser(ctx context.Context, username string, password string) error
func (d *DBUserRepository) CreateUser(ctx context.Context, username string, password string) error {
	query := `
		INSERT INTO USERS 
			   (USERNAME , PASSWORD, CREATED_AT)
		VALUES (:USERNAME, :PASSWORD, :CREATED_AT)
	`
	_, err := d.Conn.ExecContext(ctx, query,
		sql.Named("USERNAME", username),
		sql.Named("PASSWORD", password),
		sql.Named("CREATED_AT", time.Now().UTC()),
	)
	if err != nil {
		return err
	}
	return nil
}

type UserRow struct {
	Username  sql.NullString
	Password  sql.NullString
	CreatedAt sql.NullTime
}

func DTOUser(source UserRow) domain.User {
	return domain.User{
		Username:  source.Username.String,
		Password:  source.Password.String,
		CreatedAt: &source.CreatedAt.Time,
	}
}

func (d *DBUserRepository) GetUser(ctx context.Context, username string) (domain.User, error) {

	var user UserRow

	query := `SELECT USERNAME, PASSWORD, CREATED_AT FROM USERS WHERE USERNAME = :USERNAME`

	row := d.Conn.QueryRowContext(ctx, query, sql.Named("USERNAME", username))
	if err := row.Scan(
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, err
		}
		return domain.User{}, err
	}

	return DTOUser(user), nil

}
