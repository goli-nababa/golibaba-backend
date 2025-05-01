package domain

type PaymentCallback struct {
	Authority string
	Status    string
	Amount    int64
}
