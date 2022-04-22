package wallet

import (
	"bluelabs/wallet/internal/repository"
)

type wallet struct {
	repository   repository.Repository
}

type Wallet interface {
}

func New(repository repository.Repository) Wallet {
	return &wallet{
		repository: repository,
	}
}