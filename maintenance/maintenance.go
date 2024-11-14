package maintenance

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MaintenanceMiddleware struct {
	maintenanceStatus GetMaintenanceStatus
}

func NewMaintenanceMiddleware(maintenanceStatus GetMaintenanceStatus) *MaintenanceMiddleware {
	return &MaintenanceMiddleware{maintenanceStatus: maintenanceStatus}
}

type GetMaintenanceStatus interface {
	IsMaintenance() (bool, error)
}

func (m *MaintenanceMiddleware) MaintenanceStatus(c *gin.Context) {
	isMaintenanceMode, err := m.maintenanceStatus.IsMaintenance()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if isMaintenanceMode {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, "Service is under maintenance")
		return
	}
	c.Next()
}
