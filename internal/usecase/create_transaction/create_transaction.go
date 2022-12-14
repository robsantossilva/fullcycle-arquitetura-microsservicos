package createtransaction

import (
	"github.com/robsantossilva/fullcycle-arquitetura-microsservicos/internal/entity"
	"github.com/robsantossilva/fullcycle-arquitetura-microsservicos/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(in CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.Get(in.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.Get(in.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, in.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}
