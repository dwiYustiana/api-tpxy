package api

import (
	"api-tpx/config/env"
	httpHelper "api-tpx/http/helper"
	"api-tpx/model"
	"fmt"

	"github.com/labstack/echo/v4"
)

// InjectAPIHandler ...
type CreditAPIHandler struct {
	Config      env.Config
	Helper      httpHelper.HTTPHelper
	LimitModel  model.LimitModelInterface
	MemberModel model.MemberModelInterface
}

func (_h *CreditAPIHandler) PurchaseCredit(c echo.Context) error {

	var (
		err error
	)
	//get member
	memberCode := c.Param("memberCode")
	member, err := _h.MemberModel.GetMemberById(memberCode)
	if err != nil {
		return _h.Helper.SendBadRequest(c, "member not found "+err.Error(), _h.Helper.EmptyJsonMap())
	}

	fmt.Println(member)

	return _h.Helper.SendSuccess(c, "EVERYTHING IS WORKING FINE...", _h.Helper.EmptyJsonMap())
}

// HeroAddIncentiveHandler ...
// func (_h *LimitAPIHandler) GetLimitMember(c echo.Context) error {
// 	var (
// 		err error
// 	)
// 	memberCode := c.QueryParam("member_code")
// 	limits, err := _h.LimitModel.GetLimit(memberCode)
// 	if err != nil {
// 		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
// 	}
// 	return _h.Helper.SendSuccess(c, `Success`, limits)
// }
