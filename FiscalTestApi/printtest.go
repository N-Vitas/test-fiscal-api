package FiscalTestApi

func NewPrint() HeadRPCInfo {
	return HeadRPCInfo{
		HRpc{1, "1Q2w3e4r"},
		InfoReq{104, 0, "Тестовая печать"},
	}
}
