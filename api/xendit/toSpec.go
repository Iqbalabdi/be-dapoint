package xendit

type XenditCallbackPayload struct {
	Event      string      `json:"event"`
	BusinessID string      `json:"business_id"`
	Created    string      `json:"created"`
	Data       interface{} `json:"data"`
}

type XenditGetVoucher struct {
	ID      int
	Nominal float64
}
