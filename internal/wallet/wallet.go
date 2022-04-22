package wallet

import (
	"bluelabs/wallet/internal/repository"
	"errors"
	"strconv"
)

type wallet struct {
	repository   repository.Repository
}

type Wallet interface {
	Create(userId string) error
	Deposit(userId string, amount int64) error
	Withdraw(userId string, amount int64) error
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

func (w *wallet) Withdraw(userId string, amount int64) error {
	wallet, err := w.repository.Get(userId)
	if err != nil {
		return err
	}

	balance, err := getBalance(wallet.Transactions)
	if err != nil {
		return errors.New("something went wrong")
	}

	if balance-amount < 0 {
		return errors.New("not enough balance to withdraw this amount")
	}

	transaction := repository.Transaction{
		TransactionType: withdraw,
		Amount:          amount,
	}
	wallet.Transactions = append(wallet.Transactions, transaction)
	updateSequence(&wallet)

	w.repository.Update(wallet)
	return nil
}

func updateSequence(wallet *repository.Wallet) {
	updateSeq, _ := strconv.Atoi(wallet.UpdateSequence)
	updateSeq += 1
	wallet.UpdateSequence = strconv.Itoa(updateSeq)
}

func getBalance(transactions []repository.Transaction) (int64, error) {
	var balance int64 = 0
	for _, v := range transactions {
		if v.TransactionType == add {
			balance += v.Amount
		} else {
			balance -= v.Amount
		}
	}
	return balance, nil
}