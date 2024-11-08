package main

import (
	common_middleware "github.com/common-middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	checkMaintenanceStatusStore := NewCheckMaintenanceStatusStore()
	maintenanceMiddleware := common_middleware.NewMaintenanceMiddleware(checkMaintenanceStatusStore)

	r.Use(func(c *gin.Context) {
		handler := maintenanceMiddleware.MaintenanceStatus(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		}))
		handler.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}

type CheckMaintenanceStatusStore struct{}

func NewCheckMaintenanceStatusStore() *CheckMaintenanceStatusStore {
	return &CheckMaintenanceStatusStore{}
}

func (c *CheckMaintenanceStatusStore) GetMaintenanceStatus() (bool, error) {
	return true, nil
}
