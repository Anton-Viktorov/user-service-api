package user_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	host       = "localhost"
	port       = "54321"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "user_service_api"
	sslMode    = "disable"
)

func (i *Implementation) AddUser(ctx context.Context, req *desc.AddUserRequest) (*desc.AddUserResponse, error) {
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
		Values(req.GetName(), req.GetAge(), req.GetEmail()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &desc.AddUserResponse{
		Result: &desc.AddUserResponse_Result{
			Id: id,
		},
	}, nil
}
