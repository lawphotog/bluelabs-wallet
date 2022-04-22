package wallet

import (
	"errors"
	"testing"

	"bluelabs/wallet/internal/repository"
	"bluelabs/wallet/internal/wallet/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "wallet service")
}

var _ = Describe("Wallet", func() {
	Describe("when Create is called", func() {
		It("should pass the right userId", func() {
			mockRepo := &mocks.DynamoRepository{}

			//assert
			mockRepo.On("Create", "1").Return(nil)

			service := New(mockRepo)
			service.Create("1")
		})

		It("should return error if errored by repo", func() {
			mockRepo := &mocks.DynamoRepository{}

			expectErr := errors.New("something wrong")
			mockRepo.On("Create", "1").Return(expectErr)

			service := New(mockRepo)
			err := service.Create("1")
			Expect(err).To(Equal(expectErr))
		})
	})

	Describe("when Deposit is called", func() {
		It("should return error when repository returns error", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
			}

			expect := errors.New("something wrong")

			mockRepo.On("Get", mock.Anything).Return(wallet, expect)

			service := New(mockRepo)
			err := service.Deposit("1", 100)

			Expect(err).To(Equal(expect))
		})

		It("should update sequence number", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
			}

			mockRepo.On("Get", mock.Anything).Return(wallet, nil)

			//assert
			mockRepo.On("Update", mock.MatchedBy(func(wallet repository.Wallet) bool {
				return wallet.UpdateSequence == "1"
			})).Return(nil)

			service := New(mockRepo)
			service.Deposit("1", 100)
		})

		It("should add a transaction with right amount as deposit", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
			}

			mockRepo.On("Get", mock.Anything).Return(wallet, nil)

			//assert
			mockRepo.On("Update", mock.MatchedBy(func(wallet repository.Wallet) bool {
				return wallet.Transactions[0].Amount == 100 && 
					wallet.Transactions[0].TransactionType == 0
			})).Return(nil)

			service := New(mockRepo)
			service.Deposit("1", 100)
		})
	})

	Describe("when Withdraw is called", func() {
		It("should return error when repository returns error", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
				Transactions: []repository.Transaction{
					{
						TransactionType: 0,
						Amount: 100,
					},
				},
			}

			expect := errors.New("something wrong")

			mockRepo.On("Get", mock.Anything).Return(wallet, expect)

			service := New(mockRepo)
			err := service.Withdraw("1", 100)

			Expect(err).To(Equal(expect))
		})

		It("should update sequence number", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
				Transactions: []repository.Transaction{
					{
						TransactionType: 0,
						Amount: 100,
					},
				},
			}

			mockRepo.On("Get", mock.Anything).Return(wallet, nil)

			//assert
			mockRepo.On("Update", mock.MatchedBy(func(wallet repository.Wallet) bool {
				return wallet.UpdateSequence == "1"
			})).Return(nil)

			service := New(mockRepo)
			service.Withdraw("1", 100)
		})

		It("should add a transaction with right amount as withdraw", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
				Transactions: []repository.Transaction{
					{
						TransactionType: 0,
						Amount: 100,
					},
				},
			}

			mockRepo.On("Get", mock.Anything).Return(wallet, nil)

			//assert
			mockRepo.On("Update", mock.MatchedBy(func(wallet repository.Wallet) bool {
				return wallet.Transactions[1].Amount == 100 && 
					wallet.Transactions[1].TransactionType == 1
			})).Return(nil)

			service := New(mockRepo)
			service.Withdraw("1", 100)
		})

		It("should not let negative balance", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
			}

			mockRepo.On("Get", mock.Anything).Return(wallet, nil)

			service := New(mockRepo)
			err := service.Withdraw("1", 100)
			Expect(err.Error()).To(Equal("not enough balance to withdraw this amount"))
		})
	})

	Describe("when Balance is called", func() {
		It("should return zero balance on new wallet", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
			}
			mockRepo.On("Get", mock.Anything).Return(wallet, nil)
			service := New(mockRepo)
			balance, _ := service.Balance("1")

			var expect int64 = 0
			Expect(balance).To(Equal(expect))
		})

		It("should get the right balance", func() {
			mockRepo := &mocks.DynamoRepository{}
			wallet := repository.Wallet{
				UserId: "1",
				UpdateSequence: "0",
				Transactions: []repository.Transaction{
					{
						TransactionType: 0,
						Amount: 100,
					},
					{
						TransactionType: 1,
						Amount: 50,
					},
				},
			}
			mockRepo.On("Get", mock.Anything).Return(wallet, nil)
			service := New(mockRepo)
			balance, _ := service.Balance("1")

			var expect int64 = 50
			Expect(balance).To(Equal(expect))
		})
	})
})
