package order


type Order struct {
	ID              int64  `json:"id,omitempty"`
	Price           string `json:"price"`
	Amount          string `json:"amount"`
	RemainingAmount string `json:"remaining_amount,omitempty"`
	Cost            string `json:"cost,omitempty"`
	Status          string `json:"status,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	Side            string `json:"side,omitempty"`
}
