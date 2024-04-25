package payment

type Payment struct {
	Id         int64
	PaymentId  string
	UserId     string
	MerchantId string
	Amount     float32
}
