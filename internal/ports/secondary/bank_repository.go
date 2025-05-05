package secondary

import (
	"ams-service/internal/core/entities"
)

type BankRepository interface {
	CreateCreditCard(card *entities.CreditCard) (*entities.CreditCard, error)
	FindCreditCard(info *entities.CreditCardInfo) (entities.CreditCard, error)
	UpdateCreditCard(card *entities.CreditCard) (*entities.Transaction, error)
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	FindTransactionById(id string) (*entities.Transaction, error)
}
