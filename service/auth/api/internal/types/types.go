// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Code string `form:"code" @query:"code"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserId int    `json:"userId"`
	Page   string `json:"page"`
}
