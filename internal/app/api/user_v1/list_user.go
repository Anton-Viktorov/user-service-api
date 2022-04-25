package user_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"github.com/jmoiron/sqlx"
)

type user struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Age   int64  `db:"age"`
	Email string `db:"email"`
}

func (i *Implementation) ListUser(ctx context.Context, req *desc.Empty) (*desc.ListUserResponse, error) {
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
		From(usersTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []user
	err = db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	users := make([]*desc.ListUserResponse_Result_User, 0, len(res))
	for _, u := range res {
		users = append(users, &desc.ListUserResponse_Result_User{
			Id:    u.Id,
			Name:  u.Name,
			Age:   u.Age,
			Email: u.Email,
		})
	}

	return &desc.ListUserResponse{
		Result: &desc.ListUserResponse_Result{
			Users: users,
		},
	}, nil
}
