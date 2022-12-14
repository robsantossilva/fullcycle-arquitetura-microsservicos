package gateway

import "github.com/robsantossilva/fullcycle-arquitetura-microsservicos/internal/entity"

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
