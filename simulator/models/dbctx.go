package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBCtx struct {
	Pool               *pgxpool.Pool
	DBConnectionString string
	Ctx                context.Context
}
