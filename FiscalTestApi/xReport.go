package FiscalTestApi

func NewReport() HeadRPCInfo {
	return HeadRPCInfo{
		HRpc{1, "1Q2w3e4r"},
		InfoReq{106, 0, ""},
	}
}

type XReport struct {
	Header HRpc
	Status StatusReport
}
type StatusReport struct {
	Code    int64
	Message MessageReport
}
type MessageReport struct {
	BuyCount           int64
	BuySum             float64
	Com                int64
	HourWork           int64
	InCount            int64
	InSum              float64
	LeftCash           int64
	List               []interface{}
	NDS                int64
	OpenDay            string
	OutCount           int64
	OutSum             float64
	PayCardSum         float64
	PayCount           int64
	PayCreditSum       float64
	PaySum             float64
	PayTareSum         float64
	ReturnBuyCount     int64
	ReturnBuySum       float64
	ReturnPayCardSum   float64
	ReturnPayCount     int64
	ReturnPayCreditSum float64
	ReturnPaySum       float64
	ReturnPayTareSum   float64
	SysInfo            SysInfo
}
