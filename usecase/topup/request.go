package topup

type Request struct {
	Id     string  `json:"id"`
	Amount float32 `json:"amount"`
}
