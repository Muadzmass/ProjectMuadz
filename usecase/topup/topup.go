package topup

import "errors"

type Topup struct {
}

func NewTopup() *Topup {
	return &Topup{}
}

func (t *Topup) DoTopup(request Request) error {
	if !validate(request) {
		return errors.New("DoTopup error validation")
	}
	return nil
}

func validate(request Request) bool {
	if request.Id == "" || request.Amount > 0 {
		return false
	}
	return true
}
