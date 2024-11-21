package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
	*bun.DB
}

func Dial(cfg config.Config) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&Timezone=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.SSL,
		cfg.DB.Timezone)
	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(pgDB, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	var rnd float64

	if err := db.NewSelect().ColumnExpr("random()").Scan(context.Background(), &rnd); err != nil {
		return nil, errors.Wrap(err, "error on connecting to db")
	}

	return &DB{db}, nil
}
