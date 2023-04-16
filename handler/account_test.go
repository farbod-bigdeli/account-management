package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/farbod-bigdeli/account-management/domain"
	"github.com/farbod-bigdeli/account-management/repository"
	"github.com/farbod-bigdeli/account-management/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	accRepo    = repository.NewAccountRepo()
	accService = service.NewAccountService(accRepo)
	accHandler = NewAccountHandler(accService)
)



func TestAccountHandler_TransferSuccessful(t *testing.T) {
	err := repository.IngestAccounts()
	if err != nil {
		t.Errorf("error loading data %v", err)
	}

	transfer := domain.Transfer{
		FromAccountID: "4c0b1e82-2df3-48a1-a297-094cddde5546",
		ToAccountID: "db6d4e22-ed9a-4bac-8b50-859d17e83591",
		Amount: 100,
	}

	marshaledTransfer, _ := json.Marshal(transfer)
	body := bytes.NewReader(marshaledTransfer)
	req := httptest.NewRequest(http.MethodPost, "/transfer", body)
	w := httptest.NewRecorder()
	accHandler.Transfer(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "transaction successful!" {
		fmt.Println(string(data))
		t.Errorf("error in transaction")
	}
}

func TestAccountHandler_TransferErrorSenderNotFound(t *testing.T) {
	err := repository.IngestAccounts()
	if err != nil {
		t.Errorf("error loading data %v", err)
	}

	transfer := domain.Transfer{
		FromAccountID: "4c0b1e82-2df3-48a1-a297-094cddde5547",
		ToAccountID: "db6d4e22-ed9a-4bac-8b50-859d17e83591",
		Amount: 100,
	}

	marshaledTransfer, _ := json.Marshal(transfer)
	body := bytes.NewReader(marshaledTransfer)
	req := httptest.NewRequest(http.MethodPost, "/transfer", body)
	w := httptest.NewRecorder()
	accHandler.Transfer(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "wrong sender account id\n" {
		t.Errorf("bad error returned")
	}
}

func TestAccountHandler_TransferErrorInsufficientFunds(t *testing.T) {
	err := repository.IngestAccounts()
	if err != nil {
		t.Errorf("error loading data %v", err)
	}

	transfer := domain.Transfer{
		FromAccountID: "4c0b1e82-2df3-48a1-a297-094cddde5546",
		ToAccountID: "db6d4e22-ed9a-4bac-8b50-859d17e83591",
		Amount: 10000,
	}

	marshaledTransfer, _ := json.Marshal(transfer)
	body := bytes.NewReader(marshaledTransfer)
	req := httptest.NewRequest(http.MethodPost, "/transfer", body)
	w := httptest.NewRecorder()
	accHandler.Transfer(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "insufficient funds\n" {
		t.Errorf("bad error returned")
	}
}

