package alerttype

import "github.com/qwganker/boring/comm/request"

type AlertTypePageReq struct {
	request.PageRequest
}

type AlertTypeAddReq struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type AlertTypeDeleteReq struct {
	ID int64 `json:"id"`
}

type AlertTypeModifyReq struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}
