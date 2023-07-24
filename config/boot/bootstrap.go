package boot

import (
	"api-tpx/config/env"
	"api-tpx/http/api"
	httpHelper "api-tpx/http/helper"
	"api-tpx/model"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	ut "github.com/go-playground/universal-translator"
	appMiddleware "github.com/labstack/echo/v4/middleware"
)

// HTTPHandler ...
type HTTPHandler struct {
	E               *echo.Echo
	Config          env.Config
	Helper          httpHelper.HTTPHelper
	ValidatorDriver *validator.Validate
	Translator      ut.Translator
}

// RegisterAPIHandler ...
func (h *HTTPHandler) RegisterAPIHandler() *HTTPHandler {
	h.Helper = httpHelper.HTTPHelper{
		Validate:   h.ValidatorDriver,
		Translator: h.Translator,
	}

	// model initialize
	db := model.Info{Config: h.Config}
	limitModel := model.NewLimitModel(db.Connect())
	memberModel := model.NewMemberModel(db.Connect())

	//test
	apiHandler := api.InjectAPIHandler{
		Config: h.Config,
		Helper: h.Helper,
	}

	//limit
	limitHandler := api.LimitAPIHandler{
		Config:     h.Config,
		Helper:     h.Helper,
		LimitModel: limitModel,
	}

	//Transaction
	creditHandler := api.CreditAPIHandler{
		Config:      h.Config,
		Helper:      h.Helper,
		LimitModel:  limitModel,
		MemberModel: memberModel,
	}

	router := h.E
	router.GET("/ping", apiHandler.PingHandler)

	group := router.Group(`api`)
	// group.Use(appMiddleware.JWTWithConfig(appMiddleware.JWTConfig{
	// 	SigningMethod: "HS512",
	// 	SigningKey:    []byte(h.Config.GetString("app.secret")),
	// }))
	group.Use(appMiddleware.KeyAuthWithConfig(appMiddleware.KeyAuthConfig{
		KeyLookup: "header:Authorization",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == h.Config.GetString("app.secret"), nil
		},
	}))

	group.GET("/limit", limitHandler.GetLimitMember)
	group.POST("/transaction/:memberCode", creditHandler.PurchaseCredit)

	return h
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "airpax-tpx-Template/1.0")
		return next(c)
	}
}

// RegisterMiddleware ...
func (h *HTTPHandler) RegisterMiddleware() {
	h.E.Use(serverHeader)
	h.E.Use(appMiddleware.CORSWithConfig(appMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	h.E.Use(appMiddleware.GzipWithConfig(appMiddleware.GzipConfig{
		Level: 5,
	}))

	if h.Config.GetBool(`app.debug`) == true {
		h.E.Use(appMiddleware.Logger())
		h.E.HideBanner = true
		h.E.Debug = true
	} else {
		h.E.HideBanner = true
		h.E.Debug = false
		h.E.Use(appMiddleware.Recover())
	}
}
