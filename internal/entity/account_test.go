package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Robson Silva", "robson@dominio.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Robson Silva", "robson@dominio.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Robson Silva", "robson@dominio.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(60)
	assert.Equal(t, float64(40), account.Balance)
}
