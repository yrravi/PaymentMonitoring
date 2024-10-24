package models

type Payment struct {
	PaymentReference   string  `json:"payment_reference" gorm:"primaryKey"`
	Amount             float64 `json:"amount"`
	SenderAccountRef   string  `json:"sender_account_ref"`
	ReceiverAccountRef string  `json:"receiver_account_ref"`
	Mode               string  `json:"mode"`
	Currency           string  `json:"currency"`
	Source             string  `json:"source"`
	//CreatedAt          string  `json:"created_at" gorm:"autoCreateTime"`
}
