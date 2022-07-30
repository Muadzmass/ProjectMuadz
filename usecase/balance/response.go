package balance

import "apiproject/constants"

type Data struct {
	Amount float32 `json:"amount"`
}

type Response struct {
	Status bool                     `json:"status"`
	Data   Data                     `json:"data"`
	Errors []constants.GenericError `json:"errors"`
}
