package v1

import (
	"Hyuga/core/api"
	"Hyuga/core/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetRecords get records
func GetRecords(c echo.Context) error {
	rtype := c.QueryParam("type")
	token := c.QueryParam("token")
	filter := c.QueryParam("filter")

	log.Debug(fmt.Sprintf("api/v1/GetRecords: type=%s token=%s filter=%s", rtype, token, filter))

	records, err := utils.Recorder.GetRecords(rtype, token, filter)
	if err != nil {
		code, resp := api.ProcessingError(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, &api.RespJSON{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    records,
		Success: true,
	})
}
