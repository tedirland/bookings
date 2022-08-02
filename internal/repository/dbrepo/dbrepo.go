package dbrepo

import (
	"database/sql"

	"github.com/tedirland/bookings/internal/config"
	"github.com/tedirland/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// If you want to create a new database, create a new type and a new func

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}

}
