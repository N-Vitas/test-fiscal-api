package FiscalTestApi

/*
 * Авторизация в rpc
 * {"auth":{"phone": "demo","password": "demo","idTerminal":1069}}
*/

type Login struct {
	Auth Auth `json:"auth"`
}

type Auth struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
	IdTerminal int64 `json:"idTerminal"`
}

func NewLogin() Login  {
	return Login{Auth:Auth{"demo","demo",1069}}
}
