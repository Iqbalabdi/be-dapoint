package xendit

type XenditService interface {
	CreateCharge(param string) (interface{}, error)
}
