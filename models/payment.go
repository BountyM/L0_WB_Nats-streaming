package models

type Payment struct {
	Transaction   string `json:"transaction" db:"transaction"`
	Request_id    string `json:"request_id" db:"request_id"`
	Currency      string `json:"currency" db:"currency"`
	Provider      string `json:"provider" db:"provider"`
	Amount        int    `json:"amount" db:"amount"`
	Payment_dt    int    `json:"payment_dt" db:"payment_dt"`
	Bank          string `json:"bank" db:"bank"`
	Delivery_cost int    `json:"delivery_cost" db:"delivery_cost"`
	Goods_total   int    `json:"goods_total" db:"goods_total"`
	Custom_fee    int    `json:"custom_fee" db:"custom_fee"`
}
