package topup

import "apiproject/constants"

type Response struct {
	Status bool                     `json:"status"`
	Errors []constants.GenericError `json:"errors"`
}
