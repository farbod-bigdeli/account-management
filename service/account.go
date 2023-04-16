package service

import (
	"context"
	"fmt"
	"github.com/farbod-bigdeli/account-management/domain"
	"github.com/farbod-bigdeli/account-management/repository"
)

type AccountService struct {
	accountRepo  *repository.AccountRepo
}

func NewAccountService(accountRepo *repository.AccountRepo) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

func (s *AccountService) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	return s.accountRepo.GetAccounts(ctx)
}

func (s *AccountService) GetAccountByID(ctx context.Context, id string) (*domain.Account, error) {
	return s.accountRepo.GetAccountByID(ctx, id)
}

func (s *AccountService) Transfer(ctx context.Context, transfer domain.Transfer) error {
	sender, err := s.accountRepo.GetAccountByID(ctx, transfer.FromAccountID)
	if err != nil {
		return fmt.Errorf("wrong sender account id")
	}

	receiver, err := s.accountRepo.GetAccountByID(ctx, transfer.ToAccountID)
	if err != nil {
		return fmt.Errorf("wrong receiver account id")
	}


	if transfer.Amount > sender.Balance {
		return fmt.Errorf("insufficient funds")
	}

	s.accountRepo.UpdateBalance(ctx, sender.ID, sender.Balance - transfer.Amount)
	s.accountRepo.UpdateBalance(ctx, receiver.ID, receiver.Balance + transfer.Amount)

	return nil
}
