package user_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"github.com/jmoiron/sqlx"
)

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {

	// dirty implementation
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("id, name, age, email").
		PlaceholderFormat(sq.Dollar).
		From(usersTable).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []user
	err = db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	return &desc.GetUserResponse{
		Result: &desc.GetUserResponse_User{
			Id:    res[0].Id,
			Name:  res[0].Name,
			Age:   res[0].Age,
			Email: res[0].Email,
		},
	}, nil
}
