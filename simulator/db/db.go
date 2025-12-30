package db

import (
	"context"
	"fmt"

	"github.com/NaveenChander/GoFace/simulator/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DbContext models.DBCtx

func GetDBContext() {
	ctx := context.Background()
	connStr := models.DBConnectionString
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		fmt.Printf("Unable to ping database: %v\n", err)
	}
	fmt.Println("Connected to the database successfully!")

	DbContext = models.DBCtx{
		Pool:               pool,
		DBConnectionString: connStr,
		Ctx:                ctx,
	}
}
