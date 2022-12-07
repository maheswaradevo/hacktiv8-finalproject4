package repository

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type authRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewAuthRepository(db *sql.DB, logger *zap.Logger) *authRepository {
	return &authRepository{
		db:     db,
		logger: logger,
	}
}

var (
	INSERT_USER   = "INSERT INTO users(full_name, email, password, role, balance) VALUES (?, ?, ?, ?, ?);"
	FIND_BY_EMAIL = "SELECT id, email, password FROM users WHERE email = ?;"
)

func (a authRepository) InsertUser(ctx context.Context, data model.User) (userID uint64, err error) {
	query := INSERT_USER

	res, err := a.db.ExecContext(ctx, query, data.FullName, data.Email, data.Password, data.Role, data.Balance)
	if err != nil {
		a.logger.Error("[InsertUser] failed to insert data", zap.Error(err))
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return uint64(lastId), nil
}

func (a authRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := FIND_BY_EMAIL
	rows := a.db.QueryRowContext(ctx, query, email)

	user := &model.User{}

	err := rows.Scan(&user.UserID, &user.Email, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		a.logger.Error("[FindByEmail] failed to scan the data", zap.Error(err))
		return nil, err
	} else if err == sql.ErrNoRows {
		a.logger.Info("[FindByEmail] no data existed")
		return nil, errors.ErrInvalidResources
	}
	return user, nil
}
