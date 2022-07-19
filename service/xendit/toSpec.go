package xendit

type XenditService interface {
	CreateCharge(id uint64, param string) (interface{}, error)
	PaymentStatusCallback(param string) (res interface{}, err error)
}
