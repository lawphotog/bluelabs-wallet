package wallet

import (
	"bluelabs/wallet/internal/repository"
)

type wallet struct {
	repository   repository.Repository
}

type Wallet interface {
	Create(userId string) error
}

func New(repository repository.Repository) Wallet {
	return &wallet{
		repository: repository,
	}
}

func (w *wallet) Create(userId string) error {
	err := w.repository.Create(userId)
	if err != nil {
		return err
	}

	return nil
}