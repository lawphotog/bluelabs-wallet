package wallet

import (
	"bluelabs/wallet/internal/repository"
	"strconv"
)

type wallet struct {
	repository   repository.Repository
}

type Wallet interface {
	Create(userId string) error
	Deposit(userId string, amount int64) error
}

const (
	add = iota
	withdraw
)

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

func (w *wallet) Deposit(userId string, amount int64) error {
	transaction := repository.Transaction{
		TransactionType: add,
		Amount:          amount,
	}

	wallet, err := w.repository.Get(userId)
	if err != nil {
		return err
	}

	wallet.Transactions = append(wallet.Transactions, transaction)
	updateSequence(&wallet)

	err = w.repository.Update(wallet)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func updateSequence(wallet *repository.Wallet) {
	updateSeq, _ := strconv.Atoi(wallet.UpdateSequence)
	updateSeq += 1
	wallet.UpdateSequence = strconv.Itoa(updateSeq)
}