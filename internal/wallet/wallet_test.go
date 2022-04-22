package wallet

import (
	"errors"
	"testing"

	"bluelabs/wallet/internal/wallet/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
})
