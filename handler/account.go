package handler

import (
	"encoding/json"
	"github.com/farbod-bigdeli/account-management/domain"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/farbod-bigdeli/account-management/service"
)

type AccountHandler struct {
	accountService  *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}


func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountService.GetAccounts(r.Context())
	if err != nil {
		http.Error(w, "could not get accounts", http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(accounts)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *AccountHandler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accounts, err := h.accountService.GetAccountByID(r.Context(), vars["id"])
	if err != nil {
		http.Error(w, "could not get accounts", http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(accounts)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *AccountHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var transfer domain.Transfer
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		http.Error(w, "bad input", http.StatusBadRequest)
		return
	}

	err = h.accountService.Transfer(r.Context(), transfer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("transaction successful!"))
}
