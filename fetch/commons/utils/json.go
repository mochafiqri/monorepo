package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/monorepo/fetch/commons/dtos"
	"net/http"
)

func WriteResponse(c echo.Context, req dtos.StandardResponseReq) error {
	var errResp interface{}
	if req.Message == "" {
		req.Message = http.StatusText(req.Code)
	}
	if req.Error != nil {
		errResp = req.Error
	}
	return c.JSON(req.Code, dtos.StandardResponse{
		Code:    req.Code,
		Status:  http.StatusText(req.Code),
		Message: req.Message,
		Data:    req.Data,
		Meta:    req.Meta,
		Error:   errResp,
	})
}
