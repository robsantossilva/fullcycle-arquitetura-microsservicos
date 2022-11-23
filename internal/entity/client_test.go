package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Robson", "r@s.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Robson", client.Name)
	assert.Equal(t, "r@s.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, err := NewClient("Robson", "r@s.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	err = client.Update("Hugo", "h@s.com")
	assert.Nil(t, err)
	assert.Equal(t, "Hugo", client.Name)
	assert.Equal(t, "h@s.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Robson", "r@s.com")
	err := client.Update("", "h@s.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Robson", "r@s.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))

	client2, _ := NewClient("Robson 2", "r@s.com")
	err = client2.AddAccount(account)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "account does not belong to client")
}
