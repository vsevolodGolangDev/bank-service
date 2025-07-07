package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/vsevolodGolangDev/bank-service/pkg/logging"
)

type Repo struct {
	User
}

func NewRepo(db *sqlx.DB, log logging.Logger) *Repo {
	return &Repo{
		User: NewUserRepo(db, log),
	}
}
