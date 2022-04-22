package mocks

import (
	"github.com/stretchr/testify/mock"
)

type DynamoRepository struct {
	mock.Mock
}

func (g *DynamoRepository) Create(userId string) error {
	args := g.Called(userId)
	return args.Error(0)
}
