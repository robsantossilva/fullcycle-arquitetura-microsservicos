package gateway

import "github.com/robsantossilva/fullcycle-arquitetura-microsservicos/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
