// Code generated by goctl. DO NOT EDIT.
package types

type CreateCartRequest struct {
	ItemName string `json:"itemName"`
	Num      int    `json:"num"`
	Token    string `header:"token"`
}

type CreateCartResponse struct {
	Msg string `json:"msg"`
}

type CreateItemRequest struct {
	ItemName    string  `json:"itemName"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Token       string  `header:"token"`
	Type        int     `json:"type"` //0为普通商品，1为秒杀商品
}

type CreateItemResponse struct {
	ItemName    string  `json:"itemName"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type DeleteItemRequest struct {
	ItemName string `json:"itemName"`
	Token    string `json:"token"`
}

type DeleteItemResponse struct {
	ItemName string `json:"itemName"`
	Msg      string `json:"msg"`
}

type GetCartListRequest struct {
	Token string `header:"token"`
}

type GetCartListResponse struct {
	Data []byte `json:"data"`
}

type GetItemListResponse struct {
	Data []byte `json:"data"`
}