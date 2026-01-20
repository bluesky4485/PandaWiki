package v1

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/chaitin/panda-wiki/config"
	"github.com/chaitin/panda-wiki/consts"
	"github.com/chaitin/panda-wiki/handler"
	"github.com/chaitin/panda-wiki/log"
)

type LicenseHandler struct {
	*handler.BaseHandler
	logger *log.Logger
	config *config.Config
}

func NewLicenseHandler(e *echo.Echo, baseHandler *handler.BaseHandler, logger *log.Logger, config *config.Config) *LicenseHandler {
	handlerLogger := logger.WithModule("handler.v1.license")
	h := &LicenseHandler{
		BaseHandler: baseHandler,
		logger:      handlerLogger,
		config:      config,
	}

	group := e.Group("/api/v1/license")
	group.GET("", h.GetLicense)

	return h
}

func (h *LicenseHandler) GetLicense(c echo.Context) error {
	return h.NewResponseWithData(c, map[string]interface{}{
		"edition":    consts.LicenseEditionEnterprise,
		"expired_at": time.Now().AddDate(99, 0, 0).Unix(),
		"started_at": time.Now().Unix(),
	})
}
