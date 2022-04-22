package mocks

import (
	"bluelabs/wallet/internal/repository"
	"github.com/stretchr/testify/mock"
)

type DynamoRepository struct {
	mock.Mock
}

func (g *DynamoRepository) Create(userId string) error {
	args := g.Called(userId)
	return args.Error(0)
}

func (g *DynamoRepository) Update(wallet repository.Wallet) error {
	args := g.Called(wallet)
	return args.Error(0)
}

func (g *DynamoRepository) Get(userId string) (repository.Wallet, error) {
	args := g.Called(userId)
	wallet := args.Get(0).(repository.Wallet)
	return wallet, args.Error(1)
}