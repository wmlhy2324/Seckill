// Code generated by goctl. DO NOT EDIT.
package types

type AuthenticationRequest struct {
	Token     string `header:"Token,optional"`
	ValidPath string `header:"ValidPath,optional"`
}

type AuthenticationResponse struct {
	UserID int64 `json:"userID"`
	Role   int   `json:"role"`
}

type CreateManagerRequest struct {
	Nickname string `json:"nickName"`
	Password string `json:"password"`
}

type CreateManagerResponse struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickName"`
	Role     int    `json:"role"`
}

type CreateRequest struct {
	Nickname string `json:"nickName"`
	Password string `json:"password"`
}

type CreateResponse struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickName"`
	Role     int    `json:"role"`
}

type LoginRequest struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}