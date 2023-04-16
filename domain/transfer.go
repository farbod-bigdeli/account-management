package domain

type Transfer struct {
	FromAccountID string  `json:"from"`
	ToAccountID   string  `json:"to"`
	Amount        float64 `json:"amount"`
}
