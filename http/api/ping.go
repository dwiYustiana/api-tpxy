package api

import (
	"api-tpx/config/env"
	httpHelper "api-tpx/http/helper"

	"github.com/labstack/echo/v4"
)

// InjectAPIHandler ...
type InjectAPIHandler struct {
	Config env.Config
	Helper httpHelper.HTTPHelper
}

// tes API
func (_h *InjectAPIHandler) PingHandler(c echo.Context) error {
	return _h.Helper.SendSuccess(c, "EVERYTHING IS WORKING FINE...", _h.Helper.EmptyJsonMap())
}
