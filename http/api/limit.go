package api

import (
	"api-tpx/config/env"
	httpHelper "api-tpx/http/helper"
	"api-tpx/model"

	"github.com/labstack/echo/v4"
)

// InjectAPIHandler ...
type LimitAPIHandler struct {
	Config     env.Config
	Helper     httpHelper.HTTPHelper
	LimitModel model.LimitModelInterface
}

func (_h *LimitAPIHandler) GetLimitMember(c echo.Context) error {
	var (
		err error
	)
	memberCode := c.QueryParam("member_code")
	limits, err := _h.LimitModel.GetLimit(memberCode)
	if err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}
	return _h.Helper.SendSuccess(c, `Success`, limits)
}
