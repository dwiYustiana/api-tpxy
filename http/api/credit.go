package api

import (
	"api-tpx/config/env"
	httpHelper "api-tpx/http/helper"
	"api-tpx/model"
	"fmt"
	"log"

	"api-tpx/http/request"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
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
		err   error
		input request.PurchaseCreditRequest
	)
	//get member
	memberCode := c.Param("memberCode")
	member, err := _h.MemberModel.GetMemberById(memberCode)
	if err != nil {
		return _h.Helper.SendBadRequest(c, "member not found "+err.Error(), _h.Helper.EmptyJsonMap())
	}

	err = c.Bind(&input)
	if err != nil {
		log.Println("Error HeroAddIncentiveHandler Bind : " + err.Error())
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}
	if err = _h.Helper.Validate.Struct(input); err != nil {
		log.Println("Error PurchaseCreditHandler Validation input: " + err.Error())
		return _h.Helper.SendValidationError(c, err.(validator.ValidationErrors))
	}

	//get limit
	// limit := _h.LimitModel.GetLimitById(input.LimitCode)
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
