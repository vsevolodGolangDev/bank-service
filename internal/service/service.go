package service

import (
	"github.com/vsevolodGolangDev/bank-service/internal/repo"
	"github.com/vsevolodGolangDev/bank-service/models"
	"github.com/vsevolodGolangDev/bank-service/pkg/logging"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service struct {
	User
}

type User interface {
	GetBalance(id int, currency string) (float32, error)
	GetTransactions(id int, page models.Page) ([]models.Transaction, error)
	TopUp(input models.Input) (float32, error)
	Debit(input models.Input) (float32, error)
	Transfer(input models.TransferInput) (float32, error)
}

func NewService(repo *repo.Repo, log logging.Logger) *Service {
	return &Service{
		User: NewUserService(repo, log),
	}
}
