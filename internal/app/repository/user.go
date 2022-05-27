package repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/iamtonydev/user-service-api/internal/app/model"
	"github.com/iamtonydev/user-service-api/internal/app/repository/table"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user *model.UserInfo) (int64, error)
	GetUser(ctx context.Context, userId int64) (*model.UserInfo, error)
	MultiAddUser(ctx context.Context, users []*model.UserInfo) (int64, error)
	ListUser(ctx context.Context) ([]*model.UserInfo, error)
	RemoveUser(ctx context.Context, userId int64) error
	UpdateUser(ctx context.Context, user *model.UserInfo) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) AddUser(ctx context.Context, user *model.UserInfo) (int64, error) {
	builder := sq.Insert(table.UsersTable).
		PlaceholderFormat(sq.Dollar).
		Columns("name, age, email").
		Values(user.Name, user.Age, user.Email).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *userRepository) GetUser(ctx context.Context, userId int64) (*model.UserInfo, error) {
	builder := sq.Select("id, name, age, email").
		PlaceholderFormat(sq.Dollar).
		From(table.UsersTable).
		Where(sq.Eq{"id": userId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []*model.UserInfo
	err = r.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	if len(res) <= 0 {
		return nil, errors.New("user not found")
	}

	return res[0], nil
}

func (r *userRepository) MultiAddUser(ctx context.Context, users []*model.UserInfo) (int64, error) {
	builder := sq.Insert(table.UsersTable).
		PlaceholderFormat(sq.Dollar).
		Columns("name, age, email").
		Suffix("returning id")

	for _, user := range users {
		builder = builder.Values(user.Name, user.Age, user.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	userIds := make([]int64, 0, len(users))
	for row.Next() {
		var id int64
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
		userIds = append(userIds, id)
	}

	return int64(len(userIds)), nil
}

func (r *userRepository) ListUser(ctx context.Context) ([]*model.UserInfo, error) {
	builder := sq.Select("id, name, age, email").
		PlaceholderFormat(sq.Dollar).
		From(table.UsersTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var usersInfo []*model.UserInfo
	err = r.db.SelectContext(ctx, &usersInfo, query, args...)
	if err != nil {
		return nil, err
	}

	return usersInfo, nil
}

func (r *userRepository) RemoveUser(ctx context.Context, userId int64) error {
	builder := sq.Delete(table.UsersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.UserInfo) error {
	builder := sq.Update(table.UsersTable).
		PlaceholderFormat(sq.Dollar).
		Set("name", user.Name).
		Set("age", user.Age).
		Set("email", user.Email).
		Where(sq.Eq{"id": user.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
