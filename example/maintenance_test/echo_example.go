package main

import (
	common_middleware "github.com/huy-quang-vmo/common_middleware/maintenance"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	checkMaintenanceStatusStore := NewCheckMaintenanceStatusStore()
	maintenanceMiddleware := common_middleware.NewMaintenanceMiddleware(checkMaintenanceStatusStore)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			handler := maintenanceMiddleware.MaintenanceStatus(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_ = next(c)
			}))
			handler.ServeHTTP(c.Response(), c.Request())
			return nil
		}
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Start(":8080")
}
