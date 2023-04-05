package api

type Route struct {
	FuncName string
	Method   string
	Request  string
	ReqType  Type
	Response string
	ResType  Type
	Path     string
	Doc      string
}
