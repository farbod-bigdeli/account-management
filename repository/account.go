package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/farbod-bigdeli/account-management/domain"
	"io"
	"os"
	"strconv"
	"sync"
)

var mx = &sync.RWMutex{}


var accountsDB map[string]*account


type account struct {
	Name    string
	Balance float64
}

type AccountRepo struct {}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{}
}

func (r *AccountRepo) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	mx.RLock()
	defer mx.RUnlock()

	var accounts []*domain.Account
	for id, v := range accountsDB {
		accounts = append(accounts, &domain.Account{
			ID: id,
			Name: v.Name,
			Balance: v.Balance,
		})
	}

	return accounts, nil
}

func (r *AccountRepo) GetAccountByID(ctx context.Context, id string) (*domain.Account, error) {
	mx.RLock()
	defer mx.RUnlock()

	acc, ok := accountsDB[id]

	if !ok {
		return nil, fmt.Errorf("account not found %s", id)
	}

	accountDomain := &domain.Account{
		ID: id,
		Name: acc.Name,
		Balance: acc.Balance,
	}

	return accountDomain, nil
}

func (r *AccountRepo) UpdateBalance(ctx context.Context, id string, balance float64) {
	mx.Lock()
	accountsDB[id].Balance = balance
	mx.Unlock()
}

func IngestAccounts() error {
	accountsDB = make(map[string]*account)
	f, err := os.Open("accounts.json")
	if err != nil {
		return fmt.Errorf("error openning accounts file: %s", err)
	}


	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading accounts json: %s", err)
	}


	var result []struct {
		ID      string
		Name    string
		Balance string
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("error unmarshaling balances json api: %s", err)
	}


	for _, v := range result {
		b, err := strconv.ParseFloat(v.Balance, 64)
		if err != nil {
			return fmt.Errorf("id: %s - could not pars balance to float", v.ID)
		}

		accountsDB[v.ID] = &account{
			Name: v.Name,
			Balance: b,
		}
	}

	return nil
}