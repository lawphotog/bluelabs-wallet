package repository

type Wallet struct {
	UserId       string `dynamodbav:"userId"`
	Transactions []Transaction
	UpdateSequence     string
}

type Transaction struct {
	TransactionType int
	Amount          int64
}
