package user_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"github.com/jmoiron/sqlx"
)

func (i *Implementation) MultiAddUser(ctx context.Context, req *desc.MultiAddUserRequest) (*desc.MultiAddUserResponse, error) {
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

	builder := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns("name, age, email").
		Suffix("returning id")

	for _, user := range req.GetUsers() {
		builder = builder.Values(user.GetName(), user.GetAge(), user.GetEmail())
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	userIds := make([]int64, 0, len(req.GetUsers()))
	for row.Next() {
		var id int64
		err = row.Scan(&id)
		userIds = append(userIds, id)
	}

	return &desc.MultiAddUserResponse{
		Result: &desc.MultiAddUserResponse_Result{
			Count: int64(len(userIds)),
		},
	}, nil
}
