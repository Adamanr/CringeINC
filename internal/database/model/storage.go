package model

import (
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB *pgx.Conn
}
