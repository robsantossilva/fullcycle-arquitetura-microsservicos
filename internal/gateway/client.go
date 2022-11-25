package gateway

import "github.com/robsantossilva/fullcycle-arquitetura-microsservicos/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
