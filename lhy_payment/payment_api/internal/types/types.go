// Code generated by goctl. DO NOT EDIT.
package types

type PayRequest struct {
	UserId  int32
	OrderId int32
}

type PayResponse struct {
	Account float64
	Msg     string
}
