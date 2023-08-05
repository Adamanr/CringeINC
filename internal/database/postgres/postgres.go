package postgres

import (
	"context"
	"cringeinc_server/internal/config"
	"cringeinc_server/internal/database/model"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
)

func CheckTables(storage *model.Storage, tableName string) bool {
	var exist *bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema LIKE 'public' AND table_type LIKE 'BASE TABLE' AND table_name = '%s')", tableName)
	row := storage.DB.QueryRow(context.Background(), query)

	if err := row.Scan(&exist); err != nil {
		slog.Warn("table users does not exist!", slog.String("error", err.Error()))
		log.Println(err)
		return false
	}

	return *exist
}

func New(cfg *config.Database) (*model.Storage, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Name)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	m, err := migrate.New("file://internal/database/migrations", connStr)

	if err != nil {
		slog.Error("error new migrate on connect!", slog.String("error", err.Error()))
		return nil, errors.New("error new migrate on connect")
	}

	if err := m.Up(); err != nil {
		slog.Error("error up migrate!", slog.String("error", err.Error()))
	}

	slog.Info("The database is connected", slog.String("conn", connStr))
	return &model.Storage{conn}, nil
}
