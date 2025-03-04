package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func NewPostgres(cfg DBConfig) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)

	}

	return conn, nil
}
