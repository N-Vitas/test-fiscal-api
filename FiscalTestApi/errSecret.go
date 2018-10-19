package FiscalTestApi

func NewErrorSecret() HeadRPCInfo {
	return HeadRPCInfo{
		HRpc{1, ""},
		InfoReq{106, 0, ""},
	}
}
