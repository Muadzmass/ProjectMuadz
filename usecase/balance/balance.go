package balance

type Balance struct {
}

func (b *Balance) CheckBalance() (float32, error) {
	return 1000, nil
}

func NewBalance() *Balance {
	return &Balance{}
}
