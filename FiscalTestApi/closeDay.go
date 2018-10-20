package FiscalTestApi

func NewCloseDay() HeadDefRPC {
	return HeadDefRPC{
		HRpc{1, "1Q2w3e4r"},
		struct {
			OpCode int64
		}{7},
	}
}

type Operation struct {
	Header HRpc
	Status StatusOperation
}

type StatusOperation struct {
	Code        int64
	Transaction Types
}
type Types struct {
	Type     int64
}