package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/farbod-bigdeli/account-management/handler"
	"github.com/farbod-bigdeli/account-management/repository"
	"github.com/farbod-bigdeli/account-management/service"
)

func main() {
	// Ingest accounts from JSON file
	err := repository.IngestAccounts()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("accounts ingested successfully...")

	accRepo := repository.NewAccountRepo()

	accService := service.NewAccountService(accRepo)


	accHandler := handler.NewAccountHandler(accService)


	r := mux.NewRouter()
	// Define endpoints
	r.HandleFunc("/account/{id}", accHandler.GetAccountByID).Methods("GET")
	r.HandleFunc("/accounts", accHandler.GetAccounts).Methods("GET")
	r.HandleFunc("/transfer", accHandler.Transfer).Methods("POST")


	// Start server on port 8080
	log.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

//func makeTransfer(w http.ResponseWriter, r *http.Request) {
//	// Parse transfer request from JSON body
//	// Validate transfer request
//	// Update account balances
//	// Return success message in JSON format
//	var transfer Transfer
//	err := json.NewDecoder(r.Body).Decode(&transfer)
//	if err != nil {
//		http.Error(w, "bad input", http.StatusBadRequest)
//		return
//	}

	//sender, ok := accounts[transfer.FromAccountID]
	//if !ok {
	//	http.Error(w, "wrong sender account id", http.StatusBadRequest)
	//	return
	//}
	//
	//
	//_, ok = accounts[transfer.FromAccountID]
	//if !ok {
	//	http.Error(w, "wrong receiver account id", http.StatusBadRequest)
	//	return
	//}
	//
	//
	//if transfer.Amount > sender.Balance {
	//	http.Error(w, "insuficcient funds", http.StatusBadRequest)
	//	return
	//}
	//
	//
	//mutex.Lock()
	//// sender.Balance -= transfer.Amount
	//accounts[transfer.FromAccountID].Balance -= transfer.Amount
	//accounts[transfer.ToAccountID].Balance += transfer.Amount
	//mutex.Unlock()

//
//
//}
